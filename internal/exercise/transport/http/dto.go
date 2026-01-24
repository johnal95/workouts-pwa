package http

import "github.com/johnal95/workouts-pwa/internal/exercise"

type ExerciseResponse struct {
	ID          string                `json:"id"`
	Type        exercise.ExerciseType `json:"type"`
	Name        string                `json:"name"`
	Description *string               `json:"description"`
}
