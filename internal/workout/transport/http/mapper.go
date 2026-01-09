package http

import (
	"github.com/johnal95/workouts-pwa/internal/workout"
)

func ToExerciseResponse(e *workout.Exercise) *ExerciseResponse {
	return &ExerciseResponse{
		ID:              e.ID,
		Name:            e.Name,
		DefaultSetCount: e.DefaultSetCount,
		MinReps:         e.MinReps,
		MaxReps:         e.MaxReps,
	}
}

func ToWorkoutResponse(w *workout.Workout) *WorkoutResponse {
	exercises := []*ExerciseResponse{}

	for _, e := range w.Exercises {
		exercises = append(exercises, ToExerciseResponse(e))
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
