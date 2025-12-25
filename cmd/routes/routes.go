package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/johnal95/workouts-pwa/cmd/app"
	"github.com/johnal95/workouts-pwa/static"
)

func SetupRoutesHandler(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Use(app.RequestIDMiddleware.RequestID)

	r.Group(func(r chi.Router) {
		r.Use(app.AuthMiddleware.Authenticate)
		r.Use(app.AuthMiddleware.RequireUser)

		r.Get("/api/v1/workouts", app.WorkoutHandler.GetWorkouts)
		r.Get("/api/v1/workouts/{workoutId}", app.WorkoutHandler.GetWorkoutDetails)
		r.Post("/api/v1/workouts", app.WorkoutHandler.CreateWorkout)
		r.Delete("/api/v1/workouts/{workoutId}", app.WorkoutHandler.DeleteWorkout)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Handle("GET /", http.FileServerFS(static.GetDistFS()))

	return r
}
