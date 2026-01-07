package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/johnal95/workouts-pwa/cmd/requestcontext"
	"github.com/johnal95/workouts-pwa/cmd/service"
	"github.com/johnal95/workouts-pwa/cmd/store"
)

type WorkoutHandler struct {
	workoutStore             store.WorkoutStore
	workoutValidationService *service.WorkoutValidationService
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, workoutValidationService *service.WorkoutValidationService) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore:             workoutStore,
		workoutValidationService: workoutValidationService,
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
	workoutID := r.PathValue("workoutId")

	workout, err := h.workoutStore.FindById(userID, workoutID)
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
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		slog.Warn("invalid request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	workout.UserId = userID

	if err := h.workoutValidationService.ValidateWorkout(&workout); err != nil {
		slog.Warn("invalid workout", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	newWorkout, err := h.workoutStore.Create(&workout)
	if err != nil {
		if errors.Is(err, store.ErrUniqueConstraint) {
			slog.Warn("failed to create new user workout.", "error", err)
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("Conflict"))
			return
		}

		slog.Error("failed to create new user workout.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	WriteJSON(w, http.StatusCreated, newWorkout)
}

func (h *WorkoutHandler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	userID := *requestcontext.GetUserID(r)
	workoutID := r.PathValue("workoutId")

	workout, err := h.workoutStore.Delete(userID, workoutID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			slog.Warn("workout not found.", "error", err)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
			return
		}

		slog.Error("failed to create new user workout.", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	WriteJSON(w, http.StatusOK, workout)
}
