package routes

import (
	"net/http"

	"github.com/johnal95/workouts-pwa/cmd/app"
	"github.com/johnal95/workouts-pwa/static"
)

func SetupRoutesHandler(app *app.Application) http.Handler {
	h := http.NewServeMux()

	h.HandleFunc("GET /api/v1/workouts", app.WorkoutHandler.GetWorkouts)
	h.HandleFunc("GET /api/v1/workouts/{workoutId}", app.WorkoutHandler.GetWorkoutDetails)
	h.HandleFunc("POST /api/v1/workouts", app.WorkoutHandler.CreateWorkout)
	h.Handle("GET /", http.FileServerFS(static.GetDistFS()))

	return h
}
