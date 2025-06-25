package agents

import (
	"drone_simulation/store"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umahmood/haversine"
)

const (
	statusOn          string  = "on"
	statusOff         string  = "off"
	maxVisibilityInKm float64 = 0.35
	maxMemory         int     = 10
	earthRadiusInKm   int     = 6371
	nanoSecsInAnHour  float64 = 2.77778e-13
)

var trafficScores = []string{"HEAVY", "LIGHT", "MODERATE"}

// Drone defines the behaviours of a drone
type Drone interface {
	ID() int
	IsOn() bool
	HasMemory() bool
	Start()
	Move(location, nextLocation store.Location) store.Location
	calculateCurrentSpeed(previousLocation, location store.Location, timeTravelled time.Duration) (speedInKph float64)
	checkTrafficAtNearbyStations(location store.Location, currentSpeedInKph float64)
	ShutDown()
}

// DroneConfig holds configuration for creating a drone
type DroneConfig struct {
	StationRepo store.StationRepository
}

// drone struct with injected dependencies
type drone struct {
	id             int
	status         string
	stations       []store.Station
	trafficReports int
	stationRepo    store.StationRepository
}

// NewDrone returns a new drone
func NewDrone(id int, config DroneConfig) Drone {
	stations, err := config.StationRepo.GetStations()
	if err != nil {
		logrus.WithField("Drone", id).Warn("Could not parse locations of stations")
	}

	return &drone{id, statusOff, stations, 0, config.StationRepo}
}

// NewDroneWithDefaults creates a drone with default dependencies (backward compatible)
func NewDroneWithDefaults(id int) Drone {
	return NewDrone(id, DroneConfig{
		StationRepo: store.DefaultStationRepository{},
	})
}

func (d *drone) ID() int {
	return d.id
}

func (d *drone) IsOn() bool {
	return d.status == statusOn
}

func (d *drone) HasMemory() bool {
	return d.trafficReports <= maxMemory
}

func (d *drone) Start() {
	d.trafficReports = 0
	d.status = statusOn
	logrus.WithField("Drone", d.id).Info("On")
}

func (d *drone) Move(location, nextLocation store.Location) store.Location {
	logger := logrus.WithField("Drone", d.id).WithField("Time", strings.Split(location.Time.String(), " ")[1])

	if d.status == statusOff {
		logger.Error("Off")
		return location
	}

	if d.trafficReports >= maxMemory {
		logger.Error("Out of memory")
		d.ShutDown()
		return location
	}

	logger = logger.WithField("To", fmt.Sprintf("(%f, %f)", nextLocation.Latitude, nextLocation.Longitude))
	if location == nextLocation {
		logger.Info("Lifted off")
	} else {
		logger.Info("Flying")
	}

	travelTime := nextLocation.Time.Sub(location.Time)
	time.Sleep(travelTime)

	previousLocation := location
	location = nextLocation

	d.checkTrafficAtNearbyStations(location, d.calculateCurrentSpeed(previousLocation, location, travelTime))
	return location
}

func (d *drone) calculateCurrentSpeed(previousLocation, location store.Location, timeTravelled time.Duration) (speedInKph float64) {
	_, distanceTravelledInKm := haversine.Distance(
		haversine.Coord{Lat: location.Latitude, Lon: location.Longitude},
		haversine.Coord{Lat: previousLocation.Latitude, Lon: previousLocation.Longitude},
	)
	if timeTravelled == 0 {
		// time travelled less than one full second, but not zero because coordinates have changed slightly
		timeTravelled = time.Nanosecond
	}
	return distanceTravelledInKm / (float64(timeTravelled) * nanoSecsInAnHour)
}

func (d *drone) checkTrafficAtNearbyStations(location store.Location, currentSpeedInKph float64) {
	for _, station := range d.stations {
		_, distanceInKm := haversine.Distance(
			haversine.Coord{Lat: location.Latitude, Lon: location.Longitude},
			haversine.Coord{Lat: station.Latitude, Lon: station.Longitude},
		)

		if distanceInKm <= maxVisibilityInKm {
			logrus.WithField("Drone", d.id).
				WithField("Speed", fmt.Sprintf("%f km/h", currentSpeedInKph)).
				WithField("Station", station.Name).
				WithField("Time", strings.Split(location.Time.String(), " ")[1]).
				WithField("Traffic", trafficScores[rand.Intn(len(trafficScores))]).
				Info("Station in sight")

			d.trafficReports++
		}
	}
}

func (d *drone) ShutDown() {
	d.status = statusOff
	logrus.WithField("Drone", d.id).Info("Off")
}
