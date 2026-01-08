package logging

import (
	"log/slog"
	"os"

	"github.com/johnal95/workouts-pwa/internal/config"
)

func NewLogger(env config.Environment) *slog.Logger {
	var handler slog.Handler

	if env == config.EnvDevelopment {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	return slog.New(handler)
}
