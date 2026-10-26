package app

// All routes in this file are disposable, explicitly enabled probes. They must
// never become product endpoints. Existing user commands keep their authority.
import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/engine"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/gin-gonic/gin"
)

func remoteCall(c *gin.Context, audience, path string, payload []byte, a probe.Auth, invalidProtection bool) (int, []byte, error) {
	target, e := probe.RemoteOrigin(os.Getenv("SPIKE_" + strings.ToUpper(audience) + "_ORIGIN"))
	if e != nil {
		return 0, nil, e
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 12*time.Second)
	defer cancel()
	r, e := http.NewRequestWithContext(ctx, "POST", target+path, bytes.NewReader(payload))
	if e != nil {
		return 0, nil, e
	}
	claims, _ := json.Marshal(probe.Claims{Kind: "service", Subject: "game", Audience: audience, Scope: "probe", Expires: time.Now().Unix() + 60})
	p := base64.RawURLEncoding.EncodeToString(claims)
	mac := hmac.New(sha256.New, probe.TargetKey(audience, a.ServiceKey))
	mac.Write([]byte(p))
	r.Header.Set("Authorization", "Bearer "+p+"."+base64.RawURLEncoding.EncodeToString(mac.Sum(nil)))
	r.Header.Set("Content-Type", "application/json")
	protection := probe.ForwardProtection(r, c.Request)
	if invalidProtection {
		r.Header.Set("x-vercel-trusted-oidc-idp-token", "controlled-invalid-oidc")
	}
	if protection == "missing" {
		return 0, nil, errors.New("experimental OIDC unavailable")
	}
	client := &http.Client{Timeout: 12 * time.Second, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	res, e := client.Do(r)
	if e != nil {
		return 0, nil, e
	}
	defer res.Body.Close()
	body, e := io.ReadAll(io.LimitReader(res.Body, 16384))
	return res.StatusCode, body, e
}

func publishRemote(ctx context.Context, notice Notice, key string) (int, error) {
	if notice.GameID != "owned-game" {
		return 0, errors.New("unlisted synthetic channel")
	}
	data, _ := json.Marshal(map[string]any{"gameId": notice.GameID, "revision": notice.Revision})
	body, _ := json.Marshal(map[string]any{"name": "committed", "data": string(data)})
	r, _ := http.NewRequestWithContext(ctx, "POST", "https://main.realtime.ably.net/channels/spike%3Agame%3Aowned-game/messages", bytes.NewReader(body))
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(key)))
	r.Header.Set("Content-Type", "application/json")
	res, e := (&http.Client{Timeout: 4 * time.Second, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}).Do(r)
	if e != nil {
		return 0, e
	}
	defer res.Body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(res.Body, 8192))
	if res.StatusCode >= 300 {
		return res.StatusCode, errors.New("experimental publication refused")
	}
	return res.StatusCode, nil
}

func addRemoteRoutes(r *gin.Engine, a probe.Auth, store *Store) {
	if os.Getenv("SPIKE_REMOTE_ENABLED") != "1" {
		return
	}
	r.POST("/internal/remote/call/:service", a.Require("game", "service", "game"), func(c *gin.Context) {
		service := c.Param("service")
		if service != "identity" && service != "ai" {
			c.AbortWithStatus(400)
			return
		}
		code, _, e := remoteCall(c, service, "/internal/probe", nil, a, c.GetHeader("X-Spike-Fault") == "protection")
		if e != nil || code != 200 {
			c.JSON(503, gin.H{"upstreamStatus": code})
			return
		}
		c.Status(204)
	})
	r.POST("/internal/remote/ai", a.Require("game", "service", "game"), func(c *gin.Context) {
		var in struct {
			BaseRevision int64 `json:"baseRevision"`
			Work         int   `json:"work"`
		}
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1024)
		if c.ShouldBindJSON(&in) != nil {
			c.AbortWithStatus(400)
			return
		}
		body, _ := json.Marshal(in)
		code, result, e := remoteCall(c, "ai", "/internal/compute", body, a, false)
		if e != nil || code != 200 {
			c.JSON(503, gin.H{"upstreamStatus": code})
			return
		}
		var returned struct {
			BaseRevision int64 `json:"baseRevision"`
		}
		if json.Unmarshal(result, &returned) != nil {
			c.AbortWithStatus(503)
			return
		}
		st, e := store.Get(c.Request.Context(), "owned-game", "alice")
		if e != nil {
			c.AbortWithStatus(503)
			return
		}
		if AcceptSyntheticResult(st.Revision, returned.BaseRevision) != nil {
			c.AbortWithStatus(409)
			return
		}
		c.JSON(200, gin.H{"accepted": true, "baseRevision": returned.BaseRevision, "engine": engine.Version})
	})
	r.POST("/internal/remote/operation", a.Require("game", "service", "game"), func(c *gin.Context) {
		var in Intent
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 2048)
		if c.ShouldBindJSON(&in) != nil || in.OperationID == "" || len(in.OperationID) > 64 || len(in.Payload) > 512 || in.ExpectedRevision < 0 {
			c.AbortWithStatus(400)
			return
		}
		copy := *store
		if c.GetHeader("X-Spike-Fault") == "precommit" {
			copy.BeforeCommit = func() error { return errors.New("controlled precommit failure") }
		}
		status := 0
		committedBeforePublish := false
		copy.Notify = func(ctx context.Context, n Notice) error {
			st, e := store.Get(ctx, n.GameID, "alice")
			if e != nil || st.Revision != n.Revision {
				return errors.New("commit not observed")
			}
			committedBeforePublish = true
			key := os.Getenv("SPIKE_ABLY_PUBLISH_KEY")
			if c.GetHeader("X-Spike-Fault") == "publish" {
				key += "invalid"
			}
			status, e = publishRemote(ctx, n, key)
			return e
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), 12*time.Second)
		defer cancel()
		st, e := copy.Apply(ctx, "owned-game", "alice", in)
		if e != nil {
			if errors.Is(e, ErrConflict) || errors.Is(e, ErrStale) {
				c.AbortWithStatus(409)
			} else {
				c.AbortWithStatus(503)
			}
			return
		}
		c.JSON(200, gin.H{"state": st, "publishStatus": status, "committedBeforePublish": committedBeforePublish})
	})
	r.POST("/internal/remote/notice", a.Require("game", "service", "game"), func(c *gin.Context) {
		var in struct {
			Revision int64 `json:"revision"`
		}
		if c.ShouldBindJSON(&in) != nil {
			c.AbortWithStatus(400)
			return
		}
		st, e := store.Get(c.Request.Context(), "owned-game", "alice")
		if e != nil || in.Revision < 1 || in.Revision > st.Revision {
			c.AbortWithStatus(409)
			return
		}
		code, e := publishRemote(c.Request.Context(), Notice{"owned-game", in.Revision}, os.Getenv("SPIKE_ABLY_PUBLISH_KEY"))
		if e != nil {
			c.AbortWithStatus(503)
			return
		}
		c.JSON(200, gin.H{"publishStatus": code})
	})
}
