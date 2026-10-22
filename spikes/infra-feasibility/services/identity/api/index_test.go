package handler

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe/testkit"
	"net/http/httptest"
	"strings"
	"testing"
)

// Invokes the actual exported serverless Handler, with no running DB necessary.
func TestVercelHTTPEntrypoint(t *testing.T) {
	a := probe.Auth{UserKey: make([]byte, 32), ServiceKey: make([]byte, 32)}
	rand.Read(a.UserKey)
	rand.Read(a.ServiceKey)
	t.Setenv("SPIKE_USER_KEY", base64.StdEncoding.EncodeToString(a.UserKey))
	t.Setenv("SPIKE_SERVICE_KEY", base64.StdEncoding.EncodeToString(a.ServiceKey))
	t.Setenv("DATABASE_URL", "")
	t.Setenv("GAME_DATABASE_URL", "")
	t.Setenv("IDENTITY_DATABASE_URL", "")
	t.Setenv("IDENTITY_DATABASE_URL", "postgres://synthetic:local@127.0.0.1:15435/identity?sslmode=disable")
	w := httptest.NewRecorder()
	Handler(w, testkit.HTTPRequest("GET", "/version", "", testkit.Token(a, "user", "identity", "alice")))
	if w.Code != 200 || !strings.Contains(w.Body.String(), `"engine":"none"`) {
		t.Fatalf("handler %d %s", w.Code, w.Body.String())
	}
	for _, secret := range []string{base64.StdEncoding.EncodeToString(a.UserKey), base64.StdEncoding.EncodeToString(a.ServiceKey)} {
		if strings.Contains(w.Body.String(), secret) {
			t.Fatal("secret leaked")
		}
	}
}
