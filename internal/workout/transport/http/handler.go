package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/johnal95/workouts-pwa/internal/httpx"
	"github.com/johnal95/workouts-pwa/internal/requestcontext"
	"github.com/johnal95/workouts-pwa/internal/workout"
)

type Handler struct {
	service *workout.Service
}

func NewHandler(service *workout.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetWorkouts(w http.ResponseWriter, r *http.Request) {
	userID := requestcontext.MustUserID(r)
	workouts, err := h.service.GetWorkouts(r.Context(), userID)
	if err != nil {
		httpx.RespondError(w, err)
		return
	}
	httpx.RespondJSON(w, http.StatusOK, ToWorkoutListResponse(workouts))
}

func (h *Handler) GetWorkoutDetails(w http.ResponseWriter, r *http.Request) {
	userID := requestcontext.MustUserID(r)
	workoutID := r.PathValue("workoutId")

	wo, err := h.service.GetWorkout(r.Context(), userID, workoutID)
	if err != nil {
		if errors.Is(err, workout.ErrWorkoutNotFound) {
			err = httpx.NotFound("workout not found", err)
		}
		httpx.RespondError(w, err)
		return
	}
	httpx.RespondJSON(w, http.StatusOK, ToWorkoutResponse(wo))
}

func (h *Handler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var data CreateWorkoutRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		httpx.RespondError(w, httpx.InvalidRequestBody(err))
		return
	}

	userID := requestcontext.MustUserID(r)
	newWorkout, err := h.service.CreateWorkout(r.Context(), &workout.Workout{
		Name:   data.Name,
		UserID: userID,
	})
	if err != nil {
		if errors.Is(err, workout.ErrWorkoutNameAlreadyExists) {
			err = httpx.Conflict("workout name must be unique", err)
		}
		if errors.Is(err, workout.ErrWorkoutNameInvalid) {
			err = httpx.BadRequest("invalid workout name", err)
		}
		httpx.RespondError(w, err)
		return
	}

	httpx.RespondJSON(w, http.StatusCreated, ToWorkoutResponse(newWorkout))
}

func (h *Handler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	userID := requestcontext.MustUserID(r)
	workoutID := r.PathValue("workoutId")

	if err := h.service.DeleteWorkout(r.Context(), userID, workoutID); err != nil {
		if errors.Is(err, workout.ErrWorkoutNotFound) {
			err = httpx.NotFound("workout not found", err)
		}
		httpx.RespondError(w, err)
		return
	}

	httpx.RespondNoContent(w)
}
