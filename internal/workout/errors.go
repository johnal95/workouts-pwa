package workout

import "errors"

var (
	ErrWorkoutNameAlreadyExists  = errors.New("workout name already exists")
	ErrWorkoutNotFound           = errors.New("workout not found")
	ErrWorkoutLimitReached       = errors.New("workout limit of 50 reached")
	ErrExerciseNameAlreadyExists = errors.New("exercise name already exists for workout")
)
