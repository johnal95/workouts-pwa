package exercise

import "time"

type ExerciseType string

const (
	ExerciseTypeReps       = ExerciseType("REPS")
	ExerciseTypeRepsWeight = ExerciseType("REPS_WEIGHT")
	ExerciseTypeDuration   = ExerciseType("DURATION")
)

type Exercise struct {
	ID          string
	CreatedAt   time.Time
	Type        ExerciseType
	Name        string
	Description *string
}
