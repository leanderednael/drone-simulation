package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStations_WithoutMonkey(t *testing.T) {
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
						{"Aldgate", 51.474579, -0.171834},
						{"Camden Town", 51.479015, -0.172361},
					}, nil
				},
			},
			expectedOutput: []Station{
				{"Aldgate", 51.474579, -0.171834},
				{"Camden Town", 51.479015, -0.172361},
			},
		},
		{
			name: "Stations() should return error when repository fails",
			mockRepo: &MockStationRepository{
				GetStationsFunc: func() ([]Station, error) {
					return nil, errors.New("file not found")
				},
			},
			expectedOutput: nil,
			expectedError:  "file not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stations, err := tc.mockRepo.GetStations()

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, stations)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, stations)
			}
		})
	}
}
