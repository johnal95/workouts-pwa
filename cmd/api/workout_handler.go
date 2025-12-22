package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/johnal95/workouts-pwa/cmd/requestcontext"
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
	user := requestcontext.GetUser(r)

	workouts, err := h.workoutStore.FindAll(user.Id)
	if err != nil {
		slog.Error("failed to retrieve user workouts.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	WriteJSON(w, http.StatusOK, workouts)
}

func (h *WorkoutHandler) GetWorkoutDetails(w http.ResponseWriter, r *http.Request) {
	user := requestcontext.GetUser(r)
	workoutId := r.PathValue("workoutId")

	workout, err := h.workoutStore.FindById(user.Id, workoutId)
	if err != nil {
		slog.Error("failed to retrieve user workouts.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	WriteJSON(w, http.StatusOK, workout)
}

func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	user := requestcontext.GetUser(r)

	var workout store.Workout
	json.NewDecoder(r.Body).Decode(&workout)

	newWorkout, err := h.workoutStore.Create(user.Id, &workout)
	if err != nil {
		slog.Error("failed to create new user workout.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	WriteJSON(w, http.StatusCreated, newWorkout)
}
