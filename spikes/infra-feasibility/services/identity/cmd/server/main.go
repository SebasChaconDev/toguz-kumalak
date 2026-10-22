package main

import (
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/services/identity/app"
	"log"
)

func main() {
	if e := probe.Listen(app.FromEnv(), "18101"); e != nil {
		log.Fatal("experimental server stopped")
	}
}
