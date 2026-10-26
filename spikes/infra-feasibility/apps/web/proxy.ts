import { NextRequest, NextResponse } from 'next/server';

// Disposable server-side forwarding. No credentials enter HTML or browser JS.
export function proxy(request: NextRequest) {
  const headers = new Headers(request.headers);
  headers.delete('x-vercel-trusted-oidc-idp-token');
  if (process.env.SPIKE_REMOTE_ENABLED === '1') {
    const oidc = request.headers.get('x-vercel-oidc-token') ?? process.env.VERCEL_OIDC_TOKEN;
    if (!oidc) return new NextResponse('Experimental protection unavailable', { status: 503 });
    headers.set('x-vercel-trusted-oidc-idp-token', oidc);
  }
  return NextResponse.next({ request: { headers } });
}

export const config = { matcher: ['/probe/identity/:path*', '/probe/game/:path*'] };
