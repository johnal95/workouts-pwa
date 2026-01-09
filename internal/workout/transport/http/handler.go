package http

import (
	"errors"
	"net/http"

	"github.com/johnal95/workouts-pwa/internal/httpx"
	"github.com/johnal95/workouts-pwa/internal/requestcontext"
	"github.com/johnal95/workouts-pwa/internal/workout"
)

type Handler struct {
	parser  *httpx.Parser
	service *workout.Service
}

func NewHandler(parser *httpx.Parser, service *workout.Service) *Handler {
	return &Handler{
		parser:  parser,
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
			err = httpx.NotFound(err, "workout not found", nil)
		}
		httpx.RespondError(w, err)
		return
	}
	httpx.RespondJSON(w, http.StatusOK, ToWorkoutResponse(wo))
}

func (h *Handler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var data CreateWorkoutRequest
	if err := h.parser.ParseJSON(r.Body, &data); err != nil {
		httpx.RespondError(w, err)
		return
	}

	userID := requestcontext.MustUserID(r)
	newWorkout, err := h.service.CreateWorkout(r.Context(), userID, &workout.CreateWorkoutInput{
		Name: *data.Name,
	})
	if err != nil {
		if errors.Is(err, workout.ErrWorkoutNameAlreadyExists) {
			err = httpx.Conflict(err, "workout name must be unique", nil)
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
			err = httpx.NotFound(err, "workout not found", nil)
		}
		httpx.RespondError(w, err)
		return
	}

	httpx.RespondNoContent(w)
}
