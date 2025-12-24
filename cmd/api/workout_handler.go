package api

import (
	"encoding/json"
	"errors"
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
	userID := *requestcontext.GetUserID(r)

	workouts, err := h.workoutStore.FindAll(userID)
	if err != nil {
		slog.Error("failed to retrieve user workouts.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	WriteJSON(w, http.StatusOK, workouts)
}

func (h *WorkoutHandler) GetWorkoutDetails(w http.ResponseWriter, r *http.Request) {
	userID := *requestcontext.GetUserID(r)
	workoutId := r.PathValue("workoutId")

	workout, err := h.workoutStore.FindById(userID, workoutId)
	if err != nil {
		slog.Error("failed to retrieve user workouts.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	WriteJSON(w, http.StatusOK, workout)
}

func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	userID := *requestcontext.GetUserID(r)

	var workout store.Workout
	json.NewDecoder(r.Body).Decode(&workout)

	newWorkout, err := h.workoutStore.Create(userID, &workout)
	if errors.Is(err, store.ErrUniqueConstraint) {
		slog.Warn("failed to create new user workout.", "error", err)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Conflict"))
		return
	}
	if err != nil {
		slog.Error("failed to create new user workout.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	WriteJSON(w, http.StatusCreated, newWorkout)
}
