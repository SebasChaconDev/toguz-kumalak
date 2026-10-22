// Package probe contains experimental transport/auth helpers, not product contracts.
package probe

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Claims struct {
	Kind     string `json:"kind"`
	Subject  string `json:"sub"`
	Audience string `json:"aud"`
	Scope    string `json:"scope"`
	Expires  int64  `json:"exp"`
}
type Auth struct{ UserKey, ServiceKey []byte }

func (a Auth) Require(audience, kind, subject string) gin.HandlerFunc {
	return func(c *gin.Context) {
		deny := func() { c.AbortWithStatus(http.StatusUnauthorized) }
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") || len(h) > 2048 {
			deny()
			return
		}
		parts := strings.Split(strings.TrimPrefix(h, "Bearer "), ".")
		if len(parts) != 2 {
			deny()
			return
		}
		k := a.UserKey
		if kind == "service" {
			k = a.ServiceKey
		}
		if len(k) < 32 || hmac.Equal(a.UserKey, a.ServiceKey) {
			deny()
			return
		}
		sig, e := base64.RawURLEncoding.DecodeString(parts[1])
		if e != nil {
			deny()
			return
		}
		mac := hmac.New(sha256.New, k)
		mac.Write([]byte(parts[0]))
		if !hmac.Equal(sig, mac.Sum(nil)) {
			deny()
			return
		}
		payload, e := base64.RawURLEncoding.DecodeString(parts[0])
		if e != nil {
			deny()
			return
		}
		var p Claims
		d := json.NewDecoder(strings.NewReader(string(payload)))
		d.DisallowUnknownFields()
		if d.Decode(&p) != nil || p.Kind != kind || p.Audience != audience || p.Scope != "probe" || p.Subject == "" || p.Expires <= time.Now().Unix() || p.Expires > time.Now().Unix()+300 || (subject != "" && p.Subject != subject) {
			deny()
			return
		}
		c.Set("principal", p)
		c.Next()
	}
}
