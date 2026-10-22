$ErrorActionPreference='Stop'
$spikeRoot=Split-Path $PSScriptRoot
$repoRoot=(git -C $spikeRoot rev-parse --show-toplevel).Trim()
$secrets=@()
foreach($envFile in @('.local/compose.env','.local/test.env')){
 $p=Join-Path $spikeRoot $envFile
 if(Test-Path -LiteralPath $p){foreach($line in Get-Content -LiteralPath $p){$parts=$line.Split('=',2);if($parts.Length -eq 2 -and $parts[1].Length -ge 16){$secrets+=$parts[1]}}}
}
if($secrets.Count -lt 7){throw 'Expected local synthetic credentials unavailable; exposure check would be incomplete'}
$paths=@(git -C $repoRoot ls-files --cached --others --exclude-standard -- spikes/infra-feasibility | Sort-Object -Unique | ForEach-Object {Join-Path $repoRoot $_})
$binaries=@(Get-ChildItem -LiteralPath (Join-Path $spikeRoot '.local/contexts') -Filter '*.exe' -Recurse | Where-Object {$_.Directory.Name -in @('identity','game','ai')} | Select-Object -ExpandProperty FullName)
$findings=@()
foreach($path in ($paths+$binaries)){
 $content=[Text.Encoding]::UTF8.GetString([IO.File]::ReadAllBytes($path))
 foreach($secret in $secrets){if($content.Contains($secret)){$findings+=@{path=[IO.Path]::GetRelativePath($spikeRoot,$path);kind='local-generated-secret'}}}
 if($content -match '-----BEGIN (?:RSA |EC |OPENSSH )?PRIVATE KEY-----|\bgh[pousr]_[A-Za-z0-9]{30,}\b|\bAKIA[A-Z0-9]{16}\b'){$findings+=@{path=[IO.Path]::GetRelativePath($spikeRoot,$path);kind='credential-pattern'}}
}
$result=@{result=$(if($findings.Count){'FAIL'}else{'PASS'});publicFiles=$paths.Count;builtExecutables=$binaries.Count;generatedValuesCompared=$secrets.Count;findings=$findings;limits='Known synthetic values and selected patterns only; no universal vulnerability assurance; web ephemeral fixtures checked by smoke'}
$result | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath (Join-Path $spikeRoot 'evidence/exposure-check.json') -Encoding utf8
$result | ConvertTo-Json -Depth 6
if($findings.Count){throw 'Exposure check failed; secret values withheld'}
