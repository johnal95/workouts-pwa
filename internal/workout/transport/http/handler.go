package http

import (
	"errors"
	"fmt"
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

// GetWorkouts godoc
//
//	@Summary	List workouts
//	@Tags		workouts
//	@Produce	json
//	@Security	sessionCookieAuth
//	@Success	200	{array}		WorkoutResponse
//	@Failure	401	{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts [get]
func (h *Handler) GetWorkouts(w http.ResponseWriter, r *http.Request) {
	userID := requestcontext.MustUserID(r)
	workouts, err := h.service.GetWorkouts(r.Context(), userID)
	if err != nil {
		httpx.RespondError(w, err)
		return
	}
	httpx.RespondJSON(w, http.StatusOK, ToWorkoutListResponse(workouts))
}

// GetWorkoutDetails godoc
//
//	@Summary	Get workout details
//	@Tags		workouts
//	@Produce	json
//	@Security	sessionCookieAuth
//	@Param		workoutId	path		string	true	"Workout ID"
//	@Success	200			{object}	WorkoutResponse
//	@Failure	401			{object}	httpx.ErrorResponse
//	@Failure	404			{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts/{workoutId} [get]
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

// CreateWorkout godoc
//
//	@Summary	Create workout
//	@Tags		workouts
//	@Accept		json
//	@Produce	json
//	@Security	sessionCookieAuth
//	@Param		request	body		CreateWorkoutRequest	true	"Workout payload"
//	@Success	201		{object}	WorkoutResponse
//	@Failure	400		{object}	httpx.ErrorResponse
//	@Failure	401		{object}	httpx.ErrorResponse
//	@Failure	409		{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts [post]
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
		} else if errors.Is(err, workout.ErrWorkoutLimitReached) {
			err = httpx.BadRequest(err, fmt.Sprintf("workout limit reached (max: %d)", workout.MaxWorkoutsPerUser), nil)
		}
		httpx.RespondError(w, err)
		return
	}

	httpx.RespondJSON(w, http.StatusCreated, ToWorkoutResponse(newWorkout))
}

// CreateWorkoutExercise godoc
//
//	@Summary	Create workout exercise
//	@Tags		workouts
//	@Accept		json
//	@Produce	json
//	@Security	sessionCookieAuth
//	@Param		workoutId	path		string							true	"Workout ID"
//	@Param		request		body		CreateWorkoutExerciseRequest	true	"Workout exercise payload"
//	@Success	201			{object}	WorkoutExerciseResponse
//	@Failure	400			{object}	httpx.ErrorResponse
//	@Failure	401			{object}	httpx.ErrorResponse
//	@Failure	404			{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts/{workoutId}/exercises [post]
func (h *Handler) CreateWorkoutExercise(w http.ResponseWriter, r *http.Request) {
	var data CreateWorkoutExerciseRequest
	if err := h.parser.ParseJSON(r.Body, &data); err != nil {
		httpx.RespondError(w, err)
		return
	}

	userID := requestcontext.MustUserID(r)
	workoutID := r.PathValue("workoutId")

	newWorkoutExercise, err := h.service.CreateWorkoutExercise(r.Context(), userID, workoutID, &workout.CreateWorkoutExerciseInput{
		ExerciseID: *data.ExerciseID,
		Notes:      data.Notes,
	})
	if err != nil {
		if errors.Is(err, workout.ErrWorkoutNotFound) {
			err = httpx.NotFound(err, "workout not found", nil)
		}
		httpx.RespondError(w, err)
		return
	}

	httpx.RespondJSON(w, http.StatusCreated, ToWorkoutExerciseResponse(newWorkoutExercise))
}

// UpdateWorkoutExerciseOrder godoc
//
//	@Summary	Update workout exercise order
//	@Tags		workouts
//	@Accept		json
//	@Produce	json
//	@Security	sessionCookieAuth
//	@Param		workoutId	path		string								true	"Workout ID"
//	@Param		request		body		UpdateWorkoutExerciseOrderRequest	true	"Workout exercise order payload"
//	@Success	200			{object}	UpdateWorkoutExerciseOrderResponse
//	@Failure	400			{object}	httpx.ErrorResponse
//	@Failure	401			{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts/{workoutId}/exercises/order [put]
func (h *Handler) UpdateWorkoutExerciseOrder(w http.ResponseWriter, r *http.Request) {
	var data UpdateWorkoutExerciseOrderRequest
	if err := h.parser.ParseJSON(r.Body, &data); err != nil {
		httpx.RespondError(w, err)
		return
	}

	userID := requestcontext.MustUserID(r)
	workoutID := r.PathValue("workoutId")

	updatedOrder, err := h.service.UpdateWorkoutOrder(r.Context(), userID, workoutID, data.WorkoutExerciseIDs)
	if err != nil {
		if errors.Is(err, workout.ErrInvalidWorkoutExerciseIDs) {
			err = httpx.BadRequest(err, "invalid workout exercise IDs", nil)
		}
		httpx.RespondError(w, err)
		return
	}

	httpx.RespondJSON(w, http.StatusOK, UpdateWorkoutExerciseOrderResponse{
		WorkoutExerciseIDs: updatedOrder,
	})
}

// CreateWorkoutLog godoc
//
//	@Summary	Create workout log
//	@Tags		workouts
//	@Accept		json
//	@Produce	json
//	@Security	sessionCookieAuth
//	@Param		workoutId	path		string	true	"Workout ID"
//	@Failure	501			{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts/{workoutId}/logs [post]
func (h *Handler) CreateWorkoutLog(w http.ResponseWriter, r *http.Request) {
	// userID := requestcontext.MustUserID(r)
	// workoutID := r.PathValue("workoutId")

	httpx.RespondError(w, httpx.NotImplemented(nil, "not yet implemented", nil))
}

// CreateWorkoutLogExerciseSetLog godoc
//
//	@Summary	Create workout log exercise set log
//	@Tags		workouts
//	@Accept		json
//	@Produce	json
//	@Security	sessionCookieAuth
//	@Param		workoutId		path		string	true	"Workout ID"
//	@Param		workoutLogId	path		string	true	"Workout Log ID"
//	@Failure	501				{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts/{workoutId}/logs/{workoutLogId}/sets [post]
func (h *Handler) CreateWorkoutLogExerciseSetLog(w http.ResponseWriter, r *http.Request) {
	// userID := requestcontext.MustUserID(r)
	// workoutID := r.PathValue("workoutId")
	// workoutLogId := r.PathValue("workoutLogId")

	httpx.RespondError(w, httpx.NotImplemented(nil, "not yet implemented", nil))
}

// DeleteWorkout godoc
//
//	@Summary	Delete workout
//	@Tags		workouts
//	@Security	sessionCookieAuth
//	@Param		workoutId	path	string	true	"Workout ID"
//	@Success	204
//	@Failure	401	{object}	httpx.ErrorResponse
//	@Failure	404	{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts/{workoutId} [delete]
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

// DeleteWorkoutExercise godoc
//
//	@Summary	Delete workout exercise
//	@Tags		workouts
//	@Security	sessionCookieAuth
//	@Param		workoutId			path	string	true	"Workout ID"
//	@Param		workoutExerciseId	path	string	true	"Workout Exercise ID"
//	@Success	204
//	@Failure	404	{object}	httpx.ErrorResponse
//	@Failure	500	{object}	httpx.ErrorResponse
//	@Router		/api/v1/workouts/{workoutId}/exercises/{workoutExerciseId} [delete]
func (h *Handler) DeleteWorkoutExercise(w http.ResponseWriter, r *http.Request) {
	userID := requestcontext.MustUserID(r)
	workoutID := r.PathValue("workoutId")
	workoutExerciseID := r.PathValue("workoutExerciseId")

	if err := h.service.DeleteWorkoutExercise(r.Context(), userID, workoutID, workoutExerciseID); err != nil {
		if errors.Is(err, workout.ErrWorkoutExerciseNotFound) {
			err = httpx.NotFound(err, "workout exercise not found", nil)
		}
		httpx.RespondError(w, err)
		return
	}

	httpx.RespondNoContent(w)
}
