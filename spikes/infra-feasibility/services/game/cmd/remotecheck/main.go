// Disposable bounded remote verifier; never a product service or load test.
package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Project struct{ Origin, DeploymentURL, OperatorOIDC string }
type Config struct {
	Projects                       map[string]Project
	UserKeys, ServiceKeys, AppDSNs map[string]string
	VerifiedBuildAliases           map[string]string
	Ably                           struct {
		ID   string
		Keys map[string]struct{ Key string }
	}
}

var cfg Config
var mu sync.Mutex
var evidence map[string]any
var records []map[string]any
var cases []map[string]any
var sqlCount int
var output string
var transport = &http.Transport{DisableKeepAlives: true, MaxConnsPerHost: 2}
var client = &http.Client{Timeout: 30 * time.Second, Transport: transport, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}

func save() {
	evidence["records"] = records
	evidence["cases"] = cases
	evidence["sqlFunctionalStatements"] = sqlCount
	evidence["observedUTC"] = time.Now().UTC().Format(time.RFC3339Nano)
	b, _ := json.MarshalIndent(evidence, "", "  ")
	if os.WriteFile(output, b, 0600) != nil {
		panic("evidence-write-failed")
	}
}
func record(r map[string]any) {
	mu.Lock()
	defer mu.Unlock()
	r["ordinal"] = len(records) + 1
	records = append(records, r)
	save()
}
func check(id string, ok bool, data any) {
	mu.Lock()
	cases = append(cases, map[string]any{"id": id, "pass": ok, "observed": data})
	save()
	mu.Unlock()
	if !ok {
		panic(id)
	}
}
func b64(b []byte) string { return base64.RawURLEncoding.EncodeToString(b) }
func sign(payload string, key []byte) string {
	m := hmac.New(sha256.New, key)
	m.Write([]byte(payload))
	return b64(m.Sum(nil))
}
func auth(service, kind, subject, keyService string) string {
	k := cfg.UserKeys[keyService]
	if kind == "service" {
		k = cfg.ServiceKeys[keyService]
	}
	key, _ := base64.StdEncoding.DecodeString(k)
	b, _ := json.Marshal(map[string]any{"kind": kind, "sub": subject, "aud": service, "scope": "probe", "exp": time.Now().Unix() + 240})
	p := b64(b)
	return "Bearer " + p + "." + sign(p, key)
}
func ablyToken(expired, invalid bool) string {
	parts := strings.SplitN(cfg.Ably.Keys["subscription-issuer"].Key, ":", 2)
	h, _ := json.Marshal(map[string]string{"typ": "JWT", "alg": "HS256", "kid": parts[0]})
	now := time.Now().Unix()
	iat, exp := now, now+300
	if expired {
		iat = now - 600
		exp = now - 300
	}
	cap, _ := json.Marshal(map[string][]string{"spike:game:owned-game": {"message-subscribe"}})
	p, _ := json.Marshal(map[string]any{"iat": iat, "exp": exp, "x-ably-clientId": "synthetic-alice", "x-ably-capability": string(cap)})
	raw := b64(h) + "." + b64(p)
	key := parts[1]
	if invalid {
		key += "invalid"
	}
	return raw + "." + sign(raw, []byte(key))
}
func ablyBearer(token string) string {
	return "Bearer " + base64.StdEncoding.EncodeToString([]byte(token))
}
func request(method, target, token, oidc string, body any, fault string) *http.Request {
	var raw []byte
	if body != nil {
		raw, _ = json.Marshal(body)
	}
	r, e := http.NewRequest(method, target, bytes.NewReader(raw))
	if e != nil {
		panic("request-configuration")
	}
	if token != "" {
		r.Header.Set("Authorization", token)
	}
	if oidc != "" {
		r.Header.Set("x-vercel-trusted-oidc-idp-token", oidc)
	}
	if body != nil {
		r.Header.Set("Content-Type", "application/json")
	}
	if fault != "" {
		r.Header.Set("X-Spike-Fault", fault)
	}
	return r
}
func send(label, method, target, token, oidc string, body any, fault string) (int, map[string]any) {
	r := request(method, target, token, oidc, body, fault)
	start := time.Now()
	res, e := client.Do(r)
	row := map[string]any{"label": label, "method": method, "host": r.URL.Host, "path": r.URL.Path, "startedUTC": start.UTC().Format(time.RFC3339Nano), "automaticRetries": 0}
	if e != nil {
		row["status"] = 0
		row["elapsedMs"] = time.Since(start).Milliseconds()
		row["transportError"] = true
		record(row)
		return 0, nil
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(res.Body, 32768))
	var j map[string]any
	_ = json.Unmarshal(raw, &j)
	row["status"] = res.StatusCode
	row["elapsedMs"] = time.Since(start).Milliseconds()
	row["regionHeader"] = res.Header.Get("x-vercel-id")
	if j != nil {
		safe := map[string]any{}
		for _, k := range []string{"service", "build", "engine", "go", "instance", "region", "oidcPresent", "revision", "effects", "state", "publishStatus", "committedBeforePublish", "syntheticRows", "accepted", "baseRevision", "upstreamStatus", "cookieReceived"} {
			if v, ok := j[k]; ok {
				safe[k] = v
			}
		}
		if er, ok := j["error"].(map[string]any); ok {
			safe["errorCode"] = er["code"]
		}
		row["data"] = safe
	}
	record(row)
	return res.StatusCode, j
}
func app(label, service, method, path, kind string, body any, fault string) (int, map[string]any) {
	return send(label, method, cfg.Projects[service].Origin+path, auth(service, kind, map[bool]string{true: "game", false: "alice"}[kind == "service"], service), cfg.Projects[service].OperatorOIDC, body, fault)
}
func num(j map[string]any, k string) int64 {
	if v, ok := j[k].(float64); ok {
		return int64(v)
	}
	return -1
}
func state(j map[string]any) map[string]any {
	if s, ok := j["state"].(map[string]any); ok {
		return s
	}
	return j
}
func status(id string, got, want int) {
	check(id, got == want, map[string]int{"status": got, "expected": want})
}
func sqlConnect(dsn string) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	return pgx.Connect(ctx, dsn)
}
func sqlErrorCode(e error) string {
	var p *pgconn.PgError
	if errors.As(e, &p) {
		return p.Code
	}
	return "connection-error"
}
func access() {
	for _, name := range []string{"identity", "game", "ai", "web"} {
		path := "/version"
		if name == "web" {
			path = "/"
		}
		code, _ := send("public-protection-"+name, "GET", cfg.Projects[name].Origin+path, "", "", nil, "")
		check("PROTECT-"+name, code == 401 || code == 302 || code == 307, map[string]int{"status": code})
		if name != "web" {
			code, j := app("runtime-"+name, name, "GET", path, "user", nil, "")
			status("RUNTIME-"+name, code, 200)
			check("OIDC-RUNTIME-"+name, j["oidcPresent"] == true, j)
		} else {
			code, _ := send("protected-web", "GET", cfg.Projects[name].Origin+path, "", cfg.Projects[name].OperatorOIDC, nil, "")
			status("RUNTIME-web", code, 200)
		}
	}
	for _, name := range []string{"identity", "game", "ai"} {
		code, _ := app("connectivity-"+name, name, "GET", "/connectivity", "user", nil, "")
		want := 200
		if name == "game" {
			want = 204
		}
		status("CONNECT-"+name, code, want)
	}
	for _, target := range []string{"identity", "ai"} {
		code, _ := app("S2S-"+target, "game", "POST", "/internal/remote/call/"+target, "service", nil, "")
		status("S2S-"+target, code, 204)
	}
	code, j := app("S2S-invalid-protection", "game", "POST", "/internal/remote/call/ai", "service", nil, "protection")
	status("S2S-DOWN", code, 503)
	check("S2S-DENIED-AT-PROTECTION", num(j, "upstreamStatus") == 401 || num(j, "upstreamStatus") == 302 || num(j, "upstreamStatus") == 307, j)
	code, _ = app("S2S-recovered", "game", "POST", "/internal/remote/call/ai", "service", nil, "")
	status("S2S-RECOVERY", code, 204)
	for _, item := range []struct {
		id, token, path string
		expected        int
	}{
		{"AUTH-MISSING", "", "/games/owned-game", 401},
		{"AUTH-INVALID", "Bearer controlled-invalid", "/games/owned-game", 401},
		{"AUTH-AUDIENCE", auth("ai", "user", "alice", "game"), "/games/owned-game", 401},
		{"AUTH-FOREIGN", auth("game", "user", "alice", "game"), "/games/foreign-game", 403},
		{"AUTH-AI-FORGE", auth("game", "user", "alice", "ai"), "/games/owned-game", 401},
		{"AUTH-AI-WRITE", auth("game", "service", "ai", "ai"), "/internal/remote/operation", 401},
	} {
		method := "GET"
		if item.id == "AUTH-AI-WRITE" {
			method = "POST"
		}
		code, _ := send(item.id, method, cfg.Projects["game"].Origin+item.path, item.token, cfg.Projects["game"].OperatorOIDC, nil, "")
		status(item.id, code, item.expected)
	}
	for _, name := range []string{"identity", "game"} {
		code, _ := send("WEB-FORWARD-"+name, "GET", cfg.Projects["web"].Origin+"/probe/"+name+"/connectivity", auth(name, "user", "alice", name), cfg.Projects["web"].OperatorOIDC, nil, "")
		want := 200
		if name == "game" {
			want = 204
		}
		status("WEB-FORWARD-"+name, code, want)
	}
	code, _ = send("WEB-FORWARD-unauthorized", "GET", cfg.Projects["web"].Origin+"/probe/game/connectivity", "", cfg.Projects["web"].OperatorOIDC, nil, "")
	status("WEB-FORWARD-AUTH", code, 401)
	databaseAccess()
}

func databaseAccess() {
	for _, name := range []string{"identity", "game"} {
		own, e := sqlConnect(cfg.AppDSNs[name])
		check("DB-CONNECT-"+name, e == nil, nil)
		var role, db string
		var limit int
		sqlCount++
		e = own.QueryRow(context.Background(), "SELECT current_user,current_database(),rolconnlimit FROM pg_roles WHERE rolname=current_user").Scan(&role, &db, &limit)
		check("DB-ROLE-LIMIT-"+name, e == nil && role == "spike_"+name+"_app" && db == "spike_"+name && limit == 4, map[string]any{"role": role, "database": db, "connectionLimit": limit})
		sqlCount++
		_, e = own.Exec(context.Background(), "CREATE TABLE forbidden_probe(id int)")
		check("DB-NO-DDL-"+name, sqlErrorCode(e) == "42501", map[string]string{"sqlstate": sqlErrorCode(e)})
		own.Close(context.Background())
		cross, _ := url.Parse(cfg.AppDSNs[name])
		other := "game"
		if name == "game" {
			other = "identity"
		}
		cross.Path = "/spike_" + other
		wrong, e := sqlConnect(cross.String())
		if wrong != nil {
			wrong.Close(context.Background())
		}
		check("DB-CROSS-DENIED-"+name, sqlErrorCode(e) == "42501", map[string]string{"sqlstate": sqlErrorCode(e)})
	}
}

type stream struct {
	body      io.ReadCloser
	revisions chan int64
	cancel    context.CancelFunc
}

func openStream(label, channel, authorization string) *stream {
	target := "https://main.realtime.ably.net/sse?channels=" + url.QueryEscape(channel) + "&v=1.2"
	r := request("GET", target, authorization, "", nil, "")
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	r = r.WithContext(ctx)
	start := time.Now()
	res, e := (&http.Client{Transport: transport, CheckRedirect: client.CheckRedirect}).Do(r)
	if e != nil {
		cancel()
		record(map[string]any{"label": label, "method": "GET", "host": "main.realtime.ably.net", "path": "/event-stream", "status": 0})
		panic(label)
	}
	record(map[string]any{"label": label, "method": "GET", "host": "main.realtime.ably.net", "path": "/event-stream", "status": res.StatusCode, "elapsedMs": time.Since(start).Milliseconds(), "automaticRetries": 0, "channel": channel})
	if res.StatusCode != 200 {
		raw, _ := io.ReadAll(io.LimitReader(res.Body, 4096))
		var j map[string]any
		_ = json.Unmarshal(raw, &j)
		code := any(nil)
		if er, ok := j["error"].(map[string]any); ok {
			code = er["code"]
		}
		res.Body.Close()
		cancel()
		check(label, false, map[string]any{"status": res.StatusCode, "errorCode": code})
	}
	s := &stream{res.Body, make(chan int64, 16), cancel}
	go func() {
		defer close(s.revisions)
		scanner := bufio.NewScanner(res.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "data:") {
				continue
			}
			var msg struct {
				Name string
				Data string
			}
			if json.Unmarshal([]byte(strings.TrimSpace(strings.TrimPrefix(line, "data:"))), &msg) != nil || msg.Name != "committed" {
				continue
			}
			var n struct {
				GameID   string
				Revision int64
			}
			if json.Unmarshal([]byte(msg.Data), &n) == nil && n.GameID == "owned-game" {
				select {
				case s.revisions <- n.Revision:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return s
}
func (s *stream) close() { s.cancel(); s.body.Close() }
func (s *stream) expect(id string, want int64) {
	select {
	case got, ok := <-s.revisions:
		check(id, ok && got == want, map[string]any{"revision": got, "expected": want})
	case <-time.After(12 * time.Second):
		check(id, false, "no-event-before-deadline")
	}
}
func deniedStream(id, channel, authorization string) {
	r := request("GET", "https://main.realtime.ably.net/sse?channels="+url.QueryEscape(channel)+"&v=1.2", authorization, "", nil, "")
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	r = r.WithContext(ctx)
	res, e := client.Do(r)
	if e != nil {
		check(id, false, "transport-error")
		return
	}
	defer res.Body.Close()
	observed := map[string]any{"httpStatus": res.StatusCode, "method": "GET", "host": "main.realtime.ably.net", "path": "/sse", "label": id}
	scanner := bufio.NewScanner(res.Body)
	event := ""
	var code int64
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "event:") {
			event = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
		}
		if strings.HasPrefix(line, "data:") {
			var j map[string]any
			_ = json.Unmarshal([]byte(strings.TrimSpace(strings.TrimPrefix(line, "data:"))), &j)
			code = num(j, "code")
			if event == "error" {
				observed["event"] = event
				observed["errorCode"] = code
				break
			}
		}
		if res.StatusCode != 200 {
			break
		}
	}
	record(observed)
	check(id, code >= 40100 && code < 40200 || res.StatusCode == 401 || res.StatusCode == 403, observed)
}

func capability() {
	token := ablyToken(false, false)
	s := openStream("CAP-01-granular-subscribe", "spike:game:owned-game", ablyBearer(token))
	defer s.close()
	base := "https://main.realtime.ably.net/channels/" + url.PathEscape("spike:game:owned-game")
	code, _ := send("CAP-06-publisher", "POST", base+"/messages", "Basic "+base64.StdEncoding.EncodeToString([]byte(cfg.Ably.Keys["game-publisher"].Key)), "", map[string]any{"name": "committed", "data": "{\"gameId\":\"owned-game\",\"revision\":0}"}, "")
	status("CAP-06", code, 201)
	s.expect("CAP-01-REAL-EVENT", 0)
	deniedStream("CAP-03-other-channel", "spike:game:foreign-game", ablyBearer(token))
	deniedStream("CAP-07-publisher-subscribe", "spike:game:owned-game", "Basic "+base64.StdEncoding.EncodeToString([]byte(cfg.Ably.Keys["game-publisher"].Key)))
	for _, test := range []struct {
		id, method, target, authorization string
		body                              any
	}{
		{"CAP-02-client-publish", "POST", base + "/messages", ablyBearer(token), map[string]string{"name": "forbidden", "data": "synthetic"}},
		{"CAP-04-expired", "GET", base + "/presence", ablyBearer(ablyToken(true, false)), nil},
		{"CAP-05-signature", "GET", base + "/presence", ablyBearer(ablyToken(false, true)), nil},
		{"CAP-08-missing", "GET", base + "/presence", "", nil},
		{"CAP-09-no-presence-read", "GET", base + "/presence", ablyBearer(token), nil},
		{"CAP-10-no-history", "GET", base + "/messages", ablyBearer(token), nil},
	} {
		code, j := send(test.id, test.method, test.target, test.authorization, "", test.body, "")
		check(test.id, code == 401 || code == 403, map[string]any{"status": code, "providerError": j["error"]})
	}
}
func operation(id string, rev int64, payload, fault string) (int, map[string]any) {
	return app(id, "game", "POST", "/internal/remote/operation", "service", map[string]any{"operationId": id, "expectedRevision": rev, "payload": payload}, fault)
}
func transactions() {
	s := openStream("NOT-subscribe", "spike:game:owned-game", ablyBearer(ablyToken(false, false)))
	defer s.close()
	code, j := operation("remote-once", 0, "first", "")
	status("TX-FIRST", code, 200)
	check("TX-FIRST-STATE", num(state(j), "revision") == 1 && num(state(j), "effects") == 1, j)
	check("NOT-COMMIT-BEFORE-PUBLISH", j["committedBeforePublish"] == true && num(j, "publishStatus") == 201, j)
	s.expect("NOT-REAL-DELIVERY", 1)
	code, j = operation("remote-once", 0, "first", "")
	status("TX-RETRY", code, 200)
	check("TX-RETRY-STATE", num(state(j), "revision") == 1 && num(j, "publishStatus") == 0, j)
	code, _ = operation("remote-once", 0, "changed", "")
	status("TX-CONFLICT", code, 409)
	var wg sync.WaitGroup
	results := make(chan int, 2)
	for _, id := range []string{"remote-race-a", "remote-race-b"} {
		wg.Add(1)
		go func(id string) { defer wg.Done(); code, _ := operation(id, 1, "contended", ""); results <- code }(id)
	}
	wg.Wait()
	close(results)
	success, conflict := 0, 0
	for code := range results {
		if code == 200 {
			success++
		}
		if code == 409 {
			conflict++
		}
	}
	check("TX-CONCURRENT-TWO", success == 1 && conflict == 1, map[string]int{"commits": success, "conflicts": conflict})
	s.expect("NOT-RACE-ONE-EVENT", 2)
	code, _ = operation("remote-rollback", 2, "rollback", "precommit")
	status("TX-ROLLBACK", code, 503)
	code, j = app("REST-after-rollback", "game", "GET", "/games/owned-game", "user", nil, "")
	check("TX-ROLLBACK-UNCHANGED", code == 200 && num(j, "revision") == 2 && num(j, "effects") == 2, j)
	// The proxy response is deliberately discarded. The later retry must recover
	// its receipt without repeating SQL effects; this is a response-loss model.
	code, j = operation("remote-lost-response", 2, "committed-no-reply", "publish")
	status("TX-LOST-RESPONSE-COMMIT", code, 200)
	check("NOT-PUBLISH-FAIL-PRESERVES-COMMIT", num(j, "publishStatus") == 401 && num(state(j), "revision") == 3 && j["committedBeforePublish"] == true, j)
	code, j = operation("remote-lost-response", 2, "committed-no-reply", "")
	check("TX-LOST-RESPONSE-RECOVERY", code == 200 && num(state(j), "revision") == 3 && num(j, "publishStatus") == 0, j)
	code, j = app("REST-recover-missed-notice", "game", "GET", "/games/owned-game", "user", nil, "")
	check("NOT-REST-AUTHORITY", code == 200 && num(j, "revision") == 3 && num(j, "effects") == 3, j)
	retained := num(j, "revision")
	ignored := 0
	for _, rev := range []int64{2, 1, 2} {
		code, _ := app(fmt.Sprintf("NOT-replay-%d", rev), "game", "POST", "/internal/remote/notice", "service", map[string]int64{"revision": rev}, "")
		status(fmt.Sprintf("NOT-PUBLISH-%d", rev), code, 200)
		s.expect(fmt.Sprintf("NOT-RECEIVE-%d", rev), rev)
		if rev <= retained {
			ignored++
		} else {
			retained = rev
		}
	}
	check("NOT-DUPLICATE-OUT-OF-ORDER", retained == 3 && ignored == 3, map[string]any{"retainedRevision": retained, "ignoredNotices": ignored})
	s.close()
	code, _ = operation("remote-offline", 3, "subscriber-offline", "")
	status("NOT-OFFLINE-COMMIT", code, 200)
	resumed := openStream("NOT-resubscribe", "spike:game:owned-game", ablyBearer(ablyToken(false, false)))
	defer resumed.close()
	code, j = app("REST-after-reconnect", "game", "GET", "/games/owned-game", "user", nil, "")
	check("NOT-RECONNECT-REST", code == 200 && num(j, "revision") == 4 && num(j, "effects") == 4, j)
	// Endpoint budgets and authoritative stale-result rejection use real S2S.
	for _, in := range []struct {
		label      string
		work, base int
		want       int
	}{{"AI-current", 2, 4, 200}, {"AI-budget", 201, 4, 503}, {"AI-timeout", 100, 4, 503}, {"AI-stale", 2, 2, 409}} {
		code, j := app(in.label, "game", "POST", "/internal/remote/ai", "service", map[string]int{"work": in.work, "baseRevision": in.base}, "")
		status(in.label, code, in.want)
		if in.label == "AI-budget" {
			check("AI-BUDGET-UPSTREAM", num(j, "upstreamStatus") == 422, j)
		}
		if in.label == "AI-timeout" {
			check("AI-TIMEOUT-UPSTREAM", num(j, "upstreamStatus") == 408, j)
		}
	}
	c, e := sqlConnect(cfg.AppDSNs["game"])
	check("SQL-final-connect", e == nil, nil)
	defer c.Close(context.Background())
	var rev, effects, receipts int64
	sqlCount++
	e = c.QueryRow(context.Background(), "SELECT revision,effects,(SELECT count(*) FROM synthetic_receipts WHERE game_id='owned-game') FROM synthetic_games WHERE id='owned-game'").Scan(&rev, &effects, &receipts)
	check("TX-FINAL-SQL", e == nil && rev == 4 && effects == 4 && receipts == 4, map[string]int64{"revision": rev, "effects": effects, "receipts": receipts})
}
func cold() {
	for _, name := range []string{"game", "ai"} {
		code, j := app("cold-repeat-"+name, name, "GET", "/version", "user", nil, "")
		status("COLD-RUNTIME-"+name, code, 200)
		var prior map[string]any
		for _, row := range records {
			if row["label"] == "runtime-"+name {
				prior, _ = row["data"].(map[string]any)
				break
			}
		}
		oldBuild, _ := prior["build"].(string)
		newBuild, _ := j["build"].(string)
		sameSources := oldBuild == newBuild || (cfg.VerifiedBuildAliases[oldBuild] != "" && cfg.VerifiedBuildAliases[oldBuild] == cfg.VerifiedBuildAliases[newBuild])
		check("COLD-INSTANCE-"+name, prior != nil && j["instance"] != prior["instance"] && sameSources && j["engine"] == prior["engine"], map[string]any{"priorInstance": prior["instance"], "current": j, "sameSourcesVerified": sameSources, "canonicalSourceDigest": cfg.VerifiedBuildAliases[newBuild]})
	}
	code, j := operation("remote-once", 0, "first", "")
	check("TX-RECEIPT-AFTER-NEW-DEPLOYMENT", code == 200 && num(state(j), "revision") == 1, j)
	code, j = app("REST-after-new-deployment", "game", "GET", "/games/owned-game", "user", nil, "")
	check("TX-STATE-AFTER-NEW-DEPLOYMENT", code == 200 && num(j, "revision") == 4 && num(j, "effects") == 4, j)
}

func assets() {
	secrets := []string{}
	for _, k := range cfg.UserKeys {
		secrets = append(secrets, k)
	}
	for _, k := range cfg.ServiceKeys {
		secrets = append(secrets, k)
	}
	for _, k := range cfg.AppDSNs {
		secrets = append(secrets, k)
	}
	for _, k := range cfg.Ably.Keys {
		secrets = append(secrets, k.Key)
	}
	fetch := func(path string) []byte {
		r := request("GET", cfg.Projects["web"].Origin+path, "", cfg.Projects["web"].OperatorOIDC, nil, "")
		start := time.Now()
		res, e := client.Do(r)
		if e != nil {
			panic("ASSET-transport")
		}
		defer res.Body.Close()
		raw, _ := io.ReadAll(io.LimitReader(res.Body, 3*1024*1024))
		record(map[string]any{"label": "ASSET-SCAN", "method": "GET", "host": r.URL.Host, "path": path, "status": res.StatusCode, "bytes": len(raw), "elapsedMs": time.Since(start).Milliseconds()})
		clean := true
		for _, secret := range secrets {
			if len(secret) > 20 && bytes.Contains(raw, []byte(secret)) {
				clean = false
			}
		}
		check("ASSET-NO-SECRETS-"+path, res.StatusCode == 200 && clean, map[string]any{"bytes": len(raw), "status": res.StatusCode, "secretMatches": !clean})
		return raw
	}
	html := fetch("/")
	matches := regexp.MustCompile(`<script[^>]+src="([^"]+)"`).FindAllSubmatch(html, 20)
	seen := map[string]bool{}
	for _, m := range matches {
		path := string(m[1])
		if !strings.HasPrefix(path, "/_next/static/") || seen[path] {
			continue
		}
		seen[path] = true
		fetch(path)
	}
	check("ASSET-CHUNKS-CHECKED", len(seen) > 0 && len(seen) <= 12, map[string]int{"chunks": len(seen)})
}

func stats() {
	dsn := cfg.AppDSNs["game"]
	c, e := sqlConnect(dsn)
	check("DB-STATS-CONNECT", e == nil, nil)
	defer c.Close(context.Background())
	sqlCount++
	rows, e := c.Query(context.Background(), "SELECT usename,count(*) FROM pg_stat_activity WHERE usename IN ('spike_identity_app','spike_game_app') GROUP BY usename")
	check("DB-ACTIVITY-QUERY", e == nil, nil)
	defer rows.Close()
	counts := map[string]int{}
	total := 0
	within := true
	for rows.Next() {
		var role string
		var n int
		_ = rows.Scan(&role, &n)
		counts[role] = n
		total += n
		if n > 4 {
			within = false
		}
	}
	check("DB-BACKEND-CONNECTIONS", within && total <= 8, map[string]any{"observed": counts, "total": total, "limitPerRole": 4, "aggregateCeiling": 8, "scope": "PostgreSQL backends, not PgBouncer client sockets"})
}

func dataset() {
	for _, name := range []string{"identity", "game"} {
		c, e := sqlConnect(cfg.AppDSNs[name])
		check("DATASET-CONNECT-"+name, e == nil, nil)
		var relationBytes, databaseBytes int64
		sql := "SELECT pg_total_relation_size('synthetic_identity'),pg_database_size(current_database())"
		if name == "game" {
			sql = "SELECT pg_total_relation_size('synthetic_games')+pg_total_relation_size('synthetic_receipts'),pg_database_size(current_database())"
		}
		sqlCount++
		e = c.QueryRow(context.Background(), sql).Scan(&relationBytes, &databaseBytes)
		c.Close(context.Background())
		check("DATASET-BOUNDED-"+name, e == nil && relationBytes < 1024*1024, map[string]any{"relationBytesIncludingIndexes": relationBytes, "physicalDatabaseBytesIncludingPostgreSQLCatalogs": databaseBytes, "syntheticTablesUnder1MiB": relationBytes < 1024*1024})
	}
}
func main() {
	output = os.Getenv("SPIKE_REMOTE_EVIDENCE")
	if output == "" {
		fmt.Println("missing evidence target")
		os.Exit(1)
	}
	evidence = map[string]any{"schemaVersion": 1, "scope": "Remote disposable infrastructure; not product acceptance", "automaticRetries": 0}
	if raw, e := os.ReadFile(output); e == nil {
		_ = json.Unmarshal(raw, &evidence)
		if old, ok := evidence["records"].([]any); ok {
			for _, v := range old {
				records = append(records, v.(map[string]any))
			}
		}
		if old, ok := evidence["cases"].([]any); ok {
			for _, v := range old {
				cases = append(cases, v.(map[string]any))
			}
		}
		if n, ok := evidence["sqlFunctionalStatements"].(float64); ok {
			sqlCount = int(n)
		}
	}
	defer func() {
		if r := recover(); r != nil {
			evidence["lastFailure"] = fmt.Sprint(r)
			save()
			fmt.Println("Remote verifier stopped; sanitized evidence saved:", r)
			os.Exit(1)
		}
	}()
	if json.Unmarshal([]byte(os.Getenv("SPIKE_OPERATOR_CONFIG")), &cfg) != nil {
		panic("operator-config")
	}
	mode := os.Getenv("SPIKE_REMOTE_MODE")
	switch mode {
	case "access":
		access()
	case "database":
		databaseAccess()
	case "capability":
		capability()
	case "transactions":
		transactions()
	case "cold":
		cold()
	case "assets":
		assets()
	case "stats":
		stats()
	case "dataset":
		dataset()
	default:
		panic("mode")
	}
	evidence["lastCompletedMode"] = mode
	save()
	fmt.Printf("Remote verifier mode %s passed; HTTP records=%d, explicit SQL=%d\n", mode, len(records), sqlCount)
}
