package http

import (
	"net/http"
	"time"

	"github.com/johnal95/workouts-pwa/internal/auth"
	"github.com/johnal95/workouts-pwa/internal/httpx"
	"github.com/johnal95/workouts-pwa/internal/logging"
)

type Handler struct {
	parser  *httpx.Parser
	service *auth.Service
}

func NewHandler(parser *httpx.Parser, service *auth.Service) *Handler {
	return &Handler{
		parser:  parser,
		service: service,
	}
}

// Login godoc
//
//	@Summary		Login
//	@Description	Creates a session and sets a session_token cookie
//	@Tags			auth
//	@Produce		json
//	@Success		200	{object}	LoginResponse	"Sets session_token cookie"
//	@Header			200	{string}	Set-Cookie		"session_token JWT cookie"
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	token, err := h.service.CreateSessionToken("019b4388-50ee-7f94-9caf-a8ceb54ef056", "john.doe@gmail.com")
	if err != nil {
		logging.Logger(r.Context()).Error(
			"failed to create session token",
			"error", err,
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     httpx.CookieSessionToken,
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(365 * 24 * time.Hour),
	})

	httpx.RespondJSON(w, http.StatusOK, &LoginResponse{Result: "success"})
}
