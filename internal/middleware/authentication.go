package middleware

import (
	"errors"
	"net/http"

	"github.com/johnal95/workouts-pwa/internal/auth"
	"github.com/johnal95/workouts-pwa/internal/httpx"
	"github.com/johnal95/workouts-pwa/internal/logging"
	"github.com/johnal95/workouts-pwa/internal/requestcontext"
)

type AuthMiddleware struct {
	authService *auth.Service
}

func NewAuthMiddleware(authService *auth.Service) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(httpx.CookieSessionToken)

		if err != nil {
			if !errors.Is(err, http.ErrNoCookie) {
				logging.Logger(r.Context()).Error(
					"failed to read session cookie",
					"error", err,
				)
			}
		} else {
			s, err := m.authService.VerifySessionToken(cookie.Value)
			if err != nil {
				logging.Logger(r.Context()).Warn(
					"failed to verify session cookie",
					"error", err,
				)
			} else {
				r = requestcontext.SetUserID(r, s.UserID)
				ctx := logging.WithLogger(r.Context(),
					logging.Logger(r.Context()).With("user_id", s.UserID),
				)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, exists := requestcontext.UserID(r); exists {
			next.ServeHTTP(w, r)
		} else {
			httpx.RespondError(w, httpx.Unauthorized(nil, "unauthorized", nil))
		}
	})
}
