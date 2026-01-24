package workout

import (
	"time"

	"github.com/johnal95/workouts-pwa/internal/exercise"
)

type WorkoutExercise struct {
	ID        string
	WorkoutID string
	Exercise  exercise.Exercise
	Position  int
	Notes     *string
}

type Workout struct {
	ID        string
	CreatedAt time.Time
	UserID    string
	Name      string
	Exercises []*WorkoutExercise
}
