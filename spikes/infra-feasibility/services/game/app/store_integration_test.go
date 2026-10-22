//go:build integration

package app

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe/testkit"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Real PostgreSQL only; missing/unreachable DB is a failure, never a skipped PASS.
func fixture(t *testing.T) (*Store, probe.Auth) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	p, e := probe.Pool(ctx, os.Getenv("GAME_DATABASE_URL"))
	if e != nil {
		t.Fatal("BLOCKED: Game database unavailable")
	}
	t.Cleanup(p.Close)
	if p.Ping(ctx) != nil {
		t.Fatal("BLOCKED: Game database unreachable")
	}
	down, _ := Migrations.ReadFile("migrations/down.sql")
	up, _ := Migrations.ReadFile("migrations/up.sql")
	// Fresh dedicated synthetic DB only. Every test rebuilds its own fixture.
	_, _ = p.Exec(ctx, string(down))
	if _, e = p.Exec(ctx, string(up)); e != nil {
		t.Fatal("fixture migration failed")
	}
	return &Store{Pool: p}, probe.Auth{UserKey: bytes.Repeat([]byte{1}, 32), ServiceKey: bytes.Repeat([]byte{2}, 32)}
}
func TestSQLIsolationMigrationAndPool(t *testing.T) {
	s, _ := fixture(t)
	ctx := context.Background()
	if s.Pool.Config().MaxConns != 2 {
		t.Fatal("pool unbounded")
	}
	idDSN := os.Getenv("IDENTITY_DATABASE_URL")
	ip, e := probe.Pool(ctx, idDSN)
	if e != nil {
		t.Fatal("identity config missing")
	}
	defer ip.Close()
	if ip.Ping(ctx) != nil {
		t.Fatal("identity unavailable")
	}
	// Both connection directions must fail with SQLSTATE insufficient_privilege.
	for _, pair := range [][2]string{{idDSN, "game"}, {os.Getenv("GAME_DATABASE_URL"), "identity"}} {
		cfg, e := pgx.ParseConfig(pair[0])
		if e != nil {
			t.Fatal("config")
		}
		cfg.Database = pair[1]
		c, e := pgx.ConnectConfig(ctx, cfg)
		if e == nil {
			c.Close(ctx)
			t.Fatal("cross-owner database connect allowed")
		}
		var pe *pgconn.PgError
		if !errors.As(e, &pe) || pe.Code != "42501" {
			t.Fatal("cross-owner denial was not an SQL privilege denial")
		}
	}
	// Identity migration is reversible independently of Game.
	_, _ = ip.Exec(ctx, "DROP TABLE IF EXISTS synthetic_identity")
	if _, e = ip.Exec(ctx, "CREATE TABLE synthetic_identity(id text PRIMARY KEY); INSERT INTO synthetic_identity VALUES('alice')"); e != nil {
		t.Fatal("identity up migration")
	}
	if _, e = ip.Exec(ctx, "DROP TABLE synthetic_identity"); e != nil {
		t.Fatal("identity down migration")
	}
	down, _ := Migrations.ReadFile("migrations/down.sql")
	if _, e = s.Pool.Exec(ctx, string(down)); e != nil {
		t.Fatal("game down migration")
	}
	up, _ := Migrations.ReadFile("migrations/up.sql")
	if _, e = s.Pool.Exec(ctx, string(up)); e != nil {
		t.Fatal("game re-up migration")
	}
}
func TestIdempotencyRollbackAndRestart(t *testing.T) {
	s, _ := fixture(t)
	ctx := context.Background()
	in := Intent{"one", 0, "synthetic"}
	first, e := s.Apply(ctx, "owned-game", "alice", in)
	if e != nil || first != (State{1, 1}) {
		t.Fatal("initial commit")
	}
	// New Store instance simulates process-local memory loss; receipts stay in PostgreSQL.
	restarted := &Store{Pool: s.Pool}
	again, e := restarted.Apply(ctx, "owned-game", "alice", in)
	if e != nil || again != first {
		t.Fatal("retry duplicated or lost effect")
	}
	in.Payload = "different"
	if _, e = s.Apply(ctx, "owned-game", "alice", in); !errors.Is(e, ErrConflict) {
		t.Fatal("ID reuse not rejected")
	}
	s.BeforeCommit = func() error { return errors.New("injected") }
	if _, e = s.Apply(ctx, "owned-game", "alice", Intent{"rollback", 1, "x"}); e == nil {
		t.Fatal("precommit failure accepted")
	}
	s.BeforeCommit = nil
	st, e := s.Get(ctx, "owned-game", "alice")
	if e != nil || st != (State{1, 1}) {
		t.Fatal("rollback leaked partial state")
	}
	st, e = s.Apply(ctx, "owned-game", "alice", Intent{"rollback", 1, "x"})
	if e != nil || st != (State{2, 2}) {
		t.Fatal("rollback left a receipt")
	}
}
func TestConcurrentRevision(t *testing.T) {
	s, _ := fixture(t)
	start := make(chan struct{})
	var wg sync.WaitGroup
	result := make(chan error, 2)
	for _, id := range []string{"race-a", "race-b"} {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			<-start
			_, e := s.Apply(context.Background(), "owned-game", "alice", Intent{id, 0, "x"})
			result <- e
		}(id)
	}
	close(start)
	wg.Wait()
	close(result)
	ok, conflict := 0, 0
	for e := range result {
		if e == nil {
			ok++
		} else if errors.Is(e, ErrStale) {
			conflict++
		} else {
			t.Fatal("unexpected race error")
		}
	}
	st, e := s.Get(context.Background(), "owned-game", "alice")
	if e != nil || ok != 1 || conflict != 1 || st != (State{1, 1}) {
		t.Fatal("concurrency duplicated/lost effects")
	}
}
func TestNoticeFailureDuplicateOutOfOrderAndREST(t *testing.T) {
	s, a := fixture(t)
	ctx := context.Background()
	var notices []Notice
	s.Notify = func(ctx context.Context, n Notice) error {
		confirmed, e := s.Get(ctx, n.GameID, "alice")
		if e != nil || confirmed.Revision != n.Revision {
			t.Fatal("notification before commit")
		}
		notices = append(notices, n)
		return errors.New("local Ably double unavailable")
	}
	for i, id := range []string{"n1", "n2"} {
		if _, e := s.Apply(ctx, "owned-game", "alice", Intent{id, int64(i), "x"}); e != nil {
			t.Fatal("notice error invalidated commit")
		}
	}
	h := New(a, s)
	for _, n := range []Notice{notices[1], notices[1], notices[0]} {
		_ = n
		w := httptest.NewRecorder()
		h.ServeHTTP(w, testkit.HTTPRequest("GET", "/games/owned-game", "", testkit.Token(a, "user", "game", "alice")))
		if w.Code != 200 || strings.TrimSpace(w.Body.String()) != `{"revision":2,"effects":2}` {
			t.Fatal("REST reconciliation lost confirmed state")
		}
	}
	// Completely missing notice also recovers via that same authoritative GET.
	w := httptest.NewRecorder()
	h.ServeHTTP(w, testkit.HTTPRequest("GET", "/games/owned-game", "", testkit.Token(a, "user", "game", "alice")))
	if w.Code != 200 {
		t.Fatal("lost notice blocks REST")
	}
}
func TestLostHTTPResponseAndResourceAuthorization(t *testing.T) {
	s, a := fixture(t)
	h := New(a, s)
	firstStatus := make(chan int, 1)
	dropped := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buffer := httptest.NewRecorder()
		h.ServeHTTP(buffer, r)
		firstStatus <- buffer.Code
		conn, _, e := w.(http.Hijacker).Hijack()
		if e == nil {
			conn.Close()
		}
	}))
	defer dropped.Close()
	body := `{"operationId":"lost","expectedRevision":0,"payload":"x"}`
	token := testkit.Token(a, "user", "game", "alice")
	req, _ := http.NewRequest("POST", dropped.URL+"/games/owned-game/operations", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 2 * time.Second}
	resp, e := client.Do(req)
	if resp != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	if e == nil {
		t.Fatal("response was not lost")
	}
	if <-firstStatus != 200 {
		t.Fatal("original request did not confirm")
	}
	confirmed, e := s.Get(context.Background(), "owned-game", "alice")
	if e != nil || confirmed != (State{1, 1}) {
		t.Fatal("first commit was not durable before retry")
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, testkit.HTTPRequest("POST", "/games/owned-game/operations", body, token))
	if w.Code != 200 || strings.TrimSpace(w.Body.String()) != `{"revision":1,"effects":1}` {
		t.Fatal("retry after lost response incorrect")
	}
	for _, path := range []string{"/games/foreign-game", "/games/foreign-game/operations"} {
		method := "GET"
		if strings.HasSuffix(path, "operations") {
			method = "POST"
		}
		w = httptest.NewRecorder()
		h.ServeHTTP(w, testkit.HTTPRequest(method, path, body, token))
		if w.Code != 403 {
			t.Fatal("foreign resource allowed")
		}
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, testkit.HTTPRequest("POST", "/games/owned-game/operations", body, testkit.Token(a, "service", "game", "ai")))
	if w.Code != 401 {
		t.Fatal("AI can commit Game state")
	}
	st, e := s.Get(context.Background(), "owned-game", "alice")
	if e != nil || st != (State{1, 1}) {
		t.Fatal("unauthorized operation changed state")
	}
}
