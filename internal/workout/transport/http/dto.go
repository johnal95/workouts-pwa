package http

import (
	exercisehttp "github.com/johnal95/workouts-pwa/internal/exercise/transport/http"
)

type WorkoutExerciseResponse struct {
	ID       string                         `json:"id"`
	Exercise *exercisehttp.ExerciseResponse `json:"exercise"`
	Position int                            `json:"position"`
	Notes    *string                        `json:"notes"`
}

type WorkoutResponse struct {
	ID        string                     `json:"id"`
	Name      string                     `json:"name"`
	Exercises []*WorkoutExerciseResponse `json:"exercises"`
}

type CreateWorkoutRequest struct {
	Name *string `json:"name" validate:"required,min=3,max=50"`
}

type CreateWorkoutExerciseRequest struct {
	ExerciseID *string `json:"exercise_id" validate:"required"`
	Notes      *string `json:"notes" validate:"max=500"`
}
