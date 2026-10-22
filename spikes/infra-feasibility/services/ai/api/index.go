// Package handler adapts Gin to Vercel's legacy Go HTTP entrypoint.
package handler

import (
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/services/ai/app"
	"net/http"
	"sync"
)

var once sync.Once
var router http.Handler

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() { router = app.FromEnv() })
	router.ServeHTTP(w, r)
}
