// One bounded timing sample per operation, on replacement disposable resources.
package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type config struct {
	MetricsNeon struct{ ID, DSN, Endpoint string }
	MetricsAbly struct {
		ID   string
		Keys map[string]struct{ ID, Key string }
	}
	MetricsPassword string
}

var evidence = map[string]any{"scope": "C6 timing only; operator-to-provider, not Vercel Game timing or product performance", "samplesPerMetric": 1, "automaticHTTPRetries": 0, "sqlStatementsInstrumented": 0, "httpRequests": []map[string]any{}}
var sqlCount int
var target string

func save() {
	evidence["sqlStatementsInstrumented"] = sqlCount
	b, _ := json.MarshalIndent(evidence, "", "  ")
	if os.WriteFile(target, b, 0600) != nil {
		panic("evidence-write")
	}
}
func require(ok bool, label string) {
	if !ok {
		panic(label)
	}
}
func micros(d time.Duration) float64 { return float64(d.Microseconds()) / 1000 }
func enc(b []byte) string            { return base64.RawURLEncoding.EncodeToString(b) }
func main() {
	target = os.Getenv("SPIKE_METRICS_EVIDENCE")
	require(target != "", "evidence-target")
	defer func() {
		if e := recover(); e != nil {
			evidence["status"] = "FAIL"
			evidence["failure"] = fmt.Sprint(e)
			save()
			fmt.Println("Metrics stopped; sanitized evidence saved:", e)
			os.Exit(1)
		}
	}()
	var cfg config
	require(json.Unmarshal([]byte(os.Getenv("SPIKE_OPERATOR_CONFIG")), &cfg) == nil, "config")
	require(cfg.MetricsNeon.ID != "" && cfg.MetricsAbly.ID != "" && len(cfg.MetricsPassword) == 64, "owned-fixtures")
	adminURL, e := url.Parse(cfg.MetricsNeon.DSN)
	require(e == nil && strings.Contains(adminURL.Host, cfg.MetricsNeon.Endpoint), "owned-endpoint-guard")
	evidence["neonProject"] = cfg.MetricsNeon.ID
	evidence["ablyApp"] = cfg.MetricsAbly.ID
	evidence["observedUTC"] = time.Now().UTC().Format(time.RFC3339Nano)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	start := time.Now()
	admin, e := pgx.Connect(ctx, cfg.MetricsNeon.DSN)
	require(e == nil, "direct-admin-connect")
	defer admin.Close(context.Background())
	evidence["directAdminConnectMs"] = micros(time.Since(start))
	for _, sql := range []string{"CREATE TABLE latency_probe(id int PRIMARY KEY,revision bigint NOT NULL)", "INSERT INTO latency_probe VALUES(1,0)", "CREATE ROLE spike_metrics_app LOGIN PASSWORD '" + cfg.MetricsPassword + "' NOSUPERUSER NOCREATEDB NOCREATEROLE CONNECTION LIMIT 2", "REVOKE CREATE ON SCHEMA public FROM PUBLIC", "GRANT USAGE ON SCHEMA public TO spike_metrics_app", "GRANT SELECT,UPDATE ON latency_probe TO spike_metrics_app"} {
		sqlCount++
		_, e = admin.Exec(ctx, sql)
		require(e == nil, "isolated-metrics-DDL")
		save()
	}
	admin.Close(context.Background())
	pooled := *adminURL
	pooled.Host = strings.Replace(adminURL.Host, cfg.MetricsNeon.Endpoint+".", cfg.MetricsNeon.Endpoint+"-pooler.", 1)
	pooled.User = url.UserPassword("spike_metrics_app", cfg.MetricsPassword)
	pc, e := pgxpool.ParseConfig(pooled.String())
	require(e == nil, "pooled-config")
	pc.MaxConns = 2
	pc.MinConns = 0
	pc.MaxConnIdleTime = 30 * time.Second
	pc.MaxConnLifetime = 5 * time.Minute
	pool, e := pgxpool.NewWithConfig(ctx, pc)
	require(e == nil, "pool-init")
	defer pool.Close()
	start = time.Now()
	conn, e := pool.Acquire(ctx)
	require(e == nil, "pooled-acquire")
	evidence["pooledFirstAcquireMs"] = micros(time.Since(start))
	var revision int64
	start = time.Now()
	sqlCount++
	e = conn.QueryRow(ctx, "SELECT revision FROM latency_probe WHERE id=1").Scan(&revision)
	require(e == nil && revision == 0, "first-select")
	evidence["pooledFirstSelectMs"] = micros(time.Since(start))
	start = time.Now()
	sqlCount++
	e = conn.QueryRow(ctx, "SELECT revision FROM latency_probe WHERE id=1").Scan(&revision)
	require(e == nil && revision == 0, "warm-select")
	evidence["pooledWarmSelectMs"] = micros(time.Since(start))
	start = time.Now()
	sqlCount++
	tx, e := conn.Begin(ctx)
	require(e == nil, "begin")
	sqlCount++
	e = tx.QueryRow(ctx, "UPDATE latency_probe SET revision=revision+1 WHERE id=1 RETURNING revision").Scan(&revision)
	require(e == nil && revision == 1, "update")
	sqlCount++
	e = tx.Commit(ctx)
	require(e == nil, "commit")
	commitAt := time.Now()
	evidence["transactionBeginUpdateCommitMs"] = micros(commitAt.Sub(start))
	sqlCount++
	e = conn.QueryRow(ctx, "SELECT revision FROM latency_probe WHERE id=1").Scan(&revision)
	require(e == nil && revision == 1, "postcommit-read")
	conn.Release()
	evidence["sqlRevisionConfirmedBeforePublish"] = revision
	save()
	key := strings.SplitN(cfg.MetricsAbly.Keys["subscription-issuer"].Key, ":", 2)
	require(len(key) == 2, "issuer-key")
	header, _ := json.Marshal(map[string]string{"alg": "HS256", "typ": "JWT", "kid": key[0]})
	capability, _ := json.Marshal(map[string][]string{"spike:game:owned-game": {"message-subscribe"}})
	now := time.Now().Unix()
	claims, _ := json.Marshal(map[string]any{"iat": now, "exp": now + 300, "x-ably-capability": string(capability), "x-ably-clientId": "metrics-synthetic-alice"})
	unsigned := enc(header) + "." + enc(claims)
	mac := hmac.New(sha256.New, []byte(key[1]))
	mac.Write([]byte(unsigned))
	jwt := unsigned + "." + enc(mac.Sum(nil))
	transport := &http.Transport{DisableKeepAlives: true, MaxConnsPerHost: 2}
	client := &http.Client{Transport: transport, Timeout: 15 * time.Second, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	streamCtx, streamCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer streamCancel()
	req, _ := http.NewRequestWithContext(streamCtx, "GET", "https://main.realtime.ably.net/sse?channels=spike%3Agame%3Aowned-game&v=1.2", nil)
	req.Header.Set("Authorization", "Bearer "+base64.StdEncoding.EncodeToString([]byte(jwt)))
	start = time.Now()
	res, e := client.Do(req)
	require(e == nil, "subscribe-transport")
	defer res.Body.Close()
	require(res.StatusCode == 200, "subscribe-status")
	evidence["subscribeHeadersMs"] = micros(time.Since(start))
	evidence["httpRequests"] = []map[string]any{{"method": "GET", "host": "main.realtime.ably.net", "path": "/sse", "status": res.StatusCode}}
	received := make(chan time.Time, 1)
	go func() {
		scanner := bufio.NewScanner(res.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "data:") {
				continue
			}
			var msg struct{ Name, Data string }
			if json.Unmarshal([]byte(strings.TrimSpace(strings.TrimPrefix(line, "data:"))), &msg) != nil || msg.Name != "committed" {
				continue
			}
			var n struct {
				Revision    int64
				Measurement bool
			}
			if json.Unmarshal([]byte(msg.Data), &n) == nil && n.Revision == 1 && n.Measurement {
				received <- time.Now()
				return
			}
		}
	}()
	payload := []byte(`{"name":"committed","data":"{\"gameId\":\"owned-game\",\"revision\":1,\"measurement\":true}"}`)
	pub, _ := http.NewRequest("POST", "https://main.realtime.ably.net/channels/spike%3Agame%3Aowned-game/messages", bytes.NewReader(payload))
	pub.SetBasicAuth(strings.SplitN(cfg.MetricsAbly.Keys["game-publisher"].Key, ":", 2)[0], strings.SplitN(cfg.MetricsAbly.Keys["game-publisher"].Key, ":", 2)[1])
	pub.Header.Set("Content-Type", "application/json")
	publishAt := time.Now()
	ack, e := client.Do(pub)
	require(e == nil, "publish-transport")
	_, _ = io.Copy(io.Discard, io.LimitReader(ack.Body, 4096))
	ack.Body.Close()
	ackAt := time.Now()
	require(ack.StatusCode == 201, "publish-status")
	evidence["httpRequests"] = append(evidence["httpRequests"].([]map[string]any), map[string]any{"method": "POST", "host": "main.realtime.ably.net", "path": "/channels/spike:game:owned-game/messages", "status": ack.StatusCode})
	evidence["publishAckMs"] = micros(ackAt.Sub(publishAt))
	evidence["commitToPublishStartMs"] = micros(publishAt.Sub(commitAt))
	select {
	case got := <-received:
		evidence["publishStartToReceiveMs"] = micros(got.Sub(publishAt))
		evidence["commitToReceiveMs"] = micros(got.Sub(commitAt))
		evidence["receiveBeforePublishAck"] = got.Before(ackAt)
		evidence["commitBeforeReceive"] = got.After(commitAt)
	case <-time.After(10 * time.Second):
		panic("receive-deadline")
	}
	streamCancel()
	res.Body.Close()
	evidence["poolMaxConnections"] = 2
	evidence["maximumSubscribers"] = 1
	evidence["successfulPublications"] = 1
	evidence["payloadBytes"] = len(payload)
	evidence["status"] = "PASS"
	save()
	fmt.Println("C6 timing follow-up PASS; SQL statements instrumented:", sqlCount, "; successful publication1; subscriber1")
}
