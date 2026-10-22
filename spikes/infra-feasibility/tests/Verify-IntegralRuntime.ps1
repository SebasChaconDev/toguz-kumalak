param([switch]$SkipDatabaseRestart)
$ErrorActionPreference='Stop'
$spikeRoot=Split-Path $PSScriptRoot
$runId='integral-'+[Guid]::NewGuid().ToString('N')
$privateDir=Join-Path $spikeRoot ".local/integral-runtime/$runId"
$publicDir=Join-Path $spikeRoot "evidence/integral-runtime/$runId"
New-Item -ItemType Directory -Path $privateDir,$publicDir -Force | Out-Null
$values=@{}
foreach($line in Get-Content -LiteralPath (Join-Path $spikeRoot '.local/test.env')){$pair=$line.Split('=',2);$values[$pair[0]]=$pair[1]}
$secrets=@($values.Values)
foreach($line in Get-Content -LiteralPath (Join-Path $spikeRoot '.local/compose.env')){$pair=$line.Split('=',2);if($pair.Length -eq 2){$secrets+=$pair[1]}}
foreach($name in @('IDENTITY_DATABASE_URL','GAME_DATABASE_URL')){
 $uri=[uri]$values[$name]
 if($uri.Host -ne '127.0.0.1' -or $uri.Port -ne 15435 -or $uri.AbsolutePath.Trim('/') -notin @('identity','game')){throw 'Unexpected database target; this runner is local-only'}
}
$saved=@{}
foreach($name in ($values.Keys+@('GOWORK','SPIKE_RUN_ID'))){$saved[$name]=[Environment]::GetEnvironmentVariable($name,'Process')}
$children=@{};$childLogs=@();$summary=@{run=$runId;startedUTC=[DateTime]::UtcNow.ToString('o');result='FAIL';environment='Windows local processes and Docker PostgreSQL';cases=@();databaseRestart='NOT_RUN'}
function Build-Driver($root,$output,$target){
 Push-Location $root
 try{go build -mod=readonly -o $output $target;if($LASTEXITCODE -ne 0){throw "Build failed in $([IO.Path]::GetRelativePath($spikeRoot,$root))"}}finally{Pop-Location}
}
function Start-Probe($service){
 $psi=[Diagnostics.ProcessStartInfo]::new((Join-Path $privateDir "$service.exe"))
 $psi.UseShellExecute=$false;$psi.CreateNoWindow=$true;$psi.WindowStyle='Hidden';$psi.RedirectStandardOutput=$true;$psi.RedirectStandardError=$true
 $psi.Environment.Clear()
 foreach($n in @('SystemRoot','WINDIR','TEMP','TMP','PATH')){$value=[Environment]::GetEnvironmentVariable($n,'Process');if($null -ne $value){$psi.Environment[$n]=$value}}
 $mac=[Security.Cryptography.HMACSHA256]::new([Convert]::FromBase64String($values['SPIKE_USER_KEY']))
 $userKey=[Convert]::ToBase64String($mac.ComputeHash([Text.Encoding]::UTF8.GetBytes('spike-user:'+ $service)));$mac.Dispose()
 $script:secrets+=$userKey
 $psi.Environment['SPIKE_USER_KEY']=$userKey;$psi.Environment['SPIKE_SERVICE_KEY']=$values['SPIKE_SERVICE_KEY']
 if($service -ne 'ai'){$name=$service.ToUpperInvariant()+'_DATABASE_URL';$psi.Environment[$name]=$values[$name]}
 $p=[Diagnostics.Process]::new();$p.StartInfo=$psi;$null=$p.Start()
 $children[$service]=@{process=$p;stdout=$p.StandardOutput.ReadToEndAsync();stderr=$p.StandardError.ReadToEndAsync();pid=$p.Id}
 $port=@{identity=18101;game=18102;ai=18103}[$service]
 $ready=$false
 for($attempt=0;$attempt -lt 100;$attempt++){
  if($p.HasExited){throw "Experimental $service process exited before readiness"}
  try{$tcp=[Net.Sockets.TcpClient]::new();$tcp.Connect('127.0.0.1',$port);$tcp.Dispose();$ready=$true;break}catch{Start-Sleep -Milliseconds 100}
 }
 if(-not $ready){throw "Experimental $service process not listening"}
 return $p.Id
}
function Stop-Probe($service){
 if(-not $children.ContainsKey($service)){return}
 $child=$children[$service];$p=$child.process
 if(-not $p.HasExited){$p.Kill();if(-not $p.WaitForExit(5000)){throw 'Owned child did not exit'}}
 $text=$child.stdout.GetAwaiter().GetResult()+$child.stderr.GetAwaiter().GetResult()
 $leaked=$false;foreach($secret in $secrets){if($secret.Length -gt 16 -and $text.Contains($secret)){$leaked=$true}}
 $text | Set-Content -LiteralPath (Join-Path $privateDir "$service-$($child.pid).log")
 $script:childLogs+=@{service=$service;pid=$child.pid;secretScan=$(if($leaked){'FAIL'}else{'PASS'});bytes=$text.Length}
 $children.Remove($service)
 if($leaked){throw 'Generated secret in service log; value withheld'}
}
function Invoke-Driver($mode,$label=$mode){
 $clock=[Diagnostics.Stopwatch]::StartNew()
 $raw=& (Join-Path $privateDir 'driver.exe') -mode $mode 2>&1
 $exitCode=$LASTEXITCODE;$clock.Stop()
 $rawText=($raw | ForEach-Object {$_.ToString()}) -join "`n"
 foreach($secret in $secrets){if($secret.Length -gt 16 -and $rawText.Contains($secret)){throw 'Generated secret in driver output; value withheld'}}
 $rawText | Set-Content -LiteralPath (Join-Path $publicDir "$label.json")
 if($exitCode -ne 0){throw "Process verification $mode failed; inspect sanitized evidence"}
 $result=$rawText | ConvertFrom-Json -AsHashtable
 $script:summary.cases+=@{mode=$mode;result=$result.result;exit=$exitCode;seconds=$clock.Elapsed.TotalSeconds}
 return $result
}
try{
 foreach($n in $values.Keys){[Environment]::SetEnvironmentVariable($n,$values[$n],'Process')};$env:GOWORK='off';$env:SPIKE_RUN_ID=$runId
 foreach($port in @(18101,18102,18103)){if(Get-NetTCPConnection -LocalPort $port -State Listen -ErrorAction SilentlyContinue){throw "Port $port already in use; do not stop unknown process"}}
 Push-Location (Join-Path $PSScriptRoot 'processprobe')
 try{go mod tidy;if($LASTEXITCODE -ne 0){throw 'Driver dependency resolution failed'}}finally{Pop-Location}
 Build-Driver (Join-Path $PSScriptRoot 'processprobe') (Join-Path $privateDir 'driver.exe') '.'
 foreach($service in @('identity','game','ai')){Build-Driver (Join-Path $spikeRoot "services/$service") (Join-Path $privateDir "$service.exe") './cmd/server'}
 $null=Invoke-Driver 'prepare'
 foreach($service in @('identity','game','ai')){$null=Start-Probe $service}
 $null=Invoke-Driver 'transactions';$null=Invoke-Driver 'security'
 $oldPID=$children['game'].pid;Stop-Probe 'game';$newPID=Start-Probe 'game'
 if($oldPID -eq $newPID){throw 'Game process did not change'}
 $null=Invoke-Driver 'retry';$summary.gameRestart=@{beforePID=$oldPID;afterPID=$newPID;result='PASS'}
 Stop-Probe 'ai';$null=Invoke-Driver 'internal-down';$null=Start-Probe 'ai';$null=Invoke-Driver 'internal-up'
 if(-not $SkipDatabaseRestart){
  $before=Invoke-Driver 'snapshot' 'snapshot-before'
  $project=(docker inspect toguz-infra-spike-postgres-1 --format '{{index .Config.Labels "com.docker.compose.project"}}').Trim()
  if($project -ne 'toguz-infra-spike'){throw 'Unexpected container owner'}
  $volume=(docker inspect toguz-infra-spike-postgres-1 --format '{{range .Mounts}}{{if eq .Type "volume"}}{{.Name}}{{end}}{{end}}').Trim()
  if($volume -ne 'toguz-infra-spike_spike_pg'){throw 'Unexpected or missing synthetic volume'}
  docker compose --env-file (Join-Path $spikeRoot '.local/compose.env') -p toguz-infra-spike -f (Join-Path $spikeRoot 'infra/compose.yaml') restart -t 10 postgres
  if($LASTEXITCODE -ne 0){throw 'Dedicated PostgreSQL restart failed'}
  docker compose --env-file (Join-Path $spikeRoot '.local/compose.env') -p toguz-infra-spike -f (Join-Path $spikeRoot 'infra/compose.yaml') up -d --wait --wait-timeout 120
  if($LASTEXITCODE -ne 0){throw 'PostgreSQL not healthy after restart'}
  $after=Invoke-Driver 'snapshot' 'snapshot-after'
  if(($before.cases[0].after | ConvertTo-Json -Compress) -ne ($after.cases[0].after | ConvertTo-Json -Compress)){throw 'State lost across PostgreSQL restart'}
  $volumeAfter=(docker inspect toguz-infra-spike-postgres-1 --format '{{range .Mounts}}{{if eq .Type "volume"}}{{.Name}}{{end}}{{end}}').Trim()
  if($volume -ne $volumeAfter){throw 'Volume changed'}
  $summary.databaseRestart=@{result='PASS';volumePreserved=$true;before=$before.cases[0].after;after=$after.cases[0].after}
  $null=Invoke-Driver 'retry'
 }
 $summary.aiEnvironment=@{databaseCredentialsPassed=$false;environmentAllowlist=$true;userKeyScopedToAudience=$true;authenticatedConnectivity='PASS'}
 $summary.result='PASS'
}finally{
 $cleanupErrors=@()
 foreach($service in @($children.Keys)){try{Stop-Probe $service}catch{$cleanupErrors+='Owned child cleanup or secret scan failed'}}
 foreach($n in $saved.Keys){try{[Environment]::SetEnvironmentVariable($n,$saved[$n],'Process')}catch{$cleanupErrors+='Environment restoration failed'}}
 if($cleanupErrors.Count){$summary.result='FAIL';$summary.cleanupErrors=$cleanupErrors}
 $summary.logScan=$childLogs;$summary.finishedUTC=[DateTime]::UtcNow.ToString('o')
 $summary | ConvertTo-Json -Depth 12 | Set-Content -LiteralPath (Join-Path $publicDir 'summary.json')
 if($cleanupErrors.Count){throw 'Cleanup/secret scan failure; all remaining cleanup steps were attempted'}
}
$summary | ConvertTo-Json -Depth 12
