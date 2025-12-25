package store

import (
	"time"
)

type AuthProvider string

const (
	AuthProviderGoogle = AuthProvider("GOOGLE")
)

type User struct {
	Id           string       `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	Email        string       `json:"email"`
	AuthProvider AuthProvider `json:"-"`
	AuthId       string       `json:"-"`
}

type Workout struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserId    string    `json:"-"`
	Name      string    `json:"name"`
}

type WorkoutLog struct {
	Id        string    `json:"id"`
	WorkoutId string    `json:"workout_id"`
	Date      time.Time `json:"date"`
}

type Exercise struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	WorkoutId       string `json:"workout_id"`
	DefaultSetCount uint   `json:"default_set_count"`
}

type ExerciseSetLog struct {
	Id            string `json:"id"`
	ExerciseSetId string `json:"exercise_set_id"`
	Index         int    `json:"index"`
	Reps          uint   `json:"reps"`
	Kg            uint   `json:"kg"`
}
