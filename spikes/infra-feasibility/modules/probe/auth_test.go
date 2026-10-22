package probe

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
	"time"
)

// Removing signature, audience, expiry, kind or service-subject checks must fail.
func TestTrustBoundaries(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	a := Auth{bytes.Repeat([]byte{1}, 32), bytes.Repeat([]byte{2}, 32)}
	for _, tc := range []struct {
		name, kind, aud, sub, scope string
		expiry                      int64
		key                         []byte
		want                        int
	}{
		{"valid-user", "user", "game", "alice", "probe", 60, a.UserKey, 204},
		{"wrong-audience", "user", "ai", "alice", "probe", 60, a.UserKey, 401},
		{"invalid-signature", "user", "game", "alice", "probe", 60, a.ServiceKey, 401},
		{"expired", "user", "game", "alice", "probe", -1, a.UserKey, 401},
		{"overlong", "user", "game", "alice", "probe", 600, a.UserKey, 401},
		{"wrong-kind", "service", "game", "ai", "probe", 60, a.ServiceKey, 401},
		{"wrong-scope", "user", "game", "alice", "admin", 60, a.UserKey, 401},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.New()
			r.GET("/", a.Require("game", "user", ""), func(c *gin.Context) { c.Status(204) })
			claims, _ := json.Marshal(Claims{tc.kind, tc.sub, tc.aud, tc.scope, time.Now().Unix() + tc.expiry})
			p := base64.RawURLEncoding.EncodeToString(claims)
			h := hmac.New(sha256.New, tc.key)
			h.Write([]byte(p))
			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", "Bearer "+p+"."+base64.RawURLEncoding.EncodeToString(h.Sum(nil)))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Fatalf("status %d want %d", w.Code, tc.want)
			}
		})
	}
	r := gin.New()
	r.GET("/", a.Require("ai", "service", "game"), func(c *gin.Context) { c.Status(204) })
	for _, token := range []string{"", "Bearer nonsense", "Basic invalid"} {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != 401 {
			t.Fatalf("missing/malformed accepted: %d", w.Code)
		}
	}
}
