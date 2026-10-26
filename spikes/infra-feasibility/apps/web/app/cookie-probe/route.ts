export const dynamic = 'force-dynamic';

export function GET() {
  if (process.env.SPIKE_REMOTE_ENABLED !== '1') return new Response(null, { status: 404 });
  const html = `<!doctype html><html lang="es"><meta charset="utf-8"><title>Cookie probe</title>
<body><h1>Prueba sintética de cookie</h1><pre id="result">RUNNING</pre><script>
(async () => {try {
  const response = await fetch('/probe/identity/cookie-probe', {credentials:'same-origin',cache:'no-store'});
  const data = await response.json();
  const hidden = !document.cookie.includes('__Host-spike-cookie');
  document.getElementById('result').textContent = JSON.stringify({status:response.status,cookieReceived:data.cookieReceived,httpOnlyHidden:hidden,pass:response.ok && data.cookieReceived && hidden});
} catch (_) {document.getElementById('result').textContent='FAIL';}})();
</script></body></html>`;
  return new Response(html, { headers: {
    'Content-Type': 'text/html; charset=utf-8',
    'Cache-Control': 'no-store',
    'Set-Cookie': '__Host-spike-cookie=synthetic-cookie-only; Path=/; Secure; HttpOnly; SameSite=Lax',
    'Content-Security-Policy': "default-src 'none'; script-src 'unsafe-inline'; connect-src 'self'; base-uri 'none'; frame-ancestors 'none'",
  }});
}
