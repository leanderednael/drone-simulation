package agents

import (
	"drone_simulation/store"
	"io"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestDrone_WithTestifyMock(t *testing.T) {
	assert := assert.New(t)
	logrus.SetOutput(io.Discard)

	// Create mock using testify/mock
	mockStationRepo := new(store.MockStationRepositoryTestify)

	// Set up expectations
	expectedStations := []store.Station{
		{"Test Station", 51.5074, -0.1278},
	}
	mockStationRepo.On("GetStations").Return(expectedStations, nil)

	// Create drone with mock
	config := DroneConfig{
		StationRepo: mockStationRepo,
	}
	drone := NewDrone(testDroneID, config)

	// Test
	assert.Equal(testDroneID, drone.ID())

	// Verify that the mock was called as expected
	mockStationRepo.AssertExpectations(t)
}

func TestDrone_WithTestifyMockAdvanced(t *testing.T) {
	assert := assert.New(t)
	logrus.SetOutput(io.Discard)

	// Create mock using testify/mock
	mockStationRepo := new(store.MockStationRepositoryTestify)

	// Set up expectations with specific call count
	expectedStations := []store.Station{
		{"Station 1", 51.5074, -0.1278},
		{"Station 2", 51.5174, -0.1378},
	}
	mockStationRepo.On("GetStations").Return(expectedStations, nil).Once()

	// Create drone with mock
	config := DroneConfig{
		StationRepo: mockStationRepo,
	}
	drone := NewDrone(testDroneID, config)

	// Test
	assert.Equal(testDroneID, drone.ID())
	assert.True(drone.IsOn() == false) // Initially off

	// Verify expectations
	mockStationRepo.AssertExpectations(t)

	// Verify specific method was called exactly once
	mockStationRepo.AssertNumberOfCalls(t, "GetStations", 1)
}
