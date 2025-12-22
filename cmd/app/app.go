package app

import (
	"database/sql"

	"github.com/johnal95/workouts-pwa/cmd/api"
	"github.com/johnal95/workouts-pwa/cmd/middleware"
	"github.com/johnal95/workouts-pwa/cmd/store"
	"github.com/johnal95/workouts-pwa/migrations"
)

type Application struct {
	AuthMiddleware *middleware.AuthMiddleware
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

type ApplicationOptions struct {
	DatabaseURL string
}

func NewApplication(options *ApplicationOptions) (*Application, error) {

	pgDB, err := store.Open(options.DatabaseURL)
	if err != nil {
		return nil, err
	}

	err = store.Migrate(pgDB, migrations.FS, ".")
	if err != nil {
		return nil, err
	}

	// TEMPORARY USER
	pgDB.Query(`DELETE FROM users`)
	pgDB.Query(`INSERT INTO users (id, email, auth_id, auth_provider) VALUES ($1, $2, $3, $4)`,
		"019b4388-50ee-7f94-9caf-a8ceb54ef056", "john.doe@gmail.com", "random_auth_id", "GOOGLE")

	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	workoutHandler := api.NewWorkoutHandler(workoutStore)

	authMiddleware := middleware.NewAuthMiddleware()

	return &Application{
		AuthMiddleware: authMiddleware,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}, nil
}
