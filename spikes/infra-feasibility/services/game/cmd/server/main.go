package main

import (
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/modules/probe"
	"github.com/SebasChaconDev/toguz-kumalak/spikes/infra-feasibility/services/game/app"
	"log"
)

func main() {
	if e := probe.Listen(app.FromEnv(), "18102"); e != nil {
		log.Fatal("experimental server stopped")
	}
}
