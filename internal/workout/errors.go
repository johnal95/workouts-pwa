package workout

import "errors"

var (
	ErrWorkoutNameAlreadyExists = errors.New("workout name already exists")
	ErrWorkoutNotFound          = errors.New("workout not found")
)
