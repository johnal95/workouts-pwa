package workout

import "time"

type Exercise struct {
	ID              string
	CreatedAt       time.Time
	WorkoutID       string
	Name            string
	DefaultSetCount uint
	MinReps         uint
	MaxReps         uint
}

type Workout struct {
	ID        string
	CreatedAt time.Time
	UserID    string
	Name      string
	Exercises []*Exercise
}

type ExerciseSetLog struct {
	ID         string
	ExerciseID string
	Index      int
	Reps       uint
	KG         uint
}

type WorkoutLog struct {
	ID        string
	WorkoutID string
	Date      time.Time
}
