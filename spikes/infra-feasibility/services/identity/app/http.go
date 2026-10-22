package app

import (
	"context"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"os"
	"time"
)

func New(a probe.Auth, p *pgxpool.Pool) http.Handler {
	r := probe.Router("identity", "none", a)
	r.GET("/connectivity", a.Require("identity", "user", ""), func(c *gin.Context) {
		if p == nil {
			c.AbortWithStatus(503)
			return
		}
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		defer cancel()
		var n int
		if p.QueryRow(ctx, "SELECT count(*) FROM synthetic_identity").Scan(&n) != nil {
			c.AbortWithStatus(503)
			return
		}
		c.JSON(200, gin.H{"syntheticRows": n})
	})
	return r
}
func FromEnv() http.Handler {
	a, e := probe.AuthFromEnv()
	if e != nil {
		return http.HandlerFunc(probe.Unavailable)
	}
	p, e := probe.Pool(context.Background(), os.Getenv("IDENTITY_DATABASE_URL"))
	if e != nil {
		return http.HandlerFunc(probe.Unavailable)
	}
	return New(a, p)
}
