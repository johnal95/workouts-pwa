package integration

import (
	"net/http"
	"testing"

	"github.com/johnal95/workouts-pwa/internal/testutil"
	workouthttp "github.com/johnal95/workouts-pwa/internal/workout/transport/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkouts(t *testing.T) {
	a := testutil.NewTestApp(t)

	token, err := a.AuthService.CreateSessionToken(testutil.TestUserID, testutil.TestUserEmail)
	require.NoError(t, err)

	// Create mock workout
	mockWorkout := &workouthttp.CreateWorkoutRequest{
		Name: testutil.Ptr("Mock Workout"),
	}

	var createdWorkout workouthttp.WorkoutResponse

	resp := testutil.DoRequest(a, http.MethodPost, "/api/v1/workouts",
		token, &mockWorkout, &createdWorkout)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, *mockWorkout.Name, createdWorkout.Name)

	// Get workouts after creation
	var workouts []workouthttp.WorkoutResponse

	resp = testutil.DoRequest(a, http.MethodGet, "/api/v1/workouts",
		token, nil, &workouts)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	require.Len(t, workouts, 1)
	assert.Equal(t, createdWorkout.ID, workouts[0].ID)
	assert.Equal(t, createdWorkout.Name, workouts[0].Name)

	// Get specific workout created
	var workout workouthttp.WorkoutResponse

	resp = testutil.DoRequest(a, http.MethodGet, "/api/v1/workouts/"+createdWorkout.ID,
		token, nil, &workout)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, createdWorkout.ID, workout.ID)
	assert.Equal(t, createdWorkout.Name, workout.Name)
}
