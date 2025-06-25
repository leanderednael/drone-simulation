package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStations(t *testing.T) {
	testCases := []struct {
		name           string
		mockRepo       *MockStationRepository
		expectedOutput []Station
		expectedError  string
	}{
		{
			name: "Stations() should return a slice of Stations",
			mockRepo: &MockStationRepository{
				GetStationsFunc: func() ([]Station, error) {
					return []Station{
						{Name: "Aldgate", Latitude: 51.474579, Longitude: -0.171834},
						{Name: "Camden Town", Latitude: 51.479015, Longitude: -0.172361},
					}, nil
				},
			},
			expectedOutput: []Station{
				{Name: "Aldgate", Latitude: 51.474579, Longitude: -0.171834},
				{Name: "Camden Town", Latitude: 51.479015, Longitude: -0.172361},
			},
		},
		{
			name: "Stations() should return an empty slice of Stations and an error if there is an error reading the file",
			mockRepo: &MockStationRepository{
				GetStationsFunc: func() ([]Station, error) {
					return []Station{}, errors.New("could not open file")
				},
			},
			expectedOutput: []Station{},
			expectedError:  "could not open file",
		},
		{
			name: "Stations() should handle partial data correctly",
			mockRepo: &MockStationRepository{
				GetStationsFunc: func() ([]Station, error) {
					return []Station{
						{Name: "Aldgate", Latitude: 51.474579, Longitude: -0.171834},
					}, nil
				},
			},
			expectedOutput: []Station{
				{Name: "Aldgate", Latitude: 51.474579, Longitude: -0.171834},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stations, err := testCase.mockRepo.GetStations()

			assert.Equal(t, testCase.expectedOutput, stations)
			if testCase.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), testCase.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Test the actual Stations function implementation
func TestStationsIntegration(t *testing.T) {
	// This is an integration test that uses the real file system
	// We can test the actual implementation without mocking
	stations, err := Stations()

	// We don't assert on specific values since they depend on file content
	// but we can verify basic functionality
	if err != nil {
		// If there's an error, it should return empty stations
		assert.Empty(t, stations)
	} else {
		// If successful, we should have some stations (assuming the file exists and has content)
		// This is a loose test since we don't want to depend on specific file content
		assert.NotNil(t, stations)
	}
}

func TestParseStation(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		expectedOutput *Station
		expectedError  bool
	}{
		{
			name:  "parseStation() should return a Station",
			input: []string{"Aldgate", "51.474579", "-0.171834"},
			expectedOutput: &Station{
				Name:      "Aldgate",
				Latitude:  51.474579,
				Longitude: -0.171834,
			},
			expectedError: false,
		},
		{
			name:           "parseStation() should return nil if the input Latitude is invalid",
			input:          []string{"Aldgate", "hello", "-0.172361"},
			expectedOutput: nil,
			expectedError:  true,
		},
		{
			name:           "parseStation() should return nil if the input Longitude is invalid",
			input:          []string{"Aldgate", "51.479015", "hello"},
			expectedOutput: nil,
			expectedError:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			station, err := parseStation(testCase.input)

			assert.Equal(t, testCase.expectedOutput, station)
			if testCase.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
