package app

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe/testkit"
	"net/http/httptest"
	"testing"
	"time"
)

// Catches ignored cancellation, unbounded work and missing useful output.
func TestComputeBounds(t *testing.T) {
	t.Run("bounded", func(t *testing.T) {
		v, e := Compute(context.Background(), 2)
		if e != nil || v == "" {
			t.Fatalf("result=%q error=%v", v, e)
		}
	})
	t.Run("work-limit", func(t *testing.T) {
		_, e := Compute(context.Background(), 201)
		if !errors.Is(e, ErrBudget) {
			t.Fatal(e)
		}
	})
	t.Run("deadline", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
		defer cancel()
		start := time.Now()
		_, e := Compute(ctx, 100)
		if !errors.Is(e, context.DeadlineExceeded) || time.Since(start) > 200*time.Millisecond {
			t.Fatalf("deadline not enforced: %v", e)
		}
	})
	t.Run("request-cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, e := Compute(ctx, 1)
		if !errors.Is(e, context.Canceled) {
			t.Fatal(e)
		}
	})
}

func TestCanceledHTTPRequest(t *testing.T) {
	a := probe.Auth{UserKey: bytes.Repeat([]byte{1}, 32), ServiceKey: bytes.Repeat([]byte{2}, 32)}
	h := New(a)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req := testkit.HTTPRequest("POST", "/internal/compute", `{"work":100,"baseRevision":0}`, testkit.Token(a, "service", "ai", "game")).WithContext(ctx)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != 408 {
		t.Fatalf("canceled HTTP work status=%d", w.Code)
	}
}
func TestAIRefusesDatabaseEnvironment(t *testing.T) {
	a := probe.Auth{UserKey: bytes.Repeat([]byte{1}, 32), ServiceKey: bytes.Repeat([]byte{2}, 32)}
	t.Setenv("SPIKE_USER_KEY", base64.StdEncoding.EncodeToString(a.UserKey))
	t.Setenv("SPIKE_SERVICE_KEY", base64.StdEncoding.EncodeToString(a.ServiceKey))
	for _, name := range []string{"DATABASE_URL", "GAME_DATABASE_URL", "IDENTITY_DATABASE_URL"} {
		t.Setenv(name, "")
	}
	request := func() int {
		w := httptest.NewRecorder()
		FromEnv().ServeHTTP(w, testkit.HTTPRequest("GET", "/version", "", testkit.Token(a, "user", "ai", "alice")))
		return w.Code
	}
	if request() != 200 {
		t.Fatal("valid baseline configuration rejected")
	}
	for _, name := range []string{"DATABASE_URL", "GAME_DATABASE_URL", "IDENTITY_DATABASE_URL"} {
		t.Setenv(name, "synthetic-present")
		if request() != 503 {
			t.Fatal("AI accepted DB environment")
		}
		t.Setenv(name, "")
	}
}
