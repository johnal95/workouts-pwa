package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/johnal95/workouts-pwa/internal/app"
	"github.com/johnal95/workouts-pwa/static"
)

func SetupRoutesHandler(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Use(app.RequestIDMiddleware.RequestID)
	r.Use(app.LoggingMiddleware.Logger)

	r.Get("/health", app.HealthCheck)
	r.Post("/login", app.AuthHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(app.AuthMiddleware.Authenticate)
		r.Use(app.AuthMiddleware.RequireUser)

		r.Get("/api/v1/exercises", app.ExerciseHandler.GetExercises)
		r.Get("/api/v1/workouts", app.WorkoutHandler.GetWorkouts)
		r.Get("/api/v1/workouts/{workoutId}", app.WorkoutHandler.GetWorkoutDetails)
		r.Post("/api/v1/workouts", app.WorkoutHandler.CreateWorkout)
		r.Delete("/api/v1/workouts/{workoutId}", app.WorkoutHandler.DeleteWorkout)
		r.Post("/api/v1/workouts/{workoutId}/exercises", app.WorkoutHandler.CreateWorkoutExercise)
		r.Put("/api/v1/workouts/{workoutId}/exercises/order", app.WorkoutHandler.UpdateWorkoutExerciseOrder)
		r.Delete("/api/v1/workouts/{workoutId}/exercises/{workoutExerciseId}", app.WorkoutHandler.DeleteWorkoutExercise)
	})

	r.Handle("GET /", http.FileServerFS(static.GetDistFS()))

	return r
}
