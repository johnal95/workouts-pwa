package workout

import "errors"

var (
	ErrInvalidWorkoutExerciseIDs = errors.New("invalid workout exercise IDs")
	ErrWorkoutExerciseNotFound   = errors.New("workout exercise not found")
	ErrWorkoutLimitReached       = errors.New("workout limit reached")
	ErrWorkoutNameAlreadyExists  = errors.New("workout name already exists")
	ErrWorkoutNotFound           = errors.New("workout not found")
)
