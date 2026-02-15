package http

import (
	exercisehttp "github.com/johnal95/workouts-pwa/internal/exercise/transport/http"
)

type WorkoutExerciseResponse struct {
	ID       string                         `json:"id"`
	Exercise *exercisehttp.ExerciseResponse `json:"exercise"`
	Position int                            `json:"position"`
	Notes    *string                        `json:"notes"`
}

type WorkoutResponse struct {
	ID        string                     `json:"id"`
	Name      string                     `json:"name"`
	Exercises []*WorkoutExerciseResponse `json:"exercises"`
}

type WorkoutLogResponse struct {
	ID        string `json:"id"`
	WorkoutID string `json:"workout_id"`
	Date      string `json:"date"`
}

type CreateWorkoutRequest struct {
	Name *string `json:"name" validate:"required,min=3,max=50"`
}

type CreateWorkoutExerciseRequest struct {
	ExerciseID *string `json:"exercise_id" validate:"required"`
	Notes      *string `json:"notes" validate:"max=500"`
}

type CreateWorkoutLogRequest struct {
	Date *string `json:"date" validate:"required,datetime=2006-01-02"`
}

type UpdateWorkoutExerciseOrderRequest struct {
	WorkoutExerciseIDs []string `json:"workout_exercise_ids" validate:"min=2,max=50,unique,dive,uuid"`
}

type UpdateWorkoutExerciseOrderResponse struct {
	WorkoutExerciseIDs []string `json:"workout_exercise_ids" validate:"min=2,max=50,unique,dive,uuid"`
}
