package middleware

import (
	"log/slog"
	"net/http"

	"github.com/johnal95/workouts-pwa/internal/logging"
	"github.com/johnal95/workouts-pwa/internal/requestcontext"
)

type LoggingMiddleware struct {
	baseLogger *slog.Logger
}

func NewLoggingMiddleware(baseLogger *slog.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		baseLogger: baseLogger,
	}
}

func (m *LoggingMiddleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := m.baseLogger

		if requestID, ok := requestcontext.RequestID(r); ok {
			logger = logger.With(
				slog.String("request_id", requestID),
			)
		}

		ctx := logging.WithLogger(r.Context(), logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
