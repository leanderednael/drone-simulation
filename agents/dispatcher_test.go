package agents

import (
	"testing"
)

func TestShutDown(t *testing.T) {
	// Given a dispatcher and a drone
	// When the dispatcher flies the drone
	// Then before terminating the dispatcher should shut down the drone correctly
}

func TestRestart(t *testing.T) {
	// Given a dispatcher and a drone
	// When the dispatcher flies the drone and the drone runs out of memory
	// Then the dispatcher should try to restart the drone with wiped memory and continue on the given route
}
