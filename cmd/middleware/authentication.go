package middleware

import (
	"net/http"

	"github.com/johnal95/workouts-pwa/cmd/requestcontext"
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
		userID := requestcontext.GetUserID(r)
		if userID == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
