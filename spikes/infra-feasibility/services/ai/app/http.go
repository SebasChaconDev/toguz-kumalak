package app

import (
	"context"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/engine"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func New(a probe.Auth) http.Handler {
	r := probe.Router("ai", engine.Version, a)
	r.GET("/connectivity", a.Require("ai", "user", ""), func(c *gin.Context) { c.JSON(200, gin.H{"engine": engine.Digest("connectivity"), "database": "none"}) })
	r.POST("/internal/compute", a.Require("ai", "service", "game"), func(c *gin.Context) {
		var in struct {
			Work         int   `json:"work"`
			BaseRevision int64 `json:"baseRevision"`
		}
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1024)
		if c.ShouldBindJSON(&in) != nil || in.BaseRevision < 0 {
			c.AbortWithStatus(400)
			return
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), 40*time.Millisecond)
		defer cancel()
		result, e := Compute(ctx, in.Work)
		if e != nil {
			if e == ErrBudget {
				c.AbortWithStatus(422)
			} else {
				c.AbortWithStatus(408)
			}
			return
		}
		c.JSON(200, gin.H{"result": result, "baseRevision": in.BaseRevision, "engine": engine.Version})
	})
	return r
}

func FromEnv() http.Handler {
	a, e := probe.AuthFromEnv()
	if e != nil || os.Getenv("DATABASE_URL") != "" || os.Getenv("GAME_DATABASE_URL") != "" || os.Getenv("IDENTITY_DATABASE_URL") != "" {
		return http.HandlerFunc(probe.Unavailable)
	}
	return New(a)
}
