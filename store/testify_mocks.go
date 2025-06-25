package store

import (
	"github.com/stretchr/testify/mock"
)

// MockStationRepositoryTestify uses testify/mock for more advanced mocking
type MockStationRepositoryTestify struct {
	mock.Mock
}

// GetStations mocks the GetStations method
func (m *MockStationRepositoryTestify) GetStations() ([]Station, error) {
	args := m.Called()
	return args.Get(0).([]Station), args.Error(1)
}

// MockRouteRepositoryTestify uses testify/mock for more advanced mocking
type MockRouteRepositoryTestify struct {
	mock.Mock
}

// GetRoute mocks the GetRoute method
func (m *MockRouteRepositoryTestify) GetRoute(id int) ([]Location, error) {
	args := m.Called(id)
	return args.Get(0).([]Location), args.Error(1)
}
