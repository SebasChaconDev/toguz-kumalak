export function probeOrigin(value, defaultPort, remote = process.env.SPIKE_REMOTE_ENABLED === '1') {
  const candidate = value === undefined ? `http://127.0.0.1:${defaultPort}` : value;
  const invalid = () => new Error('Probe target must be a loopback HTTP origin without credentials, path, query or fragment');
  let url;
  try {
    url = new URL(candidate);
  } catch {
    throw invalid();
  }
  if (remote && url.protocol === 'https:' &&
      /^tk-spike01-(identity|game)-ultra-seneca\.vercel\.app$/.test(url.hostname) &&
      !url.username && !url.password && !url.port && url.pathname === '/' && !url.search && !url.hash) {
    return url.origin;
  }
  if (url.protocol !== 'http:' || url.hostname !== '127.0.0.1' ||
      url.username || url.password || url.pathname !== '/' ||
      url.search || url.hash || url.port === '0') {
    throw invalid();
  }
  return url.origin;
}
