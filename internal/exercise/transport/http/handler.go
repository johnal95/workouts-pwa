package http

import (
	"net/http"

	"github.com/johnal95/workouts-pwa/internal/exercise"
	"github.com/johnal95/workouts-pwa/internal/httpx"
)

type Handler struct {
	parser  *httpx.Parser
	service *exercise.Service
}

func NewHandler(parser *httpx.Parser, service *exercise.Service) *Handler {
	return &Handler{
		parser:  parser,
		service: service,
	}
}

// GetExercises godoc
//
//	@Summary	List exercises
//	@Tags		exercises
//	@Produce	json
//	@Security	sessionCookieAuth
//	@Success	200	{array}		ExerciseResponse
//	@Failure	401	{object}	httpx.ErrorResponse
//	@Router		/api/v1/exercises [get]
func (h *Handler) GetExercises(w http.ResponseWriter, r *http.Request) {
	exercises, err := h.service.GetExercises(r.Context())
	if err != nil {
		httpx.RespondError(w, httpx.InternalServerError(err))
		return
	}
	httpx.RespondJSON(w, http.StatusOK, ToExerciseListResponse(exercises))
}
