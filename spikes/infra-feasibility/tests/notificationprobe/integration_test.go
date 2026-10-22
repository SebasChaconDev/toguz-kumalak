//go:build integration

package notificationprobe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe/testkit"
	game "github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/services/game/app"
	"github.com/jackc/pgx/v5/pgxpool"
)

type fixture struct {
	t              *testing.T
	pool           *pgxpool.Pool
	gameID         string
	store          *game.Store
	auth           probe.Auth
	gameServer     *httptest.Server
	broker         *broker
	brokerServer   *httptest.Server
	endpoint       atomic.Value
	before         snapshot
	notifyFailures atomic.Int64
}

var serial atomic.Uint64

func newFixture(t *testing.T, mode string) *fixture {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	p, err := probe.Pool(ctx, os.Getenv("GAME_DATABASE_URL"))
	if err != nil {
		t.Fatal("BLOCKED: Game PostgreSQL config unavailable")
	}
	t.Cleanup(p.Close)
	if p.Ping(ctx) != nil {
		t.Fatal("BLOCKED: real Game PostgreSQL unreachable")
	}
	id := fmt.Sprintf("notification-run-%d-%d", time.Now().UnixNano(), serial.Add(1))
	if _, err = p.Exec(ctx, "INSERT INTO synthetic_games(id,owner,revision,effects) VALUES($1,$2,0,0)", id, "alice"); err != nil {
		t.Fatal("synthetic insert failed; this harness never creates or resets tables")
	}
	f := &fixture{t: t, pool: p, gameID: id, auth: probe.Auth{UserKey: bytes.Repeat([]byte{1}, 32), ServiceKey: bytes.Repeat([]byte{2}, 32)}}
	f.before = f.db()
	f.broker = newBroker()
	f.broker.mode = mode
	f.broker.readCommitted = func(ctx context.Context, n notification) (snapshot, error) { return readSnapshot(ctx, p, n.GameID) }
	f.brokerServer = httptest.NewServer(f.broker)
	f.endpoint.Store(f.brokerServer.URL)
	t.Cleanup(func() { f.brokerServer.Close() })
	f.store = &game.Store{Pool: p, Notify: func(ctx context.Context, n game.Notice) error {
		raw, _ := json.Marshal(notification{n.GameID, n.Revision})
		req, _ := http.NewRequestWithContext(ctx, "POST", f.endpoint.Load().(string)+"/channels/"+n.GameID+"/messages", bytes.NewReader(raw))
		req.Header.Set("Authorization", "Bearer "+f.broker.token("game", n.GameID, "publish", time.Now().Add(time.Minute)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			f.notifyFailures.Add(1)
			return errors.New("local controlled broker request failed")
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != 202 {
			f.notifyFailures.Add(1)
			return errors.New("local controlled broker denied publication")
		}
		return nil
	}}
	f.gameServer = httptest.NewServer(game.New(f.auth, f.store))
	t.Cleanup(f.gameServer.Close)
	return f
}
func readSnapshot(ctx context.Context, p *pgxpool.Pool, id string) (snapshot, error) {
	var s snapshot
	err := p.QueryRow(ctx, "SELECT revision,effects,(SELECT count(*) FROM synthetic_receipts WHERE game_id=$1) FROM synthetic_games WHERE id=$1", id).Scan(&s.Revision, &s.Effects, &s.Receipts)
	return s, err
}
func (f *fixture) db() snapshot {
	f.t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s, err := readSnapshot(ctx, f.pool, f.gameID)
	if err != nil {
		f.t.Fatal("Game PostgreSQL snapshot failed")
	}
	return s
}
func (f *fixture) request(method, path string, body any) (int, game.State, error) {
	var payload []byte
	if body != nil {
		payload, _ = json.Marshal(body)
	}
	r, _ := http.NewRequest(method, f.gameServer.URL+path, bytes.NewReader(payload))
	r.Header.Set("Authorization", "Bearer "+testkit.Token(f.auth, "user", "game", "alice"))
	r.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{Timeout: 2 * time.Second}).Do(r)
	if err != nil {
		return 0, game.State{}, err
	}
	defer resp.Body.Close()
	var state game.State
	if resp.StatusCode == 200 {
		if json.NewDecoder(resp.Body).Decode(&state) != nil {
			return resp.StatusCode, state, errors.New("invalid Game state response")
		}
	} else {
		io.Copy(io.Discard, resp.Body)
	}
	return resp.StatusCode, state, nil
}
func (f *fixture) apply(id string, revision int64) game.State {
	f.t.Helper()
	code, s, err := f.request("POST", "/games/"+f.gameID+"/operations", game.Intent{OperationID: id, ExpectedRevision: revision, Payload: "synthetic"})
	if err != nil || code != 200 {
		f.t.Fatalf("Game operation failed: status=%d", code)
	}
	if s != (game.State{Revision: revision + 1, Effects: revision + 1}) {
		f.t.Fatalf("unexpected operation state: %+v", s)
	}
	return s
}
func (f *fixture) poll() []notification {
	f.t.Helper()
	r, _ := http.NewRequest("GET", f.endpoint.Load().(string)+"/channels/"+f.gameID+"/messages", nil)
	r.Header.Set("Authorization", "Bearer "+f.broker.token("client", f.gameID, "subscribe", time.Now().Add(time.Minute)))
	resp, err := (&http.Client{Timeout: time.Second}).Do(r)
	if err != nil {
		f.t.Fatal("local subscriber poll failed")
	}
	defer resp.Body.Close()
	var ns []notification
	if resp.StatusCode != 200 || json.NewDecoder(resp.Body).Decode(&ns) != nil {
		f.t.Fatal("local subscriber response failed")
	}
	return ns
}

// A notice triggers a read, never an operation. Ignore its claimed revision
// (duplicates, stale values, invented future values) in favour of real REST.
func (f *fixture) reconcile(_ *notification, want int64) game.State {
	f.t.Helper()
	code, s, err := f.request("GET", "/games/"+f.gameID, nil)
	if err != nil || code != 200 || s != (game.State{Revision: want, Effects: want}) {
		f.t.Fatalf("authoritative REST recovery failed: status=%d state=%+v", code, s)
	}
	return s
}
func (f *fixture) report(id string, want int64, details map[string]any) {
	f.t.Helper()
	after := f.db()
	if f.before != (snapshot{}) || after != (snapshot{want, want, want}) {
		f.t.Fatalf("unexpected persisted effects: pre=%+v post=%+v", f.before, after)
	}
	record := map[string]any{"caseId": id, "status": "PASS", "before": f.before, "after": after, "publicationAudits": f.broker.audit(), "notifyFailures": f.notifyFailures.Load(), "ablyRealValidated": false, "broker": "local controlled HTTP double", "game": "real app.New + Store", "database": "real PostgreSQL"}
	for k, v := range details {
		record[k] = v
	}
	raw, _ := json.Marshal(record)
	f.t.Log("CASE " + string(raw))
}

// Moving Notify before commit, trusting a hint revision, or reapplying an intent
// while reconciling fails through observed DB/REST state, never mock presence.
func TestNotificationDeliveryCommitAndREST(t *testing.T) {
	f := newFixture(t, "")
	f.apply("delivery", 0)
	notices := f.poll()
	if len(notices) != 1 || notices[0] != (notification{f.gameID, 1}) {
		t.Fatalf("delivered notices=%d want=1", len(notices))
	}
	audits := f.broker.audit()
	if len(audits) != 1 || audits[0].Committed != (snapshot{1, 1, 1}) {
		t.Fatal("commit not observed by broker publication")
	}
	rest := f.reconcile(&notices[0], 1)
	future := notification{f.gameID, 999}
	futureRest := f.reconcile(&future, 1)
	f.report("NOT-01-delivery-commit-before-publication", 1, map[string]any{"delivered": notices, "rest": rest, "untrustedFutureNoticeRevision": 999, "restAfterUntrustedNotice": futureRest})
}
func TestNotificationLossStillRecoversREST(t *testing.T) {
	f := newFixture(t, "loss")
	f.apply("lost-notice", 0)
	notices := f.poll()
	if len(notices) != 0 {
		t.Fatal("loss injection did not lose notice")
	}
	rest := f.reconcile(nil, 1)
	f.report("NOT-02-lost-notice", 1, map[string]any{"delivered": notices, "rest": rest})
}
func TestDuplicateNotificationsDoNotRepeatEffects(t *testing.T) {
	f := newFixture(t, "duplicate")
	f.apply("duplicate-notice", 0)
	notices := f.poll()
	if len(notices) != 2 || notices[0] != notices[1] {
		t.Fatal("duplicate injection did not deliver two equal hints")
	}
	var states []game.State
	for i := range notices {
		states = append(states, f.reconcile(&notices[i], 1))
	}
	f.report("NOT-03-duplicate-notice", 1, map[string]any{"delivered": notices, "restAfterEach": states})
}
func TestOutOfOrderNotificationsRecoverLatestREST(t *testing.T) {
	f := newFixture(t, "reverse")
	f.apply("order-one", 0)
	f.apply("order-two", 1)
	notices := f.poll()
	if len(notices) != 2 || notices[0].Revision != 2 || notices[1].Revision != 1 {
		t.Fatal("out-of-order injection did not deliver revisions 2 then 1")
	}
	var states []game.State
	for i := range notices {
		states = append(states, f.reconcile(&notices[i], 2))
	}
	f.report("NOT-04-out-of-order-notices", 2, map[string]any{"delivered": notices, "restAfterEach": states})
}
func TestLostHTTPResponseRecoversAndRetriesOneEffect(t *testing.T) {
	f := newFixture(t, "")
	actual := game.New(f.auth, f.store)
	committedStatus := make(chan int, 1)
	drop := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		actual.ServeHTTP(rec, r)
		committedStatus <- rec.Code
		conn, _, err := w.(http.Hijacker).Hijack()
		if err == nil {
			conn.Close()
		}
	}))
	defer drop.Close()
	body, _ := json.Marshal(game.Intent{OperationID: "lost-http", ExpectedRevision: 0, Payload: "synthetic"})
	req, _ := http.NewRequest("POST", drop.URL+"/games/"+f.gameID+"/operations", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+testkit.Token(f.auth, "user", "game", "alice"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{Timeout: 2 * time.Second}).Do(req)
	if resp != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	if err == nil || <-committedStatus != 200 {
		t.Fatal("controlled response loss did not follow successful commit")
	}
	afterLoss := f.db()
	rest := f.reconcile(nil, 1)
	code, retry, retryErr := f.request("POST", "/games/"+f.gameID+"/operations", game.Intent{OperationID: "lost-http", ExpectedRevision: 0, Payload: "synthetic"})
	if retryErr != nil || code != 200 || retry != (game.State{Revision: 1, Effects: 1}) || afterLoss != (snapshot{1, 1, 1}) {
		t.Fatal("lost response retry duplicated/lost effect")
	}
	f.report("NOT-05-lost-http-response-after-commit", 1, map[string]any{"clientResponseLost": true, "serverCommitStatus": 200, "afterResponseLoss": afterLoss, "rest": rest, "retryStatus": code, "retry": retry})
}
func TestBrokerRestartPreservesConfirmedGameEffects(t *testing.T) {
	f := newFixture(t, "")
	f.apply("before-broker-restart", 0)
	first := f.poll()
	if len(first) != 1 {
		t.Fatal("initial broker delivery missing")
	}
	f.brokerServer.Close()
	f.apply("while-broker-stopped", 1)
	during := f.db()
	stoppedRest := f.reconcile(nil, 2)
	if f.notifyFailures.Load() != 1 {
		t.Fatal("stopped broker did not cause transport failure")
	}
	f.brokerServer = httptest.NewServer(f.broker)
	f.endpoint.Store(f.brokerServer.URL)
	f.apply("after-broker-restart", 2)
	notices := f.poll()
	if len(notices) != 1 || notices[0].Revision != 3 {
		t.Fatal("broker delivery did not resume")
	}
	rest := f.reconcile(&notices[0], 3)
	f.report("NOT-06-broker-restart", 3, map[string]any{"whileBrokerStopped": during, "restWhileBrokerStopped": stoppedRest, "deliveredAfterRestart": notices, "rest": rest})
}
func TestUnresponsiveBrokerRespectsNotifyContextAndCommit(t *testing.T) {
	f := newFixture(t, "unresponsive")
	started := time.Now()
	f.apply("broker-timeout", 0)
	elapsed := time.Since(started)
	if elapsed < 80*time.Millisecond || elapsed > time.Second {
		t.Fatalf("notification timeout not bounded: %s", elapsed)
	}
	select {
	case <-f.broker.cancelled:
	case <-time.After(time.Second):
		t.Fatal("notification HTTP context did not cancel")
	}
	if f.notifyFailures.Load() != 1 {
		t.Fatal("unresponsive broker did not report notification failure")
	}
	rest := f.reconcile(nil, 1)
	f.report("NOT-07-unresponsive-broker-context", 1, map[string]any{"operationElapsedMs": float64(elapsed.Microseconds()) / 1000, "notifyContextCancelled": true, "rest": rest})
}
