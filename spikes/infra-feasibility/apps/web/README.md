# EXPERIMENTAL web infrastructure probe

Throwaway local spike, not game/product code. No authentication/session implementation, external resources or deployment. The page only identifies the experiment. React server rendering and Next's own client runtime are the only web behavior.

## Independent root and versions

All required sources and the npm lock live here: `package.json`, `package-lock.json`, `next.config.mjs`, `tsconfig.json`, `next-env.d.ts`, `app/` and `lib/`. There is no root workspace, parent manifest or dependency on Go sources. Build outputs (`.next/`) and installed packages (`node_modules/`) are ignored. Next's Turbopack root is this directory. The production build creates the generated type declarations imported by `next-env.d.ts`.

Observed toolchain: Windows, Node **24.13.0**, npm **11.6.2**, Next **16.4.2**, React/React DOM **19.3.1**, TypeScript **5.9.3**, `@types/node` **24.13.0**, React types **19.3.0**. Direct dependencies are exact, and transitive resolution/integrity is locked. npm registry metadata was consulted on 2026-10-22 UTC (2026-10-21 Caracas). TypeScript 5.9.3 was deliberately selected as a stable compatible compiler; it is not claimed as the newest npm release.

Run from this directory in PowerShell:

```powershell
$env:NEXT_TELEMETRY_DISABLED='1'
npm ci --ignore-scripts --no-fund --no-audit
npm test
npm run build
npm run smoke
```

`npm run build` runs the production compiler from this root; the first observed build had no `.next` cache, exited 0, and took **11.0896239 seconds** overall. Next reported 1477 ms for compilation and 1448 ms for TypeScript. This is a single local sample, not a remote build/cold-start benchmark. npm's download cache is distinct from the absent compilation cache. Evidence: [build log](evidence/build.log), [summary](evidence/build-summary.json).

For a manual local server, run `npm start` (127.0.0.1:18100). Rewrites point to `http://127.0.0.1:18101` for `/probe/identity/:path*` and `http://127.0.0.1:18102` for `/probe/game/:path*`. Server-only build-time overrides `SPIKE_IDENTITY_ORIGIN` and `SPIKE_GAME_ORIGIN` accept HTTP origins with hostname 127.0.0.1 only and reject credentials, paths, queries and fragments. No `NEXT_PUBLIC_` credentials, request-selected destinations, `.env` files, tokens or secrets are shipped. Origins are fixed in Next's build output; a changed destination requires rebuilding. External HTTPS origins are deliberately unsupported until the remote gate and a reviewed allowlist/configuration change.

## Actual local checks and limits

`npm test` exercises the real origin validator, including remote host, wildcard interface, embedded credentials and URL component rejection. The test was observed red against a placeholder before implementing validation, then green. [Latest unit result](evidence/unit-tests.json): one test passed, zero failed. A subsequent locked `npm ci --ignore-scripts --no-fund --no-audit` also succeeded, installing 29 packages in 22 seconds; the smoke and unit suite passed after that reinstall.

`npm run smoke` starts the production Next server and **controlled Node HTTP doubles**, all on loopback, then stops them. Ports 18100/18101/18102 must be free. The doubles exist only inside the smoke process; they are not Identity or Game. One random in-memory fixture credential protects their sample endpoints; it never goes into files or logs. An opaque synthetic cookie marker has no authority/session meaning.

The smoke checks eight measured requests: experimental HTML 200; absent/invalid fixture authorization preserved as 401; successful Identity/Game path/query forwarding with Authorization; rejected cookie request without Set-Cookie; authorized Set-Cookie retaining HttpOnly, SameSite=Lax, Secure and Path; and a manually supplied cookie reaching the other rewrite. Readiness attempts are counted separately. [Sanitized smoke evidence](evidence/web-smoke.json) contains per-request times, status, counts and explicit omissions, with no cookie/credential values. Captured Next logs, returned HTML and nine emitted client JavaScript files are checked for the generated values.

An additional `npm audit --omit=dev --json` attempt exited 1 because the public npm advisories endpoint returned a socket hang-up. The dependency vulnerability assessment is BLOCKED, not a clean audit. Build/install success does not establish an absence of dependency vulnerabilities.

The cookie test verifies HTTP header transport only. It manually supplies a Cookie header over loopback HTTP; it does **not** demonstrate browser Secure/SameSite/HttpOnly enforcement, TLS, a real Identity session, OIDC or Google integration. There is no cookie/token mint route in the Next app. The cookie fixture rejects missing credentials and disappears on shutdown. Actual Go-service cookie exchange is NOT_RUN.

POC-01 **LOCAL web component PASS** (independent build); the four-project/shared-Go result must be assembled with the Go evidence. POC-06 **LOCAL transport component PASS**; complete POC-06 remains BLOCKED/NOT_RUN for browser/session and remote preview requirements. Remote POC-01/06 and Vercel environment/protection are **NOT_RUN**, blocked by the pending `APRUEBO SPIKE REMOTO` gate. No remote viability or security end-to-end PASS is asserted. Timings involve doubles on one host and cannot represent Go service latency, deployment geography or capacity.

## Official references consulted

Consulted 2026-10-22 UTC, 2026-10-21 Caracas:

- [Next.js installation](https://nextjs.org/docs/app/getting-started/installation) — framework/toolchain setup.
- [Next.js rewrites](https://nextjs.org/docs/app/api-reference/config/next-config-js/rewrites) — external destination proxy semantics.
- [Next.js cookies](https://nextjs.org/docs/app/api-reference/functions/cookies) — server cookie operations and attributes; no claim of browser verification from these docs.
- npm metadata obtained with `npm view next version engines peerDependencies --json`, `npm view react version --json`, `npm view react-dom version --json`, `npm view typescript@5.9.3 version --json`, and matching type package metadata. These inspect the public package registry, not provider resources.
