package workout

import "errors"

var (
	ErrWorkoutNameAlreadyExists = errors.New("workout name already exists")
	ErrWorkoutNameInvalid       = errors.New("workout name is not valid")
	ErrWorkoutNotFound          = errors.New("workout not found")
)
