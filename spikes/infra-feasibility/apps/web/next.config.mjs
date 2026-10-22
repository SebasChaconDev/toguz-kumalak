import { fileURLToPath } from 'node:url';
import { probeOrigin } from './lib/origins.mjs';

// Build-time, server-only inputs. Never accept rewrite destinations from a request.
const identityOrigin = probeOrigin(process.env.SPIKE_IDENTITY_ORIGIN, 18101);
const gameOrigin = probeOrigin(process.env.SPIKE_GAME_ORIGIN, 18102);

export default {
  poweredByHeader: false,
  turbopack: { root: fileURLToPath(new URL('.', import.meta.url)) },
  async rewrites() {
    return [
      { source: '/probe/identity/:path*', destination: `${identityOrigin}/:path*` },
      { source: '/probe/game/:path*', destination: `${gameOrigin}/:path*` },
    ];
  },
};
