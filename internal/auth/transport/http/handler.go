package http

import (
	"net/http"
	"time"

	"github.com/johnal95/workouts-pwa/internal/auth"
	"github.com/johnal95/workouts-pwa/internal/httpx"
	"github.com/johnal95/workouts-pwa/internal/logging"
)

type Handler struct {
	parser *httpx.Parser
}

func NewHandler(parser *httpx.Parser) *Handler {
	return &Handler{
		parser: parser,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	logger := logging.Logger(r.Context())

	token, err := auth.CreateSessionToken("019b4388-50ee-7f94-9caf-a8ceb54ef056", "john.doe@gmail.com")
	if err != nil {
		logger.Error("failed to create session token", "error", err)
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
