package store

import (
	"errors"
	"time"
)

var errTest = errors.New("could not open file")

func convertToTimeForTests(timeString string) time.Time {
	timeTime, err := time.Parse(timeLayout, timeString)
	if err != nil {
		panic(err)
	}

	return timeTime
}
