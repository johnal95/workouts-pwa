package middleware

import (
	"log/slog"
	"net/http"
	"time"

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

type responseWritter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWritter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (m *LoggingMiddleware) AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWritter{ResponseWriter: w}
		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		logging.Logger(r.Context()).Info(
			"access log",
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", duration.Milliseconds(),
			"status", rw.status,
		)
	})
}
