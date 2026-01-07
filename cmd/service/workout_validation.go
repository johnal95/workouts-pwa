package service

import (
	"fmt"

	"github.com/johnal95/workouts-pwa/cmd/store"
)

type WorkoutValidationService struct {
}

func NewWorkoutValidationService() *WorkoutValidationService {
	return &WorkoutValidationService{}
}

func (s *WorkoutValidationService) ValidateWorkout(w *store.Workout) error {
	if len(w.Name) < 1 || len(w.Name) > 40 {
		return fmt.Errorf("invalid workout name: must be between 1 and 40 characters")
	}

	return nil
}
