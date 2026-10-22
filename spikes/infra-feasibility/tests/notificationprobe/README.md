# Local notification verification harness

Test-only Go module; no new product service or broker dependency. `integration`
tests use the existing Game `app.New`, transactional `Store`, real PostgreSQL,
and a controlled HTTP notification double on ephemeral loopback ports.

Run after the existing Game fixtures have finished resetting their own tables:

```powershell
./Run.ps1
```

The runner loads the ignored `.local/test.env` in process memory, restores the
previous environment, and writes redacted JSONL plus per-case JSON under
`evidence/integral-notifications*`. It inserts only fresh `notification-run-*`
synthetic games; it does not reset tables, restart containers, or delete records.
Do not run while another harness drops tables or restarts PostgreSQL.

`NOT-01` through `NOT-07` assert delivery, loss, duplicates, reversed delivery,
lost HTTP response, stopped/restarted broker listener, and bounded unresponsive
publication. Each records PostgreSQL revision, effects and receipt counts before
and after. Publication reads the committed state over a separate DB query. The
test client always reconciles with authenticated Game REST rather than accepting
a notice's claimed revision, and retries the same intent after response loss.

`CAP-01` through `CAP-08` exercise exact channel, action, expiration and signature
restrictions using distinct synthetic client and Game publisher signing keys.
The deliberately public key literals are disposable test fixtures, not credentials.
These tests validate a simulated capability contract only: **Ably real is NOT_RUN**.
No Ably SDK, token endpoint, account, network service or production security is
validated here. Broker listener restart is simulated locally, not a real provider
restart. The client is a Go test client, not the experimental Next.js client.

`./Run.ps1 -Race -EvidenceName integral-notifications-race` requires a Go host
with CGO and a compiler. `go vet -tags=integration ./...` checks the entire module.
The evidence files named `*-red*` retain the intentional TDD failures for an
unprotected capability endpoint and an adapter that loses every delivery.
