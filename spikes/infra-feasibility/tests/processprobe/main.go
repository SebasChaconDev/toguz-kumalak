// Command processprobe is a local test driver, never a product service.
package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe/testkit"
	game "github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/services/game/app"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type snapshot struct {
	Revision int64 `json:"revision"`
	Effects  int64 `json:"effects"`
	Receipts int64 `json:"receipts"`
}
type record struct {
	ID           string  `json:"id"`
	Before       any     `json:"before,omitempty"`
	Action       string  `json:"action"`
	Status       any     `json:"status"`
	After        any     `json:"after,omitempty"`
	Result       string  `json:"result"`
	Milliseconds float64 `json:"milliseconds"`
}

var records []record
var ctx context.Context
var gp *pgxpool.Pool
var auth probe.Auth
var id string

func userKey(service string) []byte {
	m := hmac.New(sha256.New, auth.UserKey)
	m.Write([]byte("spike-user:" + service))
	return m.Sum(nil)
}
func decision(body string, want game.State) {
	var got game.State
	require(json.Unmarshal([]byte(body), &got) == nil && got == want, "returned stored decision differs")
}

func require(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}
func state() snapshot {
	var s snapshot
	e := gp.QueryRow(ctx, "SELECT revision,effects,(SELECT count(*) FROM synthetic_receipts WHERE game_id=$1) FROM synthetic_games WHERE id=$1", id).Scan(&s.Revision, &s.Effects, &s.Receipts)
	require(e == nil, "snapshot query failed")
	return s
}
func request(method, service, path, body, kind, subject, audience string) (int, string) {
	ports := map[string]string{"identity": "18101", "game": "18102", "ai": "18103"}
	r, e := http.NewRequestWithContext(ctx, method, "http://127.0.0.1:"+ports[service]+path, strings.NewReader(body))
	require(e == nil, "request config")
	token := ""
	if kind != "missing" {
		signing := auth
		tokenKind := kind
		if kind == "user" {
			signing.UserKey = userKey(service)
		}
		if kind == "ai-forged-user" {
			signing.UserKey = userKey("ai")
			tokenKind = "user"
		}
		token = testkit.Token(signing, tokenKind, audience, subject)
		if kind == "invalid" {
			token = "invalid"
		}
		r.Header.Set("Authorization", "Bearer "+token)
	}
	r.Header.Set("Content-Type", "application/json")
	res, e := (&http.Client{Timeout: 2 * time.Second}).Do(r)
	require(e == nil, "HTTP request unavailable")
	defer res.Body.Close()
	raw, e := io.ReadAll(io.LimitReader(res.Body, 8192))
	require(e == nil, "HTTP body unavailable")
	text := string(raw)
	if len(token) > 16 {
		require(!strings.Contains(text, token), "credential in response")
	}
	for _, key := range []string{"SPIKE_USER_KEY", "SPIKE_SERVICE_KEY", "GAME_DATABASE_URL", "IDENTITY_DATABASE_URL"} {
		v := os.Getenv(key)
		if v != "" {
			require(!strings.Contains(text, v), "secret in response")
		}
	}
	for _, name := range []string{"GAME_DATABASE_URL", "IDENTITY_DATABASE_URL"} {
		u, e := url.Parse(os.Getenv(name))
		if e == nil && u.User != nil {
			pw, ok := u.User.Password()
			if ok && pw != "" {
				require(!strings.Contains(text, pw) && !strings.Contains(text, url.QueryEscape(pw)), "password in response")
			}
		}
	}
	for _, s := range []string{"identity", "game", "ai"} {
		require(!strings.Contains(text, base64.StdEncoding.EncodeToString(userKey(s))), "audience key in response")
	}
	return res.StatusCode, text
}
func op(in game.Intent) (int, string) {
	b, _ := json.Marshal(in)
	return request("POST", "game", "/games/"+id+"/operations", string(b), "user", "alice", "game")
}
func add(caseID, action string, before, status, after any, start time.Time) {
	records = append(records, record{caseID, before, action, status, after, "PASS", float64(time.Since(start).Microseconds()) / 1000})
}

func prepare() {
	var table *string
	require(gp.QueryRow(ctx, "SELECT to_regclass('public.synthetic_games')::text").Scan(&table) == nil, "schema lookup")
	if table == nil {
		up, e := game.Migrations.ReadFile("migrations/up.sql")
		require(e == nil, "migration source")
		_, e = gp.Exec(ctx, string(up))
		require(e == nil, "game migration")
	}
	_, e := gp.Exec(ctx, "INSERT INTO synthetic_games VALUES($1,'alice',0,0),($2,'bob',0,0)", id, id+"-foreign")
	require(e == nil, "unique synthetic fixture insertion")
	ip, e := probe.Pool(ctx, os.Getenv("IDENTITY_DATABASE_URL"))
	require(e == nil, "identity config")
	defer ip.Close()
	_, e = ip.Exec(ctx, "CREATE TABLE IF NOT EXISTS synthetic_identity(id text PRIMARY KEY); INSERT INTO synthetic_identity VALUES('alice') ON CONFLICT DO NOTHING")
	require(e == nil, "identity migration")
	add("DB-08", "apply missing migrations as dedicated synthetic database owners; no startup migration", nil, "owner roles", state(), time.Now())
}
func transactions() {
	start := time.Now()
	before := state()
	in := game.Intent{OperationID: "first", ExpectedRevision: 0, Payload: "synthetic-x"}
	status, body := op(in)
	decision(body, game.State{Revision: 1, Effects: 1})
	require(status == 200 && state() == (snapshot{1, 1, 1}), "TX-01 first commit")
	add("TX-01", "POST first/expectedRevision=0", before, status, state(), start)
	start = time.Now()
	before = state()
	status, body = op(in)
	decision(body, game.State{Revision: 1, Effects: 1})
	require(status == 200 && state() == before, "TX-02 duplicate")
	add("TX-02", "same operationId+payload+expectedRevision", before, status, state(), start)
	start = time.Now()
	before = state()
	in.Payload = "changed"
	status, _ = op(in)
	require(status == 409 && state() == before, "TX-03 ID reuse")
	add("TX-03", "same ID changed payload", before, status, state(), start)
	start = time.Now()
	before = state()
	var wg sync.WaitGroup
	statuses := make(chan int, 2)
	for _, key := range []string{"race-a", "race-b"} {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
			code, response := op(game.Intent{OperationID: key, ExpectedRevision: 1, Payload: "race"})
			if code == 200 {
				decision(response, game.State{Revision: 2, Effects: 2})
			}
			statuses <- code
		}(key)
	}
	wg.Wait()
	close(statuses)
	counts := map[int]int{}
	for code := range statuses {
		counts[code]++
	}
	require(counts[200] == 1 && counts[409] == 1 && state() == (snapshot{2, 2, 2}), "TX-04 concurrent revision")
	add("TX-04", "two HTTP intents on revision=1", before, counts, state(), start)
	start = time.Now()
	before = state()
	s := &game.Store{Pool: gp, BeforeCommit: func() error { return errors.New("injected precommit") }}
	_, e := s.Apply(ctx, id, "alice", game.Intent{OperationID: "rollback", ExpectedRevision: 2, Payload: "fault"})
	require(e != nil && state() == before, "TX-05 rollback")
	add("TX-05 / DB-06", "injected failure before commit using in-process Game.Store; no HTTP bypass", before, "error injected", state(), start)
	start = time.Now()
	before = state()
	intent := game.Intent{OperationID: "lost-response", ExpectedRevision: 2, Payload: "lost"}
	raw, _ := json.Marshal(intent)
	confirmed := make(chan int, 1)
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code, _ := op(intent)
		confirmed <- code
		conn, _, err := w.(http.Hijacker).Hijack()
		if err == nil {
			conn.Close()
		}
	}))
	res, e := (&http.Client{Timeout: 2 * time.Second}).Post(proxy.URL, "application/json", strings.NewReader(string(raw)))
	if res != nil {
		res.Body.Close()
	}
	proxy.Close()
	require(e != nil && <-confirmed == 200 && state() == (snapshot{3, 3, 3}), "TX-06 response not lost after commit")
	status, body = op(intent)
	decision(body, game.State{Revision: 3, Effects: 3})
	require(status == 200 && state() == (snapshot{3, 3, 3}), "TX-06 retry duplicates effect")
	add("TX-06", "proxy confirms actual Game HTTP commit then drops client socket; retry same intent", before, map[string]any{"client": "transport error", "retry": status}, state(), start)
	add("DB-05", "all state effects and receipts committed atomically in tested transitions", snapshot{0, 0, 0}, "confirmed", state(), time.Now())
	before = state()
	status, body = op(game.Intent{OperationID: "first", ExpectedRevision: 0, Payload: "synthetic-x"})
	decision(body, game.State{Revision: 1, Effects: 1})
	require(status == 200 && state() == before, "late original retry changed state")
	add("TX-02-late", "retry first confirmed intent after newer revisions: original saved decision 1/1 returned", before, status, state(), time.Now())
}
func security() {
	for _, service := range []string{"identity", "game", "ai"} {
		start := time.Now()
		for _, kind := range []string{"missing", "invalid", "user"} {
			aud := service
			if kind == "user" {
				aud = "wrong"
			}
			code, _ := request("GET", service, "/health", "", kind, "alice", aud)
			require(code == 401, "credential boundary")
		}
		code, _ := request("GET", service, "/connectivity", "", "user", "alice", service)
		want := 200
		if service == "game" {
			want = 204
		}
		require(code == want, "service connectivity")
		add("DB-01 / SEC-"+service, "separate configs; missing/invalid/wrong-audience 401; authenticated connectivity", nil, code, nil, start)
	}
	for _, pair := range [][2]string{{os.Getenv("IDENTITY_DATABASE_URL"), "game"}, {os.Getenv("GAME_DATABASE_URL"), "identity"}} {
		cfg, e := pgx.ParseConfig(pair[0])
		require(e == nil, "cross config")
		cfg.Database = pair[1]
		c, e := pgx.ConnectConfig(ctx, cfg)
		if c != nil {
			c.Close(ctx)
		}
		var pe *pgconn.PgError
		require(errors.As(e, &pe) && pe.Code == "42501", "SQL cross ownership")
		add("DB-02 / DB-03", "cross-owner CONNECT to "+pair[1], nil, "42501", nil, time.Now())
	}
	before := state()
	code, _ := request("GET", "game", "/games/"+id+"-foreign", "", "user", "alice", "game")
	require(code == 403, "foreign game disclosure")
	code2, _ := request("POST", "game", "/games/"+id+"/operations", `{"operationId":"ai-forbidden","expectedRevision":3,"payload":"x"}`, "service", "ai", "game")
	require(code2 == 401 && state() == before, "AI mutation authorized")
	add("SEC-resources", "foreign game GET; AI service Game POST; no DB changes", before, []int{code, code2}, state(), time.Now())
	forged, _ := request("POST", "game", "/games/"+id+"/operations", `{"operationId":"forged","expectedRevision":3,"payload":"x"}`, "ai-forged-user", "alice", "game")
	require(forged == 401 && state() == before, "AI verification key could forge Game user")
	add("SEC-key-scope", "AI audience user key attempts to sign Game-user command", before, forged, state(), time.Now())
	for _, service := range []string{"identity", "ai"} {
		code, _ := request("POST", "game", "/internal/check/"+service, "", "service", "game", "game")
		require(code == 204, "Game outbound S2S")
		add("POC-07-"+service, "Game process authenticated POST to "+service+" process", nil, code, nil, time.Now())
	}
	code, text := request("GET", "ai", "/connectivity", "", "user", "alice", "ai")
	require(code == 200 && strings.Contains(text, `"database":"none"`), "AI DB independence")
	add("DB-04", "AI process started without any database environment, authenticated connectivity", nil, code, nil, time.Now())
	code, _ = request("POST", "ai", "/internal/compute", `{"work":2,"baseRevision":2}`, "service", "game", "ai")
	require(code == 200, "bounded AI compute")
	require(errors.Is(game.AcceptSyntheticResult(state().Revision, 2), game.ErrStale), "stale AI result accepted")
	add("AI-stale", "computed baseRevision=2; authoritative REST/SQL revision=3; revision guard rejects", state(), code, state(), time.Now())
	code, _ = request("POST", "ai", "/internal/compute", `{"work":201,"baseRevision":3}`, "service", "game", "ai")
	require(code == 422, "AI work limit")
	code2, _ = request("POST", "ai", "/internal/compute", `{"work":100,"baseRevision":3}`, "service", "game", "ai")
	require(code2 == 408, "AI deadline")
	add("AI-limits", "work=201 rejected; work=100 deadline40ms enforced", nil, []int{code, code2}, nil, time.Now())
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "processprobe FAIL:", r)
			os.Exit(1)
		}
	}()
	mode := flag.String("mode", "snapshot", "prepare, transactions, security, retry, snapshot, internal-up, internal-down")
	flag.Parse()
	id = os.Getenv("SPIKE_RUN_ID")
	require(strings.HasPrefix(id, "integral-"), "dedicated run ID required")
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var e error
	auth, e = probe.AuthFromEnv()
	require(e == nil, "synthetic auth missing")
	gp, e = probe.Pool(ctx, os.Getenv("GAME_DATABASE_URL"))
	require(e == nil, "synthetic DB config missing")
	defer gp.Close()
	switch *mode {
	case "prepare":
		prepare()
	case "transactions":
		transactions()
	case "security":
		security()
	case "snapshot":
		add("DB-07", "persistent snapshot", nil, "read", state(), time.Now())
	case "retry":
		before := state()
		code, body := op(game.Intent{OperationID: "lost-response", ExpectedRevision: 2, Payload: "lost"})
		decision(body, game.State{Revision: 3, Effects: 3})
		require(code == 200 && before == (snapshot{3, 3, 3}) && state() == before, "TX-07 process restart lost receipt")
		add("TX-07", "same confirmed intent after Game process PID changed", before, code, state(), time.Now())
	case "internal-down", "internal-up":
		code, _ := request("POST", "game", "/internal/check/ai", "", "service", "game", "game")
		want := 503
		if *mode == "internal-up" {
			want = 204
		}
		require(code == want, "internal outage status")
		add("SEC-upstream", *mode, nil, code, nil, time.Now())
	default:
		panic("unsupported mode")
	}
	require(json.NewEncoder(os.Stdout).Encode(map[string]any{"result": "PASS", "mode": *mode, "cases": records}) == nil, "evidence output")
}
