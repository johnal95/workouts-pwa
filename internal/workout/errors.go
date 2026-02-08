package workout

import "errors"

var (
	ErrWorkoutNameAlreadyExists  = errors.New("workout name already exists")
	ErrWorkoutNotFound           = errors.New("workout not found")
	ErrWorkoutLimitReached       = errors.New("workout limit reached")
	ErrExerciseNameAlreadyExists = errors.New("exercise name already exists for workout")
	ErrInvalidWorkoutExerciseIDs = errors.New("invalid workout exercise IDs")
)
