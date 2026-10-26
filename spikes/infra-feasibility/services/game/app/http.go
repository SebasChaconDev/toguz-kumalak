package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/engine"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func New(a probe.Auth, s *Store) http.Handler {
	r := probe.Router("game", engine.Version, a)
	r.POST("/internal/check/:service", a.Require("game", "service", "game"), func(c *gin.Context) {
		targets := map[string]string{"identity": "http://127.0.0.1:18101", "ai": "http://127.0.0.1:18103"}
		target, ok := targets[c.Param("service")]
		if !ok {
			c.AbortWithStatus(400)
			return
		}
		if checkInternal(c.Request.Context(), target, c.Param("service"), a) != nil {
			c.AbortWithStatus(503)
			return
		}
		c.Status(204)
	})
	errorStatus := func(c *gin.Context, e error) {
		switch {
		case errors.Is(e, ErrForbidden):
			c.AbortWithStatus(403)
		case errors.Is(e, ErrConflict), errors.Is(e, ErrStale):
			c.AbortWithStatus(409)
		default:
			c.AbortWithStatus(503)
		}
	}
	available := func(c *gin.Context) bool {
		if s == nil || s.Pool == nil {
			c.AbortWithStatus(503)
			return false
		}
		return true
	}
	r.GET("/connectivity", a.Require("game", "user", ""), func(c *gin.Context) {
		if !available(c) {
			return
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancel()
		if s.Pool.Ping(ctx) != nil {
			c.AbortWithStatus(503)
			return
		}
		c.Status(204)
	})
	r.GET("/games/:id", a.Require("game", "user", ""), func(c *gin.Context) {
		if !available(c) {
			return
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancel()
		p := c.MustGet("principal").(probe.Claims)
		st, e := s.Get(ctx, c.Param("id"), p.Subject)
		if e != nil {
			errorStatus(c, e)
			return
		}
		c.JSON(200, st)
	})
	r.POST("/games/:id/operations", a.Require("game", "user", ""), func(c *gin.Context) {
		if !available(c) {
			return
		}
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 2048)
		d := json.NewDecoder(c.Request.Body)
		d.DisallowUnknownFields()
		var in Intent
		if d.Decode(&in) != nil || d.Decode(&struct{}{}) != io.EOF || len(in.OperationID) < 1 || len(in.OperationID) > 64 || len(in.Payload) > 512 || in.ExpectedRevision < 0 {
			c.AbortWithStatus(400)
			return
		}
		p := c.MustGet("principal").(probe.Claims)
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancel()
		st, e := s.Apply(ctx, c.Param("id"), p.Subject, in)
		if e != nil {
			errorStatus(c, e)
			return
		}
		c.JSON(200, st)
	})
	addRemoteRoutes(r, a, s)
	return r
}
func FromEnv() http.Handler {
	a, e := probe.AuthFromEnv()
	if e != nil {
		return http.HandlerFunc(probe.Unavailable)
	}
	p, e := probe.Pool(context.Background(), os.Getenv("GAME_DATABASE_URL"))
	if e != nil {
		return http.HandlerFunc(probe.Unavailable)
	}
	return New(a, &Store{Pool: p})
}
