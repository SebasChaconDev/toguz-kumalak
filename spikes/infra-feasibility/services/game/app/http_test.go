package app

import (
	"bytes"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe/testkit"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Catches trusting users as services or accepting the wrong service principal.
func TestHTTPTrust(t *testing.T) {
	a := probe.Auth{UserKey: bytes.Repeat([]byte{1}, 32), ServiceKey: bytes.Repeat([]byte{2}, 32)}
	h := New(a, nil)
	server := httptest.NewServer(h)
	defer server.Close()
	for _, tc := range []struct {
		kind, aud, sub, path, method string
		want                         int
	}{
		{"user", "game", "alice", "/health", "GET", 200},
		{"user", "game", "alice", "/version", "GET", 200},
		{"user", "wrong", "alice", "/health", "GET", 401},
		{"user", "game", "alice", "/internal/probe", "POST", 401},
		{"service", "game", "ai", "/internal/probe", "POST", 401},
		{"service", "game", "game", "/internal/probe", "POST", 200},
		{"service", "game", "game", "/health", "GET", 401},
	} {
		req, _ := http.NewRequest(tc.method, server.URL+tc.path, nil)
		req.Header.Set("Authorization", "Bearer "+testkit.Token(a, tc.kind, tc.aud, tc.sub))
		res, e := server.Client().Do(req)
		if e != nil {
			t.Fatal(e)
		}
		res.Body.Close()
		if res.StatusCode != tc.want {
			t.Fatalf("%s %s %s status=%d want=%d", tc.kind, tc.sub, tc.path, res.StatusCode, tc.want)
		}
	}
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != 401 {
		t.Fatal("missing credential accepted")
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, testkit.HTTPRequest("GET", "/version", "", testkit.Token(a, "user", "game", "alice")))
	if !strings.Contains(w.Body.String(), "game") {
		t.Fatal("version missing service")
	}
}
