package store

// StationRepository defines methods for accessing station data
type StationRepository interface {
	GetStations() ([]Station, error)
}

// DefaultStationRepository implements StationRepository using file-based storage
type DefaultStationRepository struct{}

// GetStations returns a slice of all tube stations
func (r DefaultStationRepository) GetStations() ([]Station, error) {
	return Stations()
}

// RouteRepository defines methods for accessing route data
type RouteRepository interface {
	GetRoute(id int) ([]Location, error)
}

// DefaultRouteRepository implements RouteRepository using file-based storage
type DefaultRouteRepository struct{}

// GetRoute returns a slice of locations from a route file
func (r DefaultRouteRepository) GetRoute(id int) ([]Location, error) {
	return Route(id)
}
