package agents

import (
	"drone_simulation/store"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Dispatcher defines the behaviours of a dispatcher
type Dispatcher interface {
	Fly(drone Drone, wg *sync.WaitGroup)
}

type dispatcher struct {
	shutDownTime *time.Time
}

// NewDispatcher returns a new dispatcher
func NewDispatcher(shutDownTime *time.Time) Dispatcher {
	return &dispatcher{shutDownTime}
}

func (d *dispatcher) Fly(drone Drone, wg *sync.WaitGroup) {
	defer drone.ShutDown()
	defer wg.Done()

	id := drone.ID()
	logger := logrus.WithField("Drone", id)
	route, err := store.Route(id)
	if err != nil {
		logger.Error("Could not parse route, aborting")
		return
	}

	drone.Start()
	currentLocation := route[0]
	for _, nextLocation := range route {
		if d.shutDownTime != nil && d.shutDownTime.Sub(nextLocation.Time) <= 0 {
			return
		}

		location := drone.Move(currentLocation, nextLocation)
		if location == currentLocation {
			if !drone.IsOn() || !drone.HasMemory() {
				logger.Info("Trying to restart")
				drone.Start()
				location = drone.Move(currentLocation, nextLocation)
			}

			if location != nextLocation {
				logger.Error("Restart failed, aborting")
				return
			}
		}

		currentLocation = location
	}
}
