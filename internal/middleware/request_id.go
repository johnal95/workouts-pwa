package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/johnal95/workouts-pwa/internal/requestcontext"
)

type RequestIDMiddleware struct {
}

func NewRequestIDMiddleware() *RequestIDMiddleware {
	return &RequestIDMiddleware{}
}

func (m *RequestIDMiddleware) RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		r = requestcontext.SetRequestID(r, requestID)
		next.ServeHTTP(w, r)
	})
}
