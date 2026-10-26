import test from 'node:test';
import assert from 'node:assert/strict';
import { probeOrigin } from '../lib/origins.mjs';

// Catches a rewrite target becoming a configurable public proxy or carrying credentials.
test('only explicit loopback HTTP origins can be rewrite targets', () => {
  assert.equal(probeOrigin(undefined, 18101), 'http://127.0.0.1:18101');
  assert.equal(probeOrigin('http://127.0.0.1:18201', 18101), 'http://127.0.0.1:18201');
  for (const target of [
    'https://example.com', 'http://example.com', 'http://0.0.0.0:18101',
    'http://127.0.0.1:18101/private', 'http://127.0.0.1:18101?secret=value',
    'http://127.0.0.1:18101#fragment', 'http://user:password@127.0.0.1:18101',
    'file:///etc/passwd', 'http://127.0.0.1:0', '', 'not a URL',
  ]) {
    assert.throws(() => probeOrigin(target, 18101), /loopback HTTP origin/);
  }
});

test('remote spike allows only its two fixed protected service origins', () => {
  assert.equal(probeOrigin('https://tk-spike01-game-ultra-seneca.vercel.app', 18102, true), 'https://tk-spike01-game-ultra-seneca.vercel.app');
  for (const target of ['https://example.com', 'https://tk-spike01-game-ultra-seneca.vercel.app.evil.com',
    'https://tk-spike01-game-ultra-seneca.vercel.app/path', 'https://tk-spike01-game-ultra-seneca.vercel.app?token=bad',
    'https://user:password@tk-spike01-game-ultra-seneca.vercel.app', 'https://tk-spike01-ai-ultra-seneca.vercel.app']) {
    assert.throws(() => probeOrigin(target, 18102, true));
  }
});
