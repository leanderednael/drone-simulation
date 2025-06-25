package store

// MockStationRepository is a mock implementation for testing
type MockStationRepository struct {
	GetStationsFunc func() ([]Station, error)
}

// GetStations calls the mock function if set, otherwise returns empty slice
func (m *MockStationRepository) GetStations() ([]Station, error) {
	if m.GetStationsFunc != nil {
		return m.GetStationsFunc()
	}
	return []Station{}, nil
}

// MockRouteRepository is a mock implementation for testing
type MockRouteRepository struct {
	GetRouteFunc func(id int) ([]Location, error)
}

// GetRoute calls the mock function if set, otherwise returns empty slice
func (m *MockRouteRepository) GetRoute(id int) ([]Location, error) {
	if m.GetRouteFunc != nil {
		return m.GetRouteFunc(id)
	}
	return []Location{}, nil
}
