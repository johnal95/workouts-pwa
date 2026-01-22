package app

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/johnal95/workouts-pwa/internal/httpx"
	"github.com/johnal95/workouts-pwa/internal/middleware"
	"github.com/johnal95/workouts-pwa/internal/sqlx"
	"github.com/johnal95/workouts-pwa/internal/user"
	"github.com/johnal95/workouts-pwa/internal/workout"
	workouthttp "github.com/johnal95/workouts-pwa/internal/workout/transport/http"
)

type Application struct {
	AuthMiddleware      *middleware.AuthMiddleware
	RequestIDMiddleware *middleware.RequestIDMiddleware
	LoggingMiddleware   *middleware.LoggingMiddleware
	WorkoutHandler      *workouthttp.Handler
	DB                  *sql.DB
}

type ApplicationOptions struct {
	DatabaseURL string
	Logger      *slog.Logger
}

func NewApplication(options *ApplicationOptions) (*Application, error) {

	pgDB, err := sqlx.Open(options.DatabaseURL)
	if err != nil {
		return nil, err
	}

	// Validators
	validate := validator.New(validator.WithRequiredStructEnabled())

	// Parsers
	parser := httpx.NewParser(validate)

	// Repositories
	userRepo := user.NewPostgresRepository(pgDB)
	workoutRepo := workout.NewPostgresRepository(pgDB)

	// Services
	userService := user.NewService(userRepo)
	workoutService := workout.NewService(workoutRepo)

	// Middlewares
	authMiddleware := middleware.NewAuthMiddleware()
	requestIDMiddleware := middleware.NewRequestIDMiddleware()
	loggingMiddleware := middleware.NewLoggingMiddleware(options.Logger)

	// Handlers
	workoutHandler := workouthttp.NewHandler(parser, workoutService)

	// TEMPORARY TEST USER SNIPPET
	pgDB.Query(`DELETE FROM users`)
	pgDB.Query(`INSERT INTO users (id, email, auth_id, auth_provider) VALUES ($1, $2, $3, $4)`,
		"019b4388-50ee-7f94-9caf-a8ceb54ef056", "john.doe@gmail.com", "random_auth_id", "GOOGLE")
	usr, _ := userService.GetUser(context.Background(), "019b4388-50ee-7f94-9caf-a8ceb54ef056")
	fmt.Printf("TEST USER:\n%+v\n", *usr)

	return &Application{
		AuthMiddleware:      authMiddleware,
		RequestIDMiddleware: requestIDMiddleware,
		LoggingMiddleware:   loggingMiddleware,
		WorkoutHandler:      workoutHandler,
		DB:                  pgDB,
	}, nil
}
