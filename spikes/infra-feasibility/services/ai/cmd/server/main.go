package main

import (
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/services/ai/app"
	"log"
)

func main() {
	if e := probe.Listen(app.FromEnv(), "18103"); e != nil {
		log.Fatal("experimental server stopped")
	}
}
