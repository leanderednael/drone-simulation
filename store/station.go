package store

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

const stationsFilename = "tube-stations"

// Station defines the name and location of a London tube station
type Station struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// Stations returns a slice of all tube stations
func Stations() ([]Station, error) {
	lines, err := read(stationsFilename)
	if err != nil {
		return []Station{}, err
	}

	var stations []Station
	for _, line := range lines {
		station, err := parseStation(line)
		if err != nil {
			logrus.Debug(fmt.Sprintf("Could not parse station %s: %s\n", line, err))
			continue
		}

		stations = append(stations, *station)
	}

	return stations, nil
}

func parseStation(line []string) (*Station, error) {
	name := line[0]
	latitude, err := strconv.ParseFloat(line[1], 64)
	if err != nil {
		return nil, err
	}
	longitude, err := strconv.ParseFloat(line[2], 64)
	if err != nil {
		return nil, err
	}

	return &Station{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
