. (Join-Path $PSScriptRoot 'Remote-Control.ps1')

function New-SpikeManifest {
    $paths = @(& rg --files --hidden -g '!node_modules' -g '!.next' -g '!.vercel' -g '!evidence' -g '!tests' -g '!cmd' -g '!**/*_test.go' apps/web modules services)
    $allowed = @($paths | Where-Object { $_ -match '\.(go|mod|sum|json|mjs|ts|tsx|sql)$' })
    $files = @($allowed | Sort-Object | ForEach-Object {
        $bytes = [IO.File]::ReadAllBytes((Join-Path $script:SpikeRoot $_))
        [ordered]@{ path=($_ -replace '\\','/'); sha256=[Convert]::ToHexString([Security.Cryptography.SHA256]::HashData($bytes)).ToLowerInvariant(); bytes=$bytes.Length }
    })
    $digestInput = [Text.Encoding]::UTF8.GetBytes(($files | ConvertTo-Json -Depth 5 -Compress))
    $manifest = @{ sourceBase='94094620637b6705193ce984e8e5d54e39a3e45e'; sourceDigest=[Convert]::ToHexString([Security.Cryptography.SHA256]::HashData($digestInput)).ToLowerInvariant(); files=$files; scope='Disposable remote spike adapters; product unchanged' }
    $manifest | ConvertTo-Json -Depth 6 | Set-Content -LiteralPath (Join-Path $script:SpikeRoot 'evidence/remote-source-manifest-20261025.json')
    return $manifest
}

function Start-SpikeBuild([string]$Name, [switch]$Repeat) {
    $secrets=Get-SpikeSecrets
    $manifest=New-SpikeManifest
    $project=$secrets.projects[$Name]
    $root=if($Name -eq 'web'){'apps/web'}else{'services/'+$Name}
    $prefix='spikes/infra-feasibility/'
    $files=@($manifest.files | Where-Object { $_.path.StartsWith($root+'/') -or ($Name -ne 'web' -and $_.path.StartsWith('modules/')) } | ForEach-Object {
        @{file=$prefix+$_.path;data=[Convert]::ToBase64String([IO.File]::ReadAllBytes((Join-Path $script:SpikeRoot $_.path)));encoding='base64'}
    })
    $envs=@{SPIKE_REMOTE_ENABLED='1';SPIKE_SOURCE_REVISION=('9409462+'+$manifest.sourceDigest);GIN_MODE='release'}
    if($Name -ne 'web'){
        $envs.SPIKE_USER_KEY=$secrets.userKeys[$Name];$envs.SPIKE_SERVICE_KEY=$secrets.serviceKeys[$Name]
    }
    if($Name -in @('web','game')){
        $envs.SPIKE_IDENTITY_ORIGIN=$secrets.projects.identity.origin
        if($Name -eq 'web'){$envs.SPIKE_GAME_ORIGIN=$secrets.projects.game.origin}
        else{
            $envs.SPIKE_AI_ORIGIN=$secrets.projects.ai.origin
            $envs.SPIKE_IDENTITY_SERVICE_KEY=$secrets.serviceKeys.identity
            $envs.SPIKE_AI_SERVICE_KEY=$secrets.serviceKeys.ai
            $envs.SPIKE_ABLY_PUBLISH_KEY=$secrets.ably.keys['game-publisher'].key
        }
    }
    if($Name -eq 'identity'){$envs.IDENTITY_DATABASE_URL=$secrets.appDSNs.identity}
    if($Name -eq 'game'){$envs.GAME_DATABASE_URL=$secrets.appDSNs.game}
    if($Name -eq 'ai' -and @($envs.Keys | Where-Object {$_ -match 'DATABASE|GAME.*KEY'}).Count){throw 'AI isolation guard'}
    $body=@{name=$project.name;project=$project.id;files=$files;env=$envs;build=@{env=$envs};projectSettings=@{framework=$(if($Name -eq 'web'){'nextjs'}else{$null});rootDirectory=$prefix+$root;sourceFilesOutsideRootDirectory=$true};meta=@{spike='infra-feasibility-01';sourceDigest=$manifest.sourceDigest;repeatCold=[string][bool]$Repeat}}
    $j=Invoke-SpikeApi -Provider Vercel -Method POST -Path '/v13/deployments?teamId=team_xzyyrjATBH1J8bl5U89AiC59&forceNew=1&skipAutoDetectionConfirmation=1' -Body $body -Build -Purpose ('Protected remote build '+$Name+'; initial-production exception='+[string](-not $Repeat)+'; no secret output') -Summarize {param($x) @{id=$x.id;url=$x.url;target=$x.target;readyState=$x.readyState;projectId=$x.projectId;fileCount=$files.Count;environmentNames=@($envs.Keys);sourceDigest=$manifest.sourceDigest}}
    if(-not $script:RemoteLedger.deployments20261025){$script:RemoteLedger.deployments20261025=@()}
    $script:RemoteLedger.deployments20261025+=@{name=$Name;id=$j.id;url=$j.url;target=$j.target;repeat=[bool]$Repeat;sourceDigest=$manifest.sourceDigest;readyState=$j.readyState}
    Save-SpikeLedger
    $secrets.projects[$Name].deploymentId=$j.id;$secrets.projects[$Name].deploymentUrl='https://'+$j.url
    Save-SpikeSecrets $secrets
    if($Repeat -and $j.target -eq 'production'){throw 'Repeat unexpectedly targets production; stop for investigation'}
    Write-Output ($Name+' build accepted: '+$j.id+'; target='+$j.target+'; cumulative='+$script:RemoteLedger.remoteBuilds)
}

function Get-SpikeBuilds {
    foreach($item in $script:RemoteLedger.deployments20261025){
        if($item.readyState -in @('READY','ERROR','CANCELED')){continue}
        $j=Invoke-SpikeApi -Provider Vercel -Path ('/v13/deployments/'+$item.id+'?teamId=team_xzyyrjATBH1J8bl5U89AiC59') -Purpose ('Observe terminal build state '+$item.name) -Summarize {param($x) @{id=$x.id;readyState=$x.readyState;target=$x.target;errorCode=$x.errorCode;errorMessage=$x.errorMessage;alias=$x.alias;regions=$x.regions;createdAt=$x.createdAt;buildingAt=$x.buildingAt;ready=$x.ready}}
        $item.readyState=$j.readyState;$item.alias=$j.alias;$item.regions=$j.regions
        Save-SpikeLedger
        Write-Output ($item.name+': '+$j.readyState+'; '+$j.errorCode)
    }
}
