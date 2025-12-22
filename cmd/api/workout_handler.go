package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/johnal95/workouts-pwa/cmd/store"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

func NewWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
	}
}

func (h *WorkoutHandler) GetWorkouts(w http.ResponseWriter, r *http.Request) {
	workouts, err := h.workoutStore.FindAll("019b4388-50ee-7f94-9caf-a8ceb54ef056")
	if err != nil {
		slog.Error("failed to retrieve user workouts.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(workouts)
	if err != nil {
		slog.Error("failed to write user workouts.", "error", err)
	}
}

func (h *WorkoutHandler) GetWorkoutDetails(w http.ResponseWriter, r *http.Request) {
	workoutId := r.PathValue("workoutId")

	workout, err := h.workoutStore.FindById("019b4388-50ee-7f94-9caf-a8ceb54ef056", workoutId)
	if err != nil {
		slog.Error("failed to retrieve user workouts.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(workout)
	if err != nil {
		slog.Error("failed to write user workouts.", "error", err)
	}
}

func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	json.NewDecoder(r.Body).Decode(&workout)

	newWorkout, err := h.workoutStore.Create("019b4388-50ee-7f94-9caf-a8ceb54ef056", &workout)
	if err != nil {
		slog.Error("failed to create new user workout.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(newWorkout)
	if err != nil {
		slog.Error("failed to write user workout after creation.", "error", err)
	}
}
