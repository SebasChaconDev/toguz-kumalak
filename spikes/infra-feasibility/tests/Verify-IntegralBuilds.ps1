param([ValidateSet('Web','Go','LinuxRace','NotificationRace','All')][string]$Phase='All')
$ErrorActionPreference='Stop'
$PSNativeCommandUseErrorActionPreference=$false
$spikeRoot=Split-Path $PSScriptRoot
$runId=[DateTime]::UtcNow.ToString('yyyyMMddTHHmmssZ')+'-'+[guid]::NewGuid().ToString('N')
$privateRoot=Join-Path $spikeRoot ".local/integral-builds/$runId"
$evidence=Join-Path $spikeRoot "evidence/integral-builds/$runId"
New-Item -ItemType Directory -Path $privateRoot,$evidence | Out-Null
$results=[Collections.Generic.List[object]]::new()
$envNames=@('GOWORK','GOENV','GOCACHE','GOMODCACHE','GOTOOLCHAIN','CGO_ENABLED','NEXT_TELEMETRY_DISABLED','npm_config_cache','SPIKE_IDENTITY_ORIGIN','SPIKE_GAME_ORIGIN')
$saved=@{};foreach($name in $envNames){$saved[$name]=[Environment]::GetEnvironmentVariable($name,'Process')}
function Invoke-Observed([string]$Label,[string]$Directory,[string]$Program,[string[]]$Arguments){
 $clock=[Diagnostics.Stopwatch]::StartNew();Push-Location -LiteralPath $Directory
 try{$output=& $Program @Arguments 2>&1;$code=$LASTEXITCODE}finally{Pop-Location;$clock.Stop()}
 $text=($output | Out-String).Replace($spikeRoot,'<SPIKE_ROOT>').Replace($privateRoot,'<PRIVATE_CONTEXT>')
 [IO.File]::WriteAllText((Join-Path $evidence "$Label.txt"),$text)
 $commandText=((@($Program)+$Arguments)-join ' ').Replace($spikeRoot,'<SPIKE_ROOT>')
 $results.Add([ordered]@{label=$Label;command=$commandText;directory=$Directory.Replace($spikeRoot,'<SPIKE_ROOT>');exit=$code;seconds=[math]::Round($clock.Elapsed.TotalSeconds,3);log="$Label.txt"})
 Write-Host "$Label exit=$code seconds=$([math]::Round($clock.Elapsed.TotalSeconds,2))"
 if($code -ne 0){throw "$Label failed (exit $code); see evidence"}
 return $output
}
function Copy-Module([string]$Source,[string]$Target){
 New-Item -ItemType Directory -Path $Target -Force | Out-Null
 Get-ChildItem -LiteralPath $Source -File | Where-Object {$_.Extension -in @('.go','.sql') -or $_.Name -in @('go.mod','go.sum')} | Copy-Item -Destination $Target
 Get-ChildItem -LiteralPath $Source -Directory | Where-Object {$_.Name -notin @('.local','cache','modcache')} | ForEach-Object {Copy-Module $_.FullName (Join-Path $Target $_.Name)}
}
try{
 $env:GOWORK='off';$env:GOENV='off';$env:GOTOOLCHAIN='local';$env:CGO_ENABLED='0';$env:NEXT_TELEMETRY_DISABLED='1'
 $env:SPIKE_IDENTITY_ORIGIN=$null;$env:SPIKE_GAME_ORIGIN=$null
 Invoke-Observed 'git-head' $spikeRoot git @('rev-parse','HEAD') | Out-Null
 Invoke-Observed 'go-version' $spikeRoot go @('version') | Out-Null
 Invoke-Observed 'node-version' $spikeRoot node @('--version') | Out-Null
 Invoke-Observed 'npm-version' $spikeRoot npm.cmd @('--version') | Out-Null
 if($Phase -in @('Web','All')){
  $web=Join-Path $privateRoot 'web';New-Item -ItemType Directory -Path "$web/evidence" -Force | Out-Null
  foreach($dir in @('app','lib','tests')){Copy-Item -LiteralPath (Join-Path $spikeRoot "apps/web/$dir") -Destination $web -Recurse}
  Get-ChildItem -LiteralPath (Join-Path $spikeRoot 'apps/web') -File | Copy-Item -Destination $web
  $env:npm_config_cache=Join-Path $privateRoot 'npm-cache'
  Invoke-Observed 'web-ci' $web npm.cmd @('ci','--no-audit','--no-fund') | Out-Null
  Invoke-Observed 'web-tests' $web npm.cmd @('test') | Out-Null
  $package=Get-Content -LiteralPath "$web/package.json" -Raw | ConvertFrom-Json
  if($package.scripts.PSObject.Properties.Name -contains 'lint'){Invoke-Observed 'web-lint' $web npm.cmd @('run','lint') | Out-Null}
  else{$results.Add(@{label='web-lint';status='NOT_DEFINED';reason='No lint script in original package.json'})}
  Invoke-Observed 'web-build' $web npm.cmd @('run','build') | Out-Null
  Invoke-Observed 'web-smoke' $web npm.cmd @('run','smoke') | Out-Null
  Copy-Item -LiteralPath "$web/evidence/web-smoke.json" -Destination (Join-Path $evidence 'web-smoke.json')
  $lockSource=(Get-FileHash -LiteralPath (Join-Path $spikeRoot 'apps/web/package-lock.json')).Hash
  $lockCopy=(Get-FileHash -LiteralPath "$web/package-lock.json").Hash
  if($lockSource -ne $lockCopy){throw 'npm ci modified lockfile'}
  $results.Add(@{label='web-lockfile';sha256=$lockSource;unchanged=$true;context='Fresh source copy, fresh node_modules and npm cache'})
  Write-Host 'WEB_SMOKE_RELEASED_18100_18101_18102'
 }
 if($Phase -in @('Go','All')){
  $nativeCompiler=Get-Command gcc -ErrorAction SilentlyContinue
  $results.Add(@{label='windows-native-race';status=$(if($nativeCompiler){'AVAILABLE'}else{'NOT_RUN'});reason=$(if($nativeCompiler){'gcc located; run native race per module'}else{'No gcc executable resolved in Windows PATH; race detector executed separately in Linux with CGO_ENABLED=1'})})
  $env:GOCACHE=Join-Path $privateRoot 'go-cache';$env:GOMODCACHE=Join-Path $privateRoot 'go-modcache'
  foreach($unit in @('modules/engine','modules/probe','services/identity','services/game','services/ai')){
   $dir=Join-Path $spikeRoot $unit;$slug=$unit.Replace('/','-')
   Invoke-Observed "$slug-env" $dir go @('env','GOVERSION','GOOS','GOARCH','GOWORK','GOMOD','GOCACHE','GOMODCACHE','CGO_ENABLED') | Out-Null
   Invoke-Observed "$slug-test" $dir go @('test','-mod=readonly','-count=1','./...') | Out-Null
   Invoke-Observed "$slug-vet" $dir go @('vet','-mod=readonly','./...') | Out-Null
   Invoke-Observed "$slug-build" $dir go @('build','-mod=readonly','./...') | Out-Null
   if($nativeCompiler){try{$env:CGO_ENABLED='1';Invoke-Observed "$slug-native-race" $dir go @('test','-mod=readonly','-race','-count=1','./...') | Out-Null}finally{$env:CGO_ENABLED='0'}}
  }
  foreach($service in @('identity','game','ai')){
   $context=Join-Path $privateRoot "isolated-$service"
   Copy-Module (Join-Path $spikeRoot "services/$service") "$context/services/$service"
   Copy-Module (Join-Path $spikeRoot 'modules/probe') "$context/modules/probe"
   if($service -ne 'identity'){Copy-Module (Join-Path $spikeRoot 'modules/engine') "$context/modules/engine"}
   $env:GOCACHE=Join-Path $context 'cache';$env:GOMODCACHE=Join-Path $context 'modcache'
   if((Test-Path $env:GOCACHE) -or (Test-Path $env:GOMODCACHE)){throw 'Expected empty isolated caches'}
   $dir="$context/services/$service"
   Invoke-Observed "$service-isolated-env" $dir go @('env','GOWORK','GOMOD','GOCACHE','GOMODCACHE') | Out-Null
   Invoke-Observed "$service-isolated-build" $dir go @('build','-mod=readonly','-trimpath','-o',"$context/$service.exe",'./cmd/server') | Out-Null
   Invoke-Observed "$service-isolated-api" $dir go @('build','-mod=readonly','./api') | Out-Null
   Invoke-Observed "$service-isolated-modules" $dir go @('list','-m','-json','all') | Out-Null
   Invoke-Observed "$service-binary-modules" $dir go @('version','-m',"$context/$service.exe") | Out-Null
   $deps=Invoke-Observed "$service-dependencies" $dir go @('list','-deps','./...')
   $engineDep=@($deps | Where-Object {$_ -match '/modules/engine$'}).Count -gt 0
   if(($service -eq 'identity') -eq $engineDep){throw 'Unexpected shared engine dependency boundary'}
   $hash=if($service -eq 'identity'){$null}else{(Get-FileHash -LiteralPath "$context/modules/engine/engine.go").Hash}
   $results.Add(@{label="$service-isolation";gowork='off';initialCaches='Absent before build';includesEngine=$engineDep;engineSha256=$hash;engineModuleVersion=$(if($engineDep){'v0.1.0 local replacement'}else{'none'});context='Only service, probe and engine when required; no repository root'})
  }
 }
 if($Phase -in @('LinuxRace','All')){
  $image='golang:1.26.2-bookworm'
  Invoke-Observed 'docker-go-pull' $spikeRoot docker @('pull',$image) | Out-Null
  Invoke-Observed 'docker-go-image' $spikeRoot docker @('image','inspect',$image,'--format','{{json .RepoDigests}} {{.Id}}') | Out-Null
  $shell='set -eu; go version; gcc --version | head -n 1; for unit in modules/engine modules/probe services/identity services/game services/ai; do echo "BEGIN:$unit"; cd "/spike/$unit"; go test -mod=readonly -race -count=1 ./...; echo "PASS:$unit"; done'
  Invoke-Observed 'linux-race' $spikeRoot docker @('run','--rm','--name',"toguz-integral-race-$runId",'--cpus','2','--memory','4g','--network','bridge','--mount',"type=bind,source=$spikeRoot/modules,target=/spike/modules,readonly",'--mount',"type=bind,source=$spikeRoot/services,target=/spike/services,readonly",'--env','GOWORK=off','--env','GOENV=off','--env','GOTOOLCHAIN=local','--env','CGO_ENABLED=1','--env','GOCACHE=/tmp/go-cache','--env','GOMODCACHE=/tmp/go-modcache',$image,'bash','-c',$shell) | Out-Null
 }
 if($Phase -eq 'NotificationRace'){
  $dockerEnv=Join-Path $privateRoot 'notification-container.env'
  $allowed=@('IDENTITY_DATABASE_URL','GAME_DATABASE_URL','SPIKE_USER_KEY','SPIKE_SERVICE_KEY')
  $envLines=foreach($line in Get-Content -LiteralPath (Join-Path $spikeRoot '.local/test.env')){
   $parts=$line.Split('=',2)
   if($parts.Length -ne 2 -or $parts[0] -notin $allowed){throw 'Unexpected synthetic environment entry'}
   if($parts[0] -in @('IDENTITY_DATABASE_URL','GAME_DATABASE_URL')){
    $uri=[uri]$parts[1]
    if($uri.Host -ne '127.0.0.1' -or $uri.Port -ne 15435){throw 'Expected dedicated loopback PostgreSQL endpoint'}
    $parts[1]=$parts[1].Replace('127.0.0.1:15435','postgres:5432')
   }
   $parts -join '='
  }
  [IO.File]::WriteAllLines($dockerEnv,[string[]]$envLines)
  $image='golang:1.26.2-bookworm'
  Invoke-Observed 'docker-go-image' $spikeRoot docker @('image','inspect',$image,'--format','{{json .RepoDigests}} {{.Id}}') | Out-Null
  $shell='set -eu; go version; gcc --version | head -n 1; cd /spike/tests/notificationprobe; go test -mod=readonly -tags=integration -race -count=1 -json ./...'
  Invoke-Observed 'notification-linux-race' $spikeRoot docker @('run','--rm','--name',"toguz-integral-notify-$runId",'--cpus','2','--memory','4g','--network','toguz-infra-spike_default','--mount',"type=bind,source=$spikeRoot/modules,target=/spike/modules,readonly",'--mount',"type=bind,source=$spikeRoot/services,target=/spike/services,readonly",'--mount',"type=bind,source=$spikeRoot/tests/notificationprobe,target=/spike/tests/notificationprobe,readonly",'--env-file',$dockerEnv,'--env','GOWORK=off','--env','GOENV=off','--env','GOTOOLCHAIN=local','--env','CGO_ENABLED=1','--env','GOCACHE=/tmp/go-cache','--env','GOMODCACHE=/tmp/go-modcache',$image,'bash','-c',$shell) | Out-Null
 }
}finally{
 $results | ConvertTo-Json -Depth 8 | Set-Content -LiteralPath (Join-Path $evidence 'results.json') -Encoding utf8
 foreach($name in $saved.Keys){[Environment]::SetEnvironmentVariable($name,$saved[$name],'Process')}
 Write-Host "EVIDENCE=$evidence"
}
