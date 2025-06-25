package agents

import (
	"drone_simulation/store"
	"errors"
	"io"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const testDroneID = 1234

func TestID_WithDependencyInjection(t *testing.T) {
	assert := assert.New(t)

	// Create mock station repository
	mockStationRepo := &store.MockStationRepository{
		GetStationsFunc: func() ([]store.Station, error) {
			return []store.Station{}, nil
		},
	}

	// Create drone with injected dependencies
	config := DroneConfig{
		StationRepo: mockStationRepo,
	}
	drone := NewDrone(testDroneID, config)

	assert.Equal(testDroneID, drone.ID())
}

func TestStart_Move_ShutDown_WithDependencyInjection(t *testing.T) {
	assert := assert.New(t)
	logrus.SetOutput(io.Discard)

	// Mock station repository
	mockStationRepo := &store.MockStationRepository{
		GetStationsFunc: func() ([]store.Station, error) {
			return []store.Station{}, nil
		},
	}

	// Create drone with injected dependencies
	config := DroneConfig{
		StationRepo: mockStationRepo,
	}
	drone := NewDrone(testDroneID, config)

	currentLocation := store.Location{Latitude: 1, Longitude: 2}
	nextLocation := store.Location{Latitude: 3, Longitude: 4}

	// When the drone hasn't been turned on yet
	location := drone.Move(currentLocation, nextLocation)
	// Then it should not have moved
	assert.Equal(currentLocation, location)

	// When the drone is turned on
	drone.Start()
	location = drone.Move(currentLocation, nextLocation)
	// Then it should have moved to the next location
	assert.Equal(nextLocation, location)

	// When the drone has been shut down
	drone.ShutDown()
	location = drone.Move(currentLocation, nextLocation)
	// Then it should not have moved
	assert.Equal(currentLocation, location)
}

// Example of how to test error scenarios with mocks
func TestNewDrone_StationLoadError(t *testing.T) {
	assert := assert.New(t)
	logrus.SetOutput(io.Discard)

	// Mock station repository that returns an error
	mockStationRepo := &store.MockStationRepository{
		GetStationsFunc: func() ([]store.Station, error) {
			return nil, errors.New("failed to load stations")
		},
	}

	config := DroneConfig{
		StationRepo: mockStationRepo,
	}

	// This should not panic even when station loading fails
	drone := NewDrone(testDroneID, config)
	assert.Equal(testDroneID, drone.ID())
}
