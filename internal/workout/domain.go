package workout

import (
	"time"

	"github.com/johnal95/workouts-pwa/internal/exercise"
)

type WorkoutLog struct {
	ID        string
	WorkoutID string
	Date      time.Time
}

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
