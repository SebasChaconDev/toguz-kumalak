export default function ProbePage() {
  return <main>
    <h1>EXPERIMENTAL · infrastructure probe</h1>
    <p>Sonda local desechable de Next.js. No implementa Toguz Kumalak ni sesiones de producto.</p>
    <p>Rewrites locales: <code>/probe/identity/*</code> y <code>/probe/game/*</code>.</p>
    <p>Los resultados locales no verifican Vercel, OIDC, cookies del navegador ni protección de previews.</p>
  </main>;
}
