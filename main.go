package main

import (
	"drone_simulation/agents"
	"sync"
	"time"
)

const (
	shutDownTime = "2011-03-22T08:10:00Z"
	timeLayout   = time.RFC3339
)

var drones = []int{5937, 6043}

func main() {
	shutDownTime, _ := time.Parse(timeLayout, shutDownTime)
	dispatcher := agents.NewDispatcher(&shutDownTime)

	var wg sync.WaitGroup
	for _, id := range drones {
		drone := agents.NewDroneWithDefaults(id)

		go dispatcher.Fly(drone, &wg)
		wg.Add(1)
	}
	wg.Wait()
}
