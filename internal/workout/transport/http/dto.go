package http

type ExerciseResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	DefaultSetCount uint   `json:"default_set_count"`
	MinReps         uint   `json:"min_reps"`
	MaxReps         uint   `json:"max_reps"`
}

type WorkoutResponse struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	Exercises []*ExerciseResponse `json:"exercises"`
}

type CreateWorkoutRequest struct {
	Name *string `json:"name" validate:"required,min=3,max=50"`
}

type CreateExerciseRequest struct {
	Name            *string `json:"name" validate:"required,min=3,max=50"`
	DefaultSetCount *uint   `json:"default_set_count" validate:"required,min=1,max=40"`
	MinReps         *uint   `json:"min_reps" validate:"required,min=1,max=100"`
	MaxReps         *uint   `json:"max_reps" validate:"required,min=1,max=100"`
}
