param([switch]$GenerateOnly)
$ErrorActionPreference='Stop'
$spikeRoot=Split-Path $PSScriptRoot
$privateDir=Join-Path $spikeRoot '.local'
New-Item -ItemType Directory -Path $privateDir -Force | Out-Null
$envFile=Join-Path $privateDir 'compose.env'
if(Test-Path -LiteralPath $envFile){throw 'Local credentials already exist; preserve them or retire this dedicated spike first.'}
function New-SyntheticSecret { [Convert]::ToHexString([Security.Cryptography.RandomNumberGenerator]::GetBytes(32)).ToLowerInvariant() }
$adminPass=New-SyntheticSecret
$identityPass=New-SyntheticSecret
$gamePass=New-SyntheticSecret
@("ADMIN_PASSWORD=$adminPass","IDENTITY_PASSWORD=$identityPass","GAME_PASSWORD=$gamePass") | Set-Content -LiteralPath $envFile -Encoding utf8
# Paths ignored by the spike's .gitignore. Values never printed.
@("IDENTITY_DATABASE_URL=postgres://identity_owner:${identityPass}@127.0.0.1:15435/identity?sslmode=disable","GAME_DATABASE_URL=postgres://game_owner:${gamePass}@127.0.0.1:15435/game?sslmode=disable","SPIKE_USER_KEY=$([Convert]::ToBase64String([Security.Cryptography.RandomNumberGenerator]::GetBytes(32)))","SPIKE_SERVICE_KEY=$([Convert]::ToBase64String([Security.Cryptography.RandomNumberGenerator]::GetBytes(32)))") | Set-Content -LiteralPath (Join-Path $privateDir 'test.env') -Encoding utf8
docker compose --env-file $envFile -f (Join-Path $PSScriptRoot 'compose.yaml') config --quiet
if($LASTEXITCODE -ne 0){throw 'Compose config rejected'}
if(-not $GenerateOnly){
 docker compose --env-file $envFile -f (Join-Path $PSScriptRoot 'compose.yaml') up -d --wait
 if($LASTEXITCODE -ne 0){throw 'Local PostgreSQL startup failed'}
}
Write-Output 'Local synthetic credentials generated; Compose config valid. Values not logged.'
