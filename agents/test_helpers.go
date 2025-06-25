package agents

import (
	"drone_simulation/store"
	"io"

	"github.com/sirupsen/logrus"
)

// TestHelper provides utility functions for testing
type TestHelper struct{}

// NewTestHelper creates a new test helper
func NewTestHelper() *TestHelper {
	// Disable logging during tests
	logrus.SetOutput(io.Discard)
	return &TestHelper{}
}

// CreateMockStationRepo creates a mock station repository with test data
func (h *TestHelper) CreateMockStationRepo() *store.MockStationRepository {
	return &store.MockStationRepository{
		GetStationsFunc: func() ([]store.Station, error) {
			return []store.Station{
				{"Test Station 1", 51.5074, -0.1278},
				{"Test Station 2", 51.5174, -0.1378},
			}, nil
		},
	}
}

// CreateMockStationRepoWithError creates a mock that returns an error
func (h *TestHelper) CreateMockStationRepoWithError(err error) *store.MockStationRepository {
	return &store.MockStationRepository{
		GetStationsFunc: func() ([]store.Station, error) {
			return nil, err
		},
	}
}

// CreateMockStationRepoEmpty creates a mock that returns no stations
func (h *TestHelper) CreateMockStationRepoEmpty() *store.MockStationRepository {
	return &store.MockStationRepository{
		GetStationsFunc: func() ([]store.Station, error) {
			return []store.Station{}, nil
		},
	}
}

// CreateTestDrone creates a drone with mocked dependencies for testing
func (h *TestHelper) CreateTestDrone(id int, stationRepo store.StationRepository) Drone {
	config := DroneConfig{
		StationRepo: stationRepo,
	}
	return NewDrone(id, config)
}

// CreateTestDroneWithDefaults creates a drone with default mock dependencies
func (h *TestHelper) CreateTestDroneWithDefaults(id int) Drone {
	return h.CreateTestDrone(id, h.CreateMockStationRepo())
}
