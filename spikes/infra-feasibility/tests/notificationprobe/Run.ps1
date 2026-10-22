param([string]$RunPattern='.',[string]$EvidenceName='integral-notifications',[switch]$Race)
$ErrorActionPreference='Stop'
if ($EvidenceName -notmatch '^integral-notifications(?:-[a-z0-9-]+)?$') { throw 'Invalid evidence basename' }
$spikeRoot=Split-Path (Split-Path $PSScriptRoot)
$testEnv=Join-Path $spikeRoot '.local/test.env'
if(-not(Test-Path -LiteralPath $testEnv)){throw 'Local spike test environment unavailable'}
$names=@('IDENTITY_DATABASE_URL','GAME_DATABASE_URL','SPIKE_USER_KEY','SPIKE_SERVICE_KEY')
$previous=@{}
try{
 foreach($name in ($names+@('GOWORK'))){$previous[$name]=[Environment]::GetEnvironmentVariable($name,'Process')}
 foreach($line in Get-Content -LiteralPath $testEnv){$parts=$line.Split('=',2);if($parts.Length -ne 2 -or $names -notcontains $parts[0]){throw 'Unexpected local environment entry'};[Environment]::SetEnvironmentVariable($parts[0],$parts[1],'Process')}
 $env:GOWORK='off'
 Push-Location $PSScriptRoot
 try{
  $goArgs=@('test','-tags=integration','-json','-count=1','-timeout=45s','-run',$RunPattern)
  if($Race){$goArgs+='-race'}
  $goArgs += './...'
  $output=@(go @goArgs 2>&1 | ForEach-Object { [string]$_ })
  $result=$LASTEXITCODE
  $output | Set-Content -Encoding utf8 (Join-Path $spikeRoot "evidence/$EvidenceName.jsonl")
  $cases=@(); $failures=@()
  foreach($line in $output){
   try{$entry=$line|ConvertFrom-Json}catch{continue}
   if($entry.Action -eq 'fail'){$failures+=$entry.Test}
   if($entry.Output -match 'CASE (\{.*\})'){$cases+=($Matches[1]|ConvertFrom-Json)}
  }
  [ordered]@{generatedUtc=[DateTime]::UtcNow.ToString('o');command=('go '+($goArgs -join ' '));exitCode=$result;postgresql='real local spike';broker='controlled test-only HTTP double; Ably real NOT_RUN';cases=$cases;failures=$failures} | ConvertTo-Json -Depth 16 | Set-Content -Encoding utf8 (Join-Path $spikeRoot "evidence/$EvidenceName-cases.json")
  Write-Output "notificationprobe exit=$result cases=$($cases.Count) failures=$($failures.Count)"
  if($result -ne 0){throw 'Notification verification failed; inspect redacted evidence'}
 }finally{Pop-Location}
}finally{foreach($name in $previous.Keys){[Environment]::SetEnvironmentVariable($name,$previous[$name],'Process')}}
