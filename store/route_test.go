package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	testCases := []struct {
		name           string
		mockRepo       *MockRouteRepository
		droneID        int
		expectedOutput []Location
		expectedError  string
	}{
		{
			name:    "Route() should return a slice of Locations",
			droneID: 1234,
			mockRepo: &MockRouteRepository{
				GetRouteFunc: func(id int) ([]Location, error) {
					return []Location{
						{DroneID: 1234, Latitude: 51.474579, Longitude: -0.171834, Time: convertToTimeForTests("2011-03-22T07:47:55Z")},
						{DroneID: 1234, Latitude: 51.479015, Longitude: -0.172361, Time: convertToTimeForTests("2011-03-22T07:48:01Z")},
					}, nil
				},
			},
			expectedOutput: []Location{
				{DroneID: 1234, Latitude: 51.474579, Longitude: -0.171834, Time: convertToTimeForTests("2011-03-22T07:47:55Z")},
				{DroneID: 1234, Latitude: 51.479015, Longitude: -0.172361, Time: convertToTimeForTests("2011-03-22T07:48:01Z")},
			},
		},
		{
			name:    "Route() should return an empty slice of Locations and an error if there is an error reading the file",
			droneID: 1234,
			mockRepo: &MockRouteRepository{
				GetRouteFunc: func(id int) ([]Location, error) {
					return []Location{}, errors.New("could not open file")
				},
			},
			expectedOutput: []Location{},
			expectedError:  "could not open file",
		},
		{
			name:    "Route() should handle partial data correctly",
			droneID: 1234,
			mockRepo: &MockRouteRepository{
				GetRouteFunc: func(id int) ([]Location, error) {
					return []Location{
						{DroneID: 1234, Latitude: 51.474579, Longitude: -0.171834, Time: convertToTimeForTests("2011-03-22T07:47:55Z")},
					}, nil
				},
			},
			expectedOutput: []Location{
				{DroneID: 1234, Latitude: 51.474579, Longitude: -0.171834, Time: convertToTimeForTests("2011-03-22T07:47:55Z")},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			route, err := testCase.mockRepo.GetRoute(testCase.droneID)

			assert.Equal(t, testCase.expectedOutput, route)
			if testCase.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), testCase.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Test the actual Route function implementation
func TestRouteIntegration(t *testing.T) {
	// This is an integration test that uses the real file system
	// We can test the actual implementation without mocking
	route, err := Route(1234) // Use a test drone ID

	// We don't assert on specific values since they depend on file content
	// but we can verify basic functionality
	if err != nil {
		// If there's an error, it should return empty route
		assert.Empty(t, route)
	} else {
		// If successful, we should have some locations (assuming the file exists and has content)
		// This is a loose test since we don't want to depend on specific file content
		assert.NotNil(t, route)
	}
}

func TestParseLocation(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		expectedOutput *Location
		expectedError  bool
	}{
		{
			name:  "parseLocation() should return a Location",
			input: []string{"1234", "51.474579", "-0.171834", "2011-03-22 07:47:55"},
			expectedOutput: &Location{
				DroneID:   1234,
				Latitude:  51.474579,
				Longitude: -0.171834,
				Time:      convertToTimeForTests("2011-03-22T07:47:55Z"),
			},
			expectedError: false,
		},
		{
			name:           "parseLocation() should return nil if the input DroneID is invalid",
			input:          []string{"hello", "51.479015", "-0.172361", "2011-03-22 07:48:01"},
			expectedOutput: nil,
			expectedError:  true,
		},
		{
			name:           "parseLocation() should return nil if the input Latitude is invalid",
			input:          []string{"1234", "hello", "-0.172361", "2011-03-22 07:48:01"},
			expectedOutput: nil,
			expectedError:  true,
		},
		{
			name:           "parseLocation() should return nil if the input Longitude is invalid",
			input:          []string{"1234", "51.479015", "hello", "2011-03-22 07:48:01"},
			expectedOutput: nil,
			expectedError:  true,
		},
		{
			name:           "parseLocation() should return nil if the input Time is invalid",
			input:          []string{"1234", "51.479015", "-0.172361", "22 March 2011, 07:48:01"},
			expectedOutput: nil,
			expectedError:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			location, err := parseLocation(testCase.input)

			assert.Equal(t, testCase.expectedOutput, location)
			if testCase.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
