package store

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type WorkoutStore interface {
	FindAll(userId string) ([]Workout, error)
	FindById(userId string, workoutId string) (*Workout, error)
	Create(userId string, w *Workout) (*Workout, error)
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func NewPostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{
		db: db,
	}
}

func (s *PostgresWorkoutStore) FindById(userId string, workoutId string) (*Workout, error) {
	var workout Workout
	err := s.db.QueryRow(`
		SELECT
			id,
			created_at,
			name
		FROM workouts
		WHERE user_id = $1
		AND id = $2
	`,
		userId,
		workoutId,
	).Scan(
		&workout.Id,
		&workout.CreatedAt,
		&workout.Name,
	)
	if err != nil {
		return nil, err
	}

	return &workout, nil
}

func (s *PostgresWorkoutStore) FindAll(userId string) ([]Workout, error) {
	rows, err := s.db.Query(`
		SELECT
			id,
			created_at,
			name
		FROM workouts
		WHERE user_id = $1
	`, userId)
	if err != nil {
		return nil, err
	}

	workouts := []Workout{}

	for rows.Next() {
		var workout Workout
		err := rows.Scan(
			&workout.Id,
			&workout.CreatedAt,
			&workout.Name,
		)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}

	return workouts, nil
}

func (s *PostgresWorkoutStore) Create(userId string, w *Workout) (*Workout, error) {
	err := s.db.QueryRow(`
		INSERT INTO workouts
			(user_id, name)
		VALUES
			($1, $2)
		RETURNING
			id, created_at, name
	`,
		userId, w.Name,
	).Scan(
		&w.Id,
		&w.CreatedAt,
		&w.Name,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, errors.Join(ErrUniqueConstraint, err)
		}
		return nil, err
	}
	return w, nil
}
