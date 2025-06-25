package agents

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrone_WithTestHelpers(t *testing.T) {
	helper := NewTestHelper()

	t.Run("drone creation with valid stations", func(t *testing.T) {
		drone := helper.CreateTestDroneWithDefaults(123)
		assert.Equal(t, 123, drone.ID())
		assert.False(t, drone.IsOn())
		assert.True(t, drone.HasMemory())
	})

	t.Run("drone creation with station load error", func(t *testing.T) {
		mockRepo := helper.CreateMockStationRepoWithError(errors.New("file read error"))
		drone := helper.CreateTestDrone(456, mockRepo)

		// Should still create drone even with station load error
		assert.Equal(t, 456, drone.ID())
		assert.False(t, drone.IsOn())
	})

	t.Run("drone creation with empty stations", func(t *testing.T) {
		mockRepo := helper.CreateMockStationRepoEmpty()
		drone := helper.CreateTestDrone(789, mockRepo)

		assert.Equal(t, 789, drone.ID())
		assert.False(t, drone.IsOn())
	})
}
