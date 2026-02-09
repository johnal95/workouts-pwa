package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/johnal95/workouts-pwa/internal/auth"
	authhttp "github.com/johnal95/workouts-pwa/internal/auth/transport/http"
	"github.com/johnal95/workouts-pwa/internal/exercise"
	exercisehttp "github.com/johnal95/workouts-pwa/internal/exercise/transport/http"
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
	AuthHandler         *authhttp.Handler
	ExerciseHandler     *exercisehttp.Handler
	WorkoutHandler      *workouthttp.Handler
	DB                  *sql.DB
}

type ApplicationOptions struct {
	DatabaseURL      string
	SessionJWTSecret string
	Logger           *slog.Logger
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
	exerciseRepo := exercise.NewPostgresRepository(pgDB)
	workoutRepo := workout.NewPostgresRepository(pgDB)

	// Services
	authService := auth.NewService(options.SessionJWTSecret)
	userService := user.NewService(userRepo)
	exerciseService := exercise.NewService(exerciseRepo)
	workoutService := workout.NewService(workoutRepo, exerciseService)

	// Middlewares
	authMiddleware := middleware.NewAuthMiddleware(authService)
	requestIDMiddleware := middleware.NewRequestIDMiddleware()
	loggingMiddleware := middleware.NewLoggingMiddleware(options.Logger)

	// Handlers
	authHandler := authhttp.NewHandler(parser, authService)
	exerciseHandler := exercisehttp.NewHandler(parser, exerciseService)
	workoutHandler := workouthttp.NewHandler(parser, workoutService)

	// TEMPORARY TEST USER SNIPPET
	_, err = pgDB.Exec(`INSERT INTO users (id, email, auth_id, auth_provider) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`,
		"019b4388-50ee-7f94-9caf-a8ceb54ef056", "john.doe@gmail.com", "random_auth_id", "GOOGLE")
	if err != nil {
		log.Fatal(err)
	}
	usr, _ := userService.GetUser(context.Background(), "019b4388-50ee-7f94-9caf-a8ceb54ef056")
	fmt.Printf("TEST USER:\n%+v\n", *usr)

	return &Application{
		AuthMiddleware:      authMiddleware,
		RequestIDMiddleware: requestIDMiddleware,
		LoggingMiddleware:   loggingMiddleware,
		AuthHandler:         authHandler,
		ExerciseHandler:     exerciseHandler,
		WorkoutHandler:      workoutHandler,
		DB:                  pgDB,
	}, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
