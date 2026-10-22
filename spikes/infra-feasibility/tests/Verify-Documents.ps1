$ErrorActionPreference='Stop'
$spikeRoot=Split-Path $PSScriptRoot
$repoRoot=(git -C $spikeRoot rev-parse --show-toplevel).Trim()
$files=@(git -C $repoRoot ls-files --cached --others --exclude-standard -- spikes/infra-feasibility | Sort-Object -Unique | ForEach-Object {Join-Path $repoRoot $_})
$problems=@();$links=0;$json=0;$markdown=0
foreach($p in $files){
 if([IO.Path]::GetExtension($p) -eq '.json'){
  try{$null=Get-Content -Raw -LiteralPath $p | ConvertFrom-Json -Depth 100 -AsHashtable;$json++}catch{$problems+="Invalid JSON: $([IO.Path]::GetRelativePath($spikeRoot,$p))"}
 }
 if([IO.Path]::GetExtension($p) -ne '.md'){continue}
 $markdown++;$text=Get-Content -Raw -LiteralPath $p
 $fences=[regex]::Matches($text,'(?m)^```').Count;if($fences%2){$problems+="Unbalanced fences: $p"}
 foreach($m in [regex]::Matches($text,'\[[^\]]+\]\(([^)]+)\)')){
  $target=$m.Groups[1].Value
  if($target -match '^[a-z][a-z0-9+.-]*:|^#'){continue}
  $target=($target -split '#',2)[0]
  $resolved=[IO.Path]::GetFullPath((Join-Path (Split-Path $p) $target))
  $links++;if(-not(Test-Path -LiteralPath $resolved)){$problems+="Missing target: $([IO.Path]::GetRelativePath($spikeRoot,$p)) -> $target"}
 }
}
$matrix=Get-Content -Raw -LiteralPath (Join-Path $spikeRoot 'TEST_MATRIX.md')
foreach($id in 1..11){$needle='POC-{0:D2}' -f $id;if([regex]::Matches($matrix,"(?m)^\| $needle —").Count -ne 1){$problems+="Missing/duplicate POC row: $needle"}}
git -C $repoRoot diff --exit-code origin/main -- . ':!spikes/**' | Out-Null
if($LASTEXITCODE -ne 0){$problems+='Changes outside spike detected'}
git -C $repoRoot diff --check
if($LASTEXITCODE -ne 0){$problems+='Whitespace errors'}
$result=@{result=$(if($problems.Count){'FAIL'}else{'PASS'});files=$files.Count;markdown=$markdown;relativeLinks=$links;json=$json;pocRows=11;approvedFilesUnmodified=($problems -notcontains 'Changes outside spike detected');problems=$problems;limits='Relative file targets/fences/JSON/POC row IDs and tracked outside-scope diff; does not validate provider behavior'}
$result | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath (Join-Path $spikeRoot 'evidence/document-check.json') -Encoding utf8
$result | ConvertTo-Json -Depth 6
if($problems.Count){throw 'Documentary verification failed'}
