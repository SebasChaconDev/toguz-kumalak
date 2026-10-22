// Package testkit is linked by tests only. No HTTP token-minting endpoint exists.
package testkit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
)

func Token(a probe.Auth, kind, aud, sub string) string {
	b, _ := json.Marshal(probe.Claims{Kind: kind, Subject: sub, Audience: aud, Scope: "probe", Expires: time.Now().Unix() + 60})
	p := base64.RawURLEncoding.EncodeToString(b)
	k := a.UserKey
	if kind == "service" {
		k = a.ServiceKey
	}
	h := hmac.New(sha256.New, k)
	h.Write([]byte(p))
	return p + "." + base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
func HTTPRequest(method, path, body, token string) *http.Request {
	r := httptest.NewRequest(method, path, strings.NewReader(body))
	r.Header.Set("Authorization", "Bearer "+token)
	r.Header.Set("Content-Type", "application/json")
	return r
}
