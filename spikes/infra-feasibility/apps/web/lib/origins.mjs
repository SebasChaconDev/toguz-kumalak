export function probeOrigin(value, defaultPort) {
  const candidate = value === undefined ? `http://127.0.0.1:${defaultPort}` : value;
  const invalid = () => new Error('Probe target must be a loopback HTTP origin without credentials, path, query or fragment');
  let url;
  try {
    url = new URL(candidate);
  } catch {
    throw invalid();
  }
  if (url.protocol !== 'http:' || url.hostname !== '127.0.0.1' ||
      url.username || url.password || url.pathname !== '/' ||
      url.search || url.hash || url.port === '0') {
    throw invalid();
  }
  return url.origin;
}
