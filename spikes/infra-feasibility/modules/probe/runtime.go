package probe

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var BuildVersion = "local-unpublished"

func init() { gin.SetMode(gin.ReleaseMode) }

// Synthetic HMAC credentials only. This is NOT Identity/JWT/OIDC implementation.
func AuthFromEnv() (Auth, error) {
	u, e := base64.StdEncoding.DecodeString(os.Getenv("SPIKE_USER_KEY"))
	if e != nil {
		return Auth{}, errors.New("experimental auth unavailable")
	}
	s, e := base64.StdEncoding.DecodeString(os.Getenv("SPIKE_SERVICE_KEY"))
	if e != nil || len(u) < 32 || len(s) < 32 || bytes.Equal(u, s) {
		return Auth{}, errors.New("experimental auth unavailable")
	}
	return Auth{u, s}, nil
}

func Router(service, engine string, a Auth) *gin.Engine {
	r := gin.New()
	r.Use(gin.CustomRecoveryWithWriter(io.Discard, func(c *gin.Context, _ any) { c.AbortWithStatus(500) }))
	r.GET("/health", a.Require(service, "user", ""), func(c *gin.Context) { c.JSON(200, gin.H{"status": "experimental", "service": service}) })
	r.GET("/version", a.Require(service, "user", ""), func(c *gin.Context) { c.JSON(200, gin.H{"build": BuildVersion, "engine": engine, "service": service}) })
	r.POST("/internal/probe", a.Require(service, "service", "game"), func(c *gin.Context) { c.JSON(200, gin.H{"authenticated": true, "service": service}) })
	return r
}

func Pool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	if dsn == "" {
		return nil, errors.New("database unavailable")
	}
	cfg, e := pgxpool.ParseConfig(dsn)
	if e != nil {
		return nil, errors.New("database config invalid")
	}
	cfg.MaxConns = 2
	cfg.MinConns = 0
	cfg.MaxConnLifetime = 5 * time.Minute
	cfg.MaxConnIdleTime = 30 * time.Second
	p, e := pgxpool.NewWithConfig(ctx, cfg)
	if e != nil {
		return nil, errors.New("database unavailable")
	}
	return p, nil
}

func Listen(h http.Handler, port string) error {
	return (&http.Server{Addr: "127.0.0.1:" + port, Handler: h, ReadHeaderTimeout: 2 * time.Second, ReadTimeout: 3 * time.Second, WriteTimeout: 3 * time.Second, IdleTimeout: 10 * time.Second, MaxHeaderBytes: 8192}).ListenAndServe()
}

func Unavailable(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "experimental configuration unavailable", 503)
}
