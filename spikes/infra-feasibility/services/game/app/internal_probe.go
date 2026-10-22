package app

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"net/http"
	"time"
)

var ErrInternalUnavailable = errors.New("experimental upstream unavailable")

func checkInternal(ctx context.Context, target, audience string, a probe.Auth) error {
	// A bounded throwaway S2S probe, never a user-selected proxy destination.
	raw, _ := json.Marshal(probe.Claims{Kind: "service", Subject: "game", Audience: audience, Scope: "probe", Expires: time.Now().Unix() + 60})
	p := base64.RawURLEncoding.EncodeToString(raw)
	mac := hmac.New(sha256.New, a.ServiceKey)
	mac.Write([]byte(p))
	req, e := http.NewRequestWithContext(ctx, "POST", target+"/internal/probe", nil)
	if e != nil {
		return ErrInternalUnavailable
	}
	req.Header.Set("Authorization", "Bearer "+p+"."+base64.RawURLEncoding.EncodeToString(mac.Sum(nil)))
	res, e := (&http.Client{Timeout: 250 * time.Millisecond, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}).Do(req)
	if e != nil {
		return ErrInternalUnavailable
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return ErrInternalUnavailable
	}
	return nil
}
