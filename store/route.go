package store

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const timeLayout = time.RFC3339

// Location defines the location of a drone at a given time
type Location struct {
	DroneID   int
	Latitude  float64
	Longitude float64
	Time      time.Time
}

// Route returns the route of a drone with given ID as a slice of Locations
func Route(id int) ([]Location, error) {
	lines, err := read(strconv.Itoa(id))
	if err != nil {
		return []Location{}, err
	}

	var locations []Location
	for _, line := range lines {
		location, err := parseLocation(line)
		if err != nil {
			logrus.Debug(fmt.Sprintf("Could not parse location %s: %s\n", line, err))
			continue
		}

		locations = append(locations, *location)
	}

	return locations, nil
}

func parseLocation(line []string) (*Location, error) {
	droneID, err := strconv.Atoi(line[0])
	if err != nil || droneID == 0 {
		return nil, err
	}
	latitude, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		return nil, err
	}
	longitude, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		return nil, err
	}
	time, err := time.Parse(timeLayout, strings.Replace(line[3], " ", "T", 1)+"Z")
	if err != nil {
		return nil, err
	}

	return &Location{
		DroneID:   droneID,
		Latitude:  latitude,
		Longitude: longitude,
		Time:      time,
	}, nil
}
