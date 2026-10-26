# Throwaway spike operator. No product implementation or unmetered SDK retries.
$ErrorActionPreference = 'Stop'
$script:SpikeRoot = Split-Path $PSScriptRoot
$script:LedgerPath = Join-Path $script:SpikeRoot 'evidence/remote-preconditions-budget-20261022.json'
$script:SecretPath = Join-Path $script:SpikeRoot '.local/remote-20261025.clixml'
$script:RemoteLedger = Get-Content -LiteralPath $script:LedgerPath -Raw | ConvertFrom-Json -AsHashtable -DateKind String -Depth 100

function Save-SpikeLedger {
    $script:RemoteLedger.counters.newOperationalRequests = $script:RemoteLedger.requests.Count
    $script:RemoteLedger.counters.cumulativeOperationalRequests = 1 + $script:RemoteLedger.requests.Count
    $script:RemoteLedger.counters.directApiRequestsThroughCurrentRecord = 1 + $script:RemoteLedger.requests.Count
    $script:RemoteLedger.observedUTC = [DateTimeOffset]::UtcNow.ToString('o')
    $script:RemoteLedger | ConvertTo-Json -Depth 100 | Set-Content -LiteralPath $script:LedgerPath -Encoding utf8
}

function Get-SpikeSecrets {
    if (-not (Test-Path -LiteralPath $script:SecretPath)) { return @{} }
    $credential = Import-Clixml -LiteralPath $script:SecretPath
    return ConvertFrom-Json -InputObject $credential.GetNetworkCredential().Password -AsHashtable -DateKind String -Depth 40
}

function Save-SpikeSecrets($Values) {
    # Windows DPAPI encryption, decryptable only by this Windows user on this host.
    $secure = ConvertTo-SecureString -String ($Values | ConvertTo-Json -Depth 40 -Compress) -AsPlainText -Force
    [PSCredential]::new('toguz-disposable-spike', $secure) | Export-Clixml -LiteralPath $script:SecretPath
}

function Get-AblyOwnerToken {
    $sessionRoot = if ($env:CODEX_HOME) { Join-Path $env:CODEX_HOME 'sessions' } else { Join-Path $env:USERPROFILE '.codex/sessions' }
    $paths = @(rg --files --hidden $sessionRoot -g ('*' + $env:CODEX_THREAD_ID + '*.jsonl'))
    if ($paths.Count -ne 1) { throw 'Cannot identify this chat credential source' }
    $token = $null
    $stream = [IO.FileStream]::new($paths[0], [IO.FileMode]::Open, [IO.FileAccess]::Read, [IO.FileShare]::ReadWrite)
    $reader = [IO.StreamReader]::new($stream)
    try {
        while ($null -ne ($line = $reader.ReadLine())) {
            try { $item = $line | ConvertFrom-Json -AsHashtable -Depth 100 } catch { continue }
            $texts = @()
            if ($item.type -eq 'event_msg' -and $item.payload.type -eq 'user_message') { $texts += [string]$item.payload.message }
            if ($item.type -eq 'response_item' -and $item.payload.role -eq 'user') { $texts += @($item.payload.content | ForEach-Object { $_.text }) }
            foreach ($text in $texts) {
                foreach ($m in [regex]::Matches([string]$text, '(?m)^(eyJ[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+)[ \t]*$')) { $token = $m.Groups[1].Value }
            }
        }
    } finally { $reader.Dispose(); $stream.Dispose() }
    if (-not $token) { throw 'Owner Control API token unavailable; never substitute an app key' }
    return $token
}

function Invoke-SpikeApi {
    param(
        [ValidateSet('Vercel','Neon','Ably','GitHub')][string]$Provider,
        [ValidateSet('GET','POST','PATCH','DELETE')][string]$Method = 'GET',
        [string]$Path,
        $Body = $null,
        [string]$Purpose,
        [scriptblock]$Summarize = { param($json) @{} },
        [switch]$Build,
        [switch]$RecordNonsecretError
    )
    $base = switch ($Provider) { Vercel { 'https://api.vercel.com' } Neon { 'https://console.neon.tech/api/v2' } Ably { 'https://control.ably.net/v1' } GitHub { 'https://api.github.com' } }
    $token = switch ($Provider) {
        Vercel { (Get-Content -LiteralPath 'C:/Users/swmfi/AppData/Roaming/com.vercel.cli/Data/auth.json' -Raw | ConvertFrom-Json -AsHashtable).token }
        Neon { (Get-Content -LiteralPath 'C:/Users/swmfi/.config/neonctl/credentials.json' -Raw | ConvertFrom-Json -AsHashtable).access_token }
        Ably { Get-AblyOwnerToken }
        GitHub { & gh auth token --hostname github.com }
    }
    if (-not $token -or $Path -notmatch '^/' -or $Path -match '://|[\r\n]') { throw 'Invalid request configuration' }
    if ($Build -and $script:RemoteLedger.remoteBuilds -ge $script:RemoteLedger.limits.builds) { throw 'Cumulative build ceiling reached' }
    $entry = [ordered]@{
        id = 'REMOTE-{0:D3}' -f (2 + $script:RemoteLedger.requests.Count)
        provider = $Provider; method = $Method; host = ([Uri]$base).Host; endpoint = $Path
        purpose = $Purpose; readOnly = ($Method -eq 'GET'); startedUTC = [DateTimeOffset]::UtcNow.ToString('o')
        httpStatus = $null; automaticRetries = 0; followRedirects = $false
    }
    $script:RemoteLedger.requests += $entry
    if ($Method -ne 'GET') { $script:RemoteLedger.externalWrites++; $script:RemoteLedger.counters.externalWrites++ }
    if ($Build) { $script:RemoteLedger.remoteBuilds++; $script:RemoteLedger.counters.remoteBuilds++ }
    Save-SpikeLedger
    $handler = [Net.Http.HttpClientHandler]::new(); $handler.AllowAutoRedirect = $false
    $client = [Net.Http.HttpClient]::new($handler); $client.Timeout = [TimeSpan]::FromSeconds(45)
    $request = [Net.Http.HttpRequestMessage]::new([Net.Http.HttpMethod]::new($Method), $base + $Path)
    $request.Headers.Authorization = [Net.Http.Headers.AuthenticationHeaderValue]::new('Bearer', [string]$token)
    if($Provider -eq 'GitHub'){$request.Headers.UserAgent.ParseAdd('toguz-disposable-spike/0.0.0');$request.Headers.Add('X-GitHub-Api-Version','2022-11-28')}
    if ($null -ne $Body) { $request.Content = [Net.Http.StringContent]::new(($Body | ConvertTo-Json -Depth 60 -Compress), [Text.Encoding]::UTF8, 'application/json') }
    try {
        $response = $client.SendAsync($request).GetAwaiter().GetResult()
        $entry.httpStatus = [int]$response.StatusCode
        $raw = $response.Content.ReadAsStringAsync().GetAwaiter().GetResult()
        $json = @{}
        if ($raw) {
            try { $json = ConvertFrom-Json -InputObject $raw -AsHashtable -NoEnumerate -Depth 100 }
            catch { if ($entry.httpStatus -lt 300) { throw 'Successful response was not JSON' } }
        }
        if ($entry.httpStatus -ge 300) {
            $entry.result = 'PROVIDER_REJECTED'
            $entry.errorCode = if ($json.error -is [Collections.IDictionary]) { $json.error.code } elseif ($json.code) { $json.code } else { 'HTTP_' + $entry.httpStatus }
            if ($RecordNonsecretError) {
                $entry.errorMessage = if ($json.error -is [Collections.IDictionary]) { $json.error.message } elseif ($json.message) { $json.message } elseif ($json.error -is [string]) { $json.error } else { 'No diagnostic message' }
            }
            throw 'Provider rejected the request; see sanitized ledger'
        }
        $entry.observation = & $Summarize $json
        $entry.result = 'OBSERVED'
        return $json
    } catch {
        if (-not $entry.result) { $entry.result = 'REQUEST_OR_COLLECTION_FAILED'; $entry.exceptionType = $_.Exception.GetType().Name }
        throw 'Spike API operation failed; credentials and raw response suppressed'
    } finally {
        $entry.completedUTC = [DateTimeOffset]::UtcNow.ToString('o'); Save-SpikeLedger
        if ($response) { $response.Dispose() }; $request.Dispose(); $client.Dispose(); $handler.Dispose()
    }
}
