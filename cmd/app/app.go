package app

import (
	"database/sql"
	"fmt"

	"github.com/johnal95/workouts-pwa/cmd/middleware"
	"github.com/johnal95/workouts-pwa/cmd/sqlx"
	"github.com/johnal95/workouts-pwa/cmd/user"
	"github.com/johnal95/workouts-pwa/cmd/workout"
	workouthttp "github.com/johnal95/workouts-pwa/cmd/workout/transport/http"
	"github.com/johnal95/workouts-pwa/migrations"
)

type Application struct {
	AuthMiddleware      *middleware.AuthMiddleware
	RequestIDMiddleware *middleware.RequestIDMiddleware
	WorkoutHandler      *workouthttp.Handler
	DB                  *sql.DB
}

type ApplicationOptions struct {
	DatabaseURL string
}

func NewApplication(options *ApplicationOptions) (*Application, error) {

	pgDB, err := sqlx.Open(options.DatabaseURL)
	if err != nil {
		return nil, err
	}

	err = sqlx.Migrate(pgDB, migrations.FS, ".")
	if err != nil {
		return nil, err
	}

	// Repositories
	userRepo := user.NewPostgresRepository(pgDB)
	workoutRepo := workout.NewPostgresRepository(pgDB)

	// Services
	userService := user.NewService(userRepo)
	workoutService := workout.NewService(workoutRepo)

	// Middlewares
	authMiddleware := middleware.NewAuthMiddleware()
	requestIDMiddleware := middleware.NewRequestIDMiddleware()

	// Handlers
	workoutHandler := workouthttp.NewHandler(workoutService)

	// TEMPORARY TEST USER SNIPPET
	pgDB.Query(`DELETE FROM users`)
	pgDB.Query(`INSERT INTO users (id, email, auth_id, auth_provider) VALUES ($1, $2, $3, $4)`,
		"019b4388-50ee-7f94-9caf-a8ceb54ef056", "john.doe@gmail.com", "random_auth_id", "GOOGLE")
	usr, _ := userService.GetUser("019b4388-50ee-7f94-9caf-a8ceb54ef056")
	fmt.Printf("TEST USER:\n%+v\n", *usr)

	return &Application{
		AuthMiddleware:      authMiddleware,
		RequestIDMiddleware: requestIDMiddleware,
		WorkoutHandler:      workoutHandler,
		DB:                  pgDB,
	}, nil
}
