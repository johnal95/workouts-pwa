package workout

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

type Repository interface {
	FindAll(ctx context.Context, userID string) ([]*Workout, error)
	FindByID(ctx context.Context, userID, workoutID string) (*Workout, error)
	Create(ctx context.Context, userID string, w *Workout) (*Workout, error)
	CreateExercise(ctx context.Context, userID, workoutID string, e *Exercise) (*Exercise, error)
	Delete(ctx context.Context, userID, workoutID string) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) FindByID(ctx context.Context, userID, workoutID string) (*Workout, error) {
	rows, err := r.db.Query(`
		SELECT w.id, w.created_at, w.name, e.id, e.name, e.default_set_count, e.min_reps, e.max_reps
		FROM workouts w
		LEFT JOIN exercises e
		ON e.workout_id = w.id
		WHERE w.user_id = $1
		AND w.id = $2
	`, userID, workoutID)
	if err != nil {
		return nil, err
	}

	workout := &Workout{
		Exercises: []*Exercise{},
	}

	for rows.Next() {
		var exerciseID sql.NullString
		var exerciseName sql.NullString
		var exerciseDefaultSetCount sql.NullInt16
		var exerciseMinReps sql.NullInt16
		var exerciseMaxReps sql.NullInt16

		if err := rows.Scan(&workout.ID, &workout.CreatedAt, &workout.Name, &exerciseID, &exerciseID, &exerciseDefaultSetCount, &exerciseMinReps, &exerciseMaxReps); err != nil {
			return nil, err
		}

		if exerciseID.Valid {
			workout.Exercises = append(workout.Exercises, &Exercise{
				ID:              exerciseID.String,
				Name:            exerciseName.String,
				DefaultSetCount: uint(exerciseDefaultSetCount.Int16),
				MinReps:         uint(exerciseMinReps.Int16),
				MaxReps:         uint(exerciseMaxReps.Int16),
			})
		}
	}

	if workout.ID == "" {
		return nil, ErrWorkoutNotFound
	}

	return workout, nil
}

func (r *PostgresRepository) FindAll(ctx context.Context, userID string) ([]*Workout, error) {
	rows, err := r.db.Query(`
		SELECT w.id, w.created_at, w.name, e.id, e.name, e.default_set_count, e.min_reps, e.max_reps
		FROM workouts w
		LEFT JOIN exercises e
		ON e.workout_id = w.id
		WHERE w.user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}

	workoutsMap := map[string]*Workout{}

	for rows.Next() {
		var workoutID string
		var createdAt time.Time
		var name string
		var exerciseID sql.NullString
		var exerciseName sql.NullString
		var exerciseDefaultSetCount sql.NullInt16
		var exerciseMinReps sql.NullInt16
		var exerciseMaxReps sql.NullInt16

		if err := rows.Scan(&workoutID, &createdAt, &name, &exerciseID, &exerciseName, &exerciseDefaultSetCount, &exerciseMinReps, &exerciseMaxReps); err != nil {
			return nil, err
		}

		workout, exists := workoutsMap[workoutID]
		if !exists {
			workout = &Workout{
				ID:        workoutID,
				CreatedAt: createdAt,
				Name:      name,
				Exercises: []*Exercise{},
			}
			workoutsMap[workoutID] = workout
		}
		if exerciseID.Valid {
			workout.Exercises = append(workout.Exercises, &Exercise{
				ID:              exerciseID.String,
				Name:            exerciseName.String,
				DefaultSetCount: uint(exerciseDefaultSetCount.Int16),
				MinReps:         uint(exerciseMinReps.Int16),
				MaxReps:         uint(exerciseMaxReps.Int16),
			})
		}
	}

	workouts := []*Workout{}
	for _, w := range workoutsMap {
		workouts = append(workouts, w)
	}

	return workouts, nil
}

func (r *PostgresRepository) Create(ctx context.Context, userID string, w *Workout) (*Workout, error) {
	newWorkout := &Workout{
		Exercises: []*Exercise{},
	}

	if err := r.db.QueryRow(`
		INSERT INTO workouts (user_id, name)
		VALUES ($1, $2)
		RETURNING id, created_at, name
	`, userID, w.Name,
	).Scan(&newWorkout.ID, &newWorkout.CreatedAt, &newWorkout.Name); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "workouts_user_id_name_key" {
			return nil, ErrWorkoutNameAlreadyExists
		}
		return nil, err

	}

	return newWorkout, nil
}

func (r *PostgresRepository) CreateExercise(ctx context.Context, userID, workoutID string, e *Exercise) (*Exercise, error) {
	newExercise := &Exercise{}

	if err := r.db.QueryRow(`
		INSERT INTO exercises (workout_id, name, default_set_count, min_reps, max_reps)
		SELECT w.id, $1, $2, $3, $4
		FROM workouts w
		WHERE w.id = $5
		AND w.user_id = $6
		RETURNING id, created_at, workout_id, name, default_set_count, min_reps, max_reps
	`, e.Name, e.DefaultSetCount, e.MinReps, e.MaxReps, workoutID, userID,
	).Scan(&newExercise.ID, &newExercise.CreatedAt, &newExercise.WorkoutID, &newExercise.Name, &newExercise.DefaultSetCount, &newExercise.MinReps, &newExercise.MaxReps); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrWorkoutNotFound
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "exercises_workout_id_name_key" {
			return nil, ErrExerciseNameAlreadyExists
		}
		return nil, err

	}

	return newExercise, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, userID, workoutID string) error {
	var w Workout
	err := r.db.QueryRow(`
		DELETE FROM workouts
		WHERE user_id = $1
		AND id = $2
		RETURNING id, created_at, name
	`, userID, workoutID,
	).Scan(&w.ID, &w.CreatedAt, &w.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Join(ErrWorkoutNotFound, err)
		}
		return err
	}
	return nil
}
