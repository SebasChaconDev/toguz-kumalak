package probe

// Disposable remote-spike transport. Callers choose from a fixed configured map;
// no client-supplied URL is ever used as a proxy destination.
import (
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func RemoteOrigin(value string) (string, error) {
	u, e := url.Parse(value)
	if e != nil || u.Scheme != "https" || u.User != nil || u.Path != "" || u.RawQuery != "" || u.Fragment != "" || u.Port() != "" {
		return "", errors.New("invalid experimental HTTPS origin")
	}
	allowed := map[string]bool{}
	for _, name := range []string{"web", "identity", "game", "ai"} {
		allowed["tk-spike01-"+name+"-ultra-seneca.vercel.app"] = true
	}
	if !allowed[u.Host] {
		return "", errors.New("unlisted experimental HTTPS origin")
	}
	return u.String(), nil
}

func TargetKey(audience string, fallback []byte) []byte {
	value := os.Getenv("SPIKE_" + strings.ToUpper(audience) + "_SERVICE_KEY")
	if value == "" {
		return fallback
	}
	k, e := base64.StdEncoding.DecodeString(value)
	if e != nil || len(k) < 32 {
		return nil
	}
	return k
}

func ForwardProtection(out, incoming *http.Request) string {
	token := incoming.Header.Get("x-vercel-oidc-token")
	if token == "" {
		token = os.Getenv("VERCEL_OIDC_TOKEN")
	}
	if token == "" {
		return "missing"
	}
	out.Header.Set("x-vercel-trusted-oidc-idp-token", token)
	return "oidc"
}
