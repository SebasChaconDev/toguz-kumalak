$ErrorActionPreference='Stop'
$spikeRoot=Split-Path $PSScriptRoot
$testEnv=Join-Path $spikeRoot '.local/test.env'
if(-not(Test-Path -LiteralPath $testEnv)){throw 'Run infra/Setup-Local.ps1 first'}
$names=@('IDENTITY_DATABASE_URL','GAME_DATABASE_URL','SPIKE_USER_KEY','SPIKE_SERVICE_KEY')
$previous=@{}
try{
 foreach($n in ($names+@('GOWORK'))){$previous[$n]=[Environment]::GetEnvironmentVariable($n,'Process')}
 foreach($line in Get-Content -LiteralPath $testEnv){$parts=$line.Split('=',2);if($parts.Length -ne 2 -or $names -notcontains $parts[0]){throw 'Unexpected local environment entry'};[Environment]::SetEnvironmentVariable($parts[0],$parts[1],'Process')}
 Push-Location (Join-Path $spikeRoot 'services/game')
 try{$env:GOWORK='off';go test -tags=integration -count=1 -v ./app;if($LASTEXITCODE -ne 0){throw 'Integration tests did not pass'}}finally{Pop-Location}
}finally{foreach($n in $previous.Keys){[Environment]::SetEnvironmentVariable($n,$previous[$n],'Process')}}
