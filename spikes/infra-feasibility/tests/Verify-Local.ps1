param([switch]$ColdBuild)
$ErrorActionPreference='Stop'
$spikeRoot=Split-Path $PSScriptRoot
$evidence=Join-Path $spikeRoot 'evidence'
New-Item -ItemType Directory -Path $evidence -Force | Out-Null
$saved=@{}
foreach($n in @('GOWORK','GOENV','GOCACHE','GOMODCACHE','GOTOOLCHAIN','CGO_ENABLED')){$saved[$n]=[Environment]::GetEnvironmentVariable($n,'Process')}
$results=@()
try{
 $env:GOWORK='off';$env:GOENV='off';$env:GOTOOLCHAIN='local';$env:CGO_ENABLED='0'
 foreach($unit in @('modules/engine','modules/probe','services/identity','services/game','services/ai')){
  Push-Location (Join-Path $spikeRoot $unit)
  try{
   $slug=$unit.Replace('/','-')
   $start=[Diagnostics.Stopwatch]::StartNew()
   $output=go test -json -count=1 ./... 2>&1
   $exitCode=$LASTEXITCODE
   $output | Set-Content -LiteralPath (Join-Path $evidence "$slug-tests.jsonl") -Encoding utf8
   $start.Stop();$results+=@{unit=$unit;action='unit-tests';exit=$exitCode;seconds=$start.Elapsed.TotalSeconds}
   if($exitCode -ne 0){throw "Tests failed for $unit"}
  }finally{Pop-Location}
 }
 if($ColdBuild){
  foreach($service in @('identity','game','ai')){
   $ctxRoot=Join-Path $spikeRoot ".local/contexts/$service"
   if(Test-Path -LiteralPath $ctxRoot){throw 'Cold context exists; choose a fresh spike checkout or archive its evidence first.'}
   New-Item -ItemType Directory -Path "$ctxRoot/services","$ctxRoot/modules" -Force | Out-Null
   Copy-Item -LiteralPath (Join-Path $spikeRoot "services/$service") -Destination "$ctxRoot/services/$service" -Recurse
   Copy-Item -LiteralPath (Join-Path $spikeRoot 'modules/probe') -Destination "$ctxRoot/modules/probe" -Recurse
   if($service -ne 'identity'){Copy-Item -LiteralPath (Join-Path $spikeRoot 'modules/engine') -Destination "$ctxRoot/modules/engine" -Recurse}
   $env:GOCACHE=Join-Path $ctxRoot 'cache'
   $env:GOMODCACHE=Join-Path $ctxRoot 'modcache'
   Push-Location "$ctxRoot/services/$service"
   try{
    $start=[Diagnostics.Stopwatch]::StartNew()
    $log=go build -mod=readonly -trimpath -o "$ctxRoot/$service.exe" ./cmd/server 2>&1
    $exitCode=$LASTEXITCODE
    $log | Set-Content -LiteralPath (Join-Path $evidence "$service-cold-build.txt") -Encoding utf8
    $start.Stop();$results+=@{unit="services/$service";action='cold-build';exit=$exitCode;seconds=$start.Elapsed.TotalSeconds;work='off';cache='fresh-per-project';moduleCache='fresh-per-project'}
    if($exitCode -ne 0){throw "Cold build failed for $service"}
    $apiClock=[Diagnostics.Stopwatch]::StartNew()
    $apiLog=go build -mod=readonly ./api 2>&1
    $apiExit=$LASTEXITCODE
    $apiLog | Set-Content -LiteralPath (Join-Path $evidence "$service-api-build.txt") -Encoding utf8
    $apiClock.Stop();$results+=@{unit="services/$service";action='api-build';exit=$apiExit;seconds=$apiClock.Elapsed.TotalSeconds;work='off'}
    if($apiExit -ne 0){throw "API build failed for $service"}
    go version -m "$ctxRoot/$service.exe" | Set-Content -LiteralPath (Join-Path $evidence "$service-build-modules.txt") -Encoding utf8
    $engine=if($service -eq 'identity'){'none'}else{Get-FileHash "$ctxRoot/modules/engine/engine.go" -Algorithm SHA256 | Select-Object -ExpandProperty Hash}
    $deps=go list -m -json all
    if($LASTEXITCODE -ne 0){throw 'Module inventory failed'}
    $deps | Set-Content -LiteralPath (Join-Path $evidence "$service-module-inventory.jsonl") -Encoding utf8
    $results+=@{unit="services/$service";action='shared-engine-source';sha256=$engine;engineVersion=$(if($service -eq 'identity'){'none'}else{'v0.1.0 (local replace; no release tag)'})}
   }finally{Pop-Location}
  }
 }
}finally{
 $resultFile=if($ColdBuild){'local-verification.json'}else{'unit-verification.json'}
 $results | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath (Join-Path $evidence $resultFile) -Encoding utf8
 foreach($n in $saved.Keys){[Environment]::SetEnvironmentVariable($n,$saved[$n],'Process')}
}
$results | ConvertTo-Json -Depth 6
