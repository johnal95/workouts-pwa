package middleware

import (
	"net/http"

	"github.com/johnal95/workouts-pwa/internal/requestcontext"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = requestcontext.SetUserID(r, "019b4388-50ee-7f94-9caf-a8ceb54ef056")
		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) RequireUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, exists := requestcontext.UserID(r); exists {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	})
}
