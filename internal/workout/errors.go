package workout

import "errors"

var (
	ErrInvalidWorkoutExerciseIDs = errors.New("invalid workout exercise IDs")
	ErrWorkoutLimitReached       = errors.New("workout limit reached")
	ErrWorkoutNameAlreadyExists  = errors.New("workout name already exists")
	ErrWorkoutNotFound           = errors.New("workout not found")
)
