package http

import (
	"time"

	exercisehttp "github.com/johnal95/workouts-pwa/internal/exercise/transport/http"
	"github.com/johnal95/workouts-pwa/internal/workout"
)

func ToWorkoutLogResponse(w *workout.WorkoutLog) *WorkoutLogResponse {
	return &WorkoutLogResponse{
		ID:        w.ID,
		WorkoutID: w.WorkoutID,
		Date:      w.Date.Format(time.DateOnly),
	}
}

func ToWorkoutExerciseResponse(e *workout.WorkoutExercise) *WorkoutExerciseResponse {
	return &WorkoutExerciseResponse{
		ID: e.ID,
		Exercise: &exercisehttp.ExerciseResponse{
			ID:          e.Exercise.ID,
			Type:        e.Exercise.Type,
			Name:        e.Exercise.Name,
			Description: e.Exercise.Description,
		},
		Position: e.Position,
		Notes:    e.Notes,
	}
}

func ToWorkoutResponse(w *workout.Workout) *WorkoutResponse {
	exercises := []*WorkoutExerciseResponse{}

	for _, e := range w.Exercises {
		exercises = append(exercises, ToWorkoutExerciseResponse(e))
	}

	return &WorkoutResponse{
		ID:        w.ID,
		Name:      w.Name,
		Exercises: exercises,
	}
}

func ToWorkoutListResponse(workouts []*workout.Workout) []*WorkoutResponse {
	workoutListResponse := []*WorkoutResponse{}

	for _, w := range workouts {
		workoutListResponse = append(workoutListResponse, ToWorkoutResponse(w))
	}

	return workoutListResponse
}
