// Controlled local doubles exercise Next's HTTP boundary. They are NOT Go services,
// an Identity implementation, an OIDC provider, or a browser cookie jar.
import assert from 'node:assert/strict';
import { createServer } from 'node:http';
import { randomBytes } from 'node:crypto';
import { spawn } from 'node:child_process';
import { performance } from 'node:perf_hooks';
import { readFile, readdir, writeFile } from 'node:fs/promises';
import { join } from 'node:path';

const credential = randomBytes(32).toString('hex');
const marker = randomBytes(16).toString('hex');
let receivedCookie = false;
let requests = 0;
let readinessAttempts = 0;
let child;
const servers = [];
const samples = [];
const next = 'http://127.0.0.1:18100';

async function request(path, options = {}) {
  const start = performance.now();
  const response = await fetch(`${next}${path}`, { ...options, signal: AbortSignal.timeout(5000) });
  requests++;
  samples.push({ path: path.split('?')[0], status: response.status, elapsedMs: +(performance.now() - start).toFixed(2) });
  return response;
}

async function fixture(service, port) {
  const server = createServer((req, res) => {
    res.setHeader('Content-Type', 'application/json');
    if (req.url?.startsWith('/cookie')) {
      if (req.headers['x-local-fixture-credential'] !== credential) {
        res.writeHead(401).end('{"error":"unauthorized_fixture_request"}');
        return;
      }
      // An opaque test marker, with NO session/authentication meaning.
      res.setHeader('Set-Cookie', `spike_probe=${marker}; HttpOnly; SameSite=Lax; Secure; Path=/; Max-Age=300`);
      res.end('{"synthetic":true,"authority":false}');
      return;
    }
    if (req.headers.authorization !== `Bearer ${credential}`) {
      res.writeHead(401).end('{"error":"unauthorized_fixture_request"}');
      return;
    }
    receivedCookie ||= req.headers.cookie === `spike_probe=${marker}`;
    res.end(JSON.stringify({ service, path: req.url, cookiePresent: Boolean(req.headers.cookie) }));
  });
  await new Promise((resolve, reject) => {
    server.once('error', reject);
    server.listen(port, '127.0.0.1', resolve);
  });
  servers.push(server);
}

try {
  await fixture('identity-double', 18101);
  await fixture('game-double', 18102);
  child = spawn(process.execPath, ['node_modules/next/dist/bin/next', 'start', '--hostname', '127.0.0.1', '--port', '18100'], {
    env: { ...process.env, NEXT_TELEMETRY_DISABLED: '1' },
    stdio: ['ignore', 'pipe', 'pipe'],
    windowsHide: true,
  });
  let startup = '';
  child.stdout.on('data', chunk => { startup += chunk; });
  child.stderr.on('data', chunk => { startup += chunk; });
  for (let i = 0; i < 100; i++) {
    if (child.exitCode !== null) throw new Error('Next process exited before readiness');
    try {
      readinessAttempts++;
      const ready = await fetch(next, { signal: AbortSignal.timeout(500) });
      if (ready.ok) break;
    } catch {}
    if (i === 99) throw new Error('Next readiness deadline exceeded');
    await new Promise(resolve => setTimeout(resolve, 100));
  }
  const page = await request('/');
  assert.equal(page.status, 200);
  const html = await page.text();
  assert.match(html, /EXPERIMENTAL/);
  assert.ok(!html.includes(credential));
  assert.ok(!html.includes(marker));
  const absent = await request('/probe/identity/health');
  assert.equal(absent.status, 401);
  const invalid = await request('/probe/game/health', { headers: { Authorization: 'Bearer invalid-synthetic-fixture' } });
  assert.equal(invalid.status, 401);
  for (const service of ['identity', 'game']) {
    const response = await request(`/probe/${service}/health?sample=synthetic`, { headers: { Authorization: `Bearer ${credential}` } });
    assert.equal(response.status, 200);
    const body = await response.json();
    assert.equal(body.service, `${service}-double`);
    assert.equal(body.path, '/health?sample=synthetic');
  }
  const deniedCookie = await request('/probe/identity/cookie', { method: 'POST' });
  assert.equal(deniedCookie.status, 401);
  assert.equal(deniedCookie.headers.get('set-cookie'), null);
  const cookieResponse = await request('/probe/identity/cookie', {
    method: 'POST', headers: { 'X-Local-Fixture-Credential': credential },
  });
  assert.equal(cookieResponse.status, 200);
  const setCookie = cookieResponse.headers.get('set-cookie');
  assert.ok(setCookie?.includes('HttpOnly'));
  assert.ok(setCookie?.includes('SameSite=Lax'));
  assert.ok(setCookie?.includes('Secure'));
  assert.ok(setCookie?.includes('Path=/'));
  // Manual header intentionally proves forwarding only, NOT Secure/SameSite browser enforcement.
  const withCookie = await request('/probe/game/health', {
    headers: { Authorization: `Bearer ${credential}`, Cookie: `spike_probe=${marker}` },
  });
  assert.equal(withCookie.status, 200);
  assert.equal((await withCookie.json()).cookiePresent, true);
  assert.equal(receivedCookie, true);
  assert.ok(!startup.includes(credential));
  assert.ok(!startup.includes(marker));
  const clientFiles = (await readdir('.next/static', { recursive: true, withFileTypes: true }))
    .filter(entry => entry.isFile() && entry.name.endsWith('.js'));
  assert.ok(clientFiles.length > 0);
  for (const entry of clientFiles) {
    const artifact = await readFile(join(entry.parentPath, entry.name), 'utf8');
    assert.ok(!artifact.includes(credential));
    assert.ok(!artifact.includes(marker));
  }
  const evidence = {
    timestampUtc: new Date().toISOString(), node: process.version,
    environment: 'Windows local loopback, production Next server, controlled Node HTTP doubles',
    result: 'PASS', measuredRequests: requests, readinessAttempts, totalHttpAttempts: requests + readinessAttempts,
    clientJavaScriptFilesChecked: clientFiles.length, samples,
    observations: [
      'Root HTML 200 with EXPERIMENTAL marker and without generated fixture credentials',
      'Two fixed upstream rewrites preserve paths, query and Authorization; upstream 401 is preserved',
      'Unauthenticated cookie fixture request returns 401 and no Set-Cookie',
      'Authorized controlled fixture Set-Cookie attributes survive rewrite',
      'Explicit manually supplied cookie header reaches second same-origin rewrite',
      'Generated fixture credentials/marker absent from captured Next logs',
      'Generated fixture credentials/marker absent from emitted client JavaScript files',
    ],
    notVerified: ['Browser enforcement of HttpOnly/Secure/SameSite', 'Real Identity session and OIDC', 'Actual Go-service cookie flow', 'TLS and Vercel preview protection', 'Remote origins, environment isolation and remote build'],
  };
  await writeFile('evidence/web-smoke.json', `${JSON.stringify(evidence, null, 2)}\n`);
  console.log(JSON.stringify(evidence, null, 2));
} finally {
  if (child && child.exitCode === null) {
    child.kill();
    await new Promise(resolve => { child.once('exit', resolve); setTimeout(resolve, 3000); });
  }
  await Promise.all(servers.map(server => new Promise(resolve => { server.closeAllConnections(); server.close(resolve); })));
}
