package app

import (
	"bytes"
	"context"
	"errors"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

// Catches wrong S2S audience/principal, missing signature and swallowed failure.
func TestInternalProbe(t *testing.T) {
	a := probe.Auth{UserKey: bytes.Repeat([]byte{1}, 32), ServiceKey: bytes.Repeat([]byte{2}, 32)}
	r := gin.New()
	r.POST("/internal/probe", a.Require("ai", "service", "game"), func(c *gin.Context) { c.Status(200) })
	s := httptest.NewServer(r)
	if e := checkInternal(context.Background(), s.URL, "ai", a); e != nil {
		t.Fatal("valid upstream rejected")
	}
	if e := checkInternal(context.Background(), s.URL, "identity", a); !errors.Is(e, ErrInternalUnavailable) {
		t.Fatal("wrong audience not rejected")
	}
	s.Close()
	if e := checkInternal(context.Background(), s.URL, "ai", a); !errors.Is(e, ErrInternalUnavailable) {
		t.Fatal("upstream outage not controlled")
	}
}
