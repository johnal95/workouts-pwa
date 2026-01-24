package workout

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/johnal95/workouts-pwa/internal/exercise"
)

type Repository interface {
	FindAll(ctx context.Context, userID string) ([]*Workout, error)
	FindByID(ctx context.Context, userID, workoutID string) (*Workout, error)
	Create(ctx context.Context, userID string, w *Workout) (*Workout, error)
	CreateWorkoutExercise(ctx context.Context, userID string, workoutID string, exerciseID string, notes *string) (*CreatedWorkoutExercise, error)
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

const baseWorkoutSelectQuery = `
SELECT
	w.id,
	w.created_at,
	w.name,
	e.id,
	e.created_at,
	e.type,
	et.name,
	et.description,
	we.id,
	we.position,
	we.notes
FROM workouts w
LEFT JOIN workout_exercises we
ON we.workout_id = w.id
LEFT JOIN exercises e
ON we.exercise_id = e.id
LEFT JOIN exercise_translations et
ON et.exercise_id = e.id AND et.locale = 'en-US'
WHERE w.user_id = $1
`

type workoutRow struct {
	WorkoutID        string
	WorkoutCreatedAt time.Time
	WorkoutName      string

	ExerciseID          sql.NullString
	ExerciseCreatedAt   sql.NullTime
	ExerciseType        sql.NullString
	ExerciseName        sql.NullString
	ExerciseDescription sql.NullString

	WorkoutExerciseID       sql.NullString
	WorkoutExercisePosition sql.NullInt16
	WorkoutExerciseNotes    sql.NullString
}

func scanWorkoutRow(rows *sql.Rows) (*workoutRow, error) {
	var r workoutRow
	if err := rows.Scan(
		&r.WorkoutID,
		&r.WorkoutCreatedAt,
		&r.WorkoutName,
		&r.ExerciseID,
		&r.ExerciseCreatedAt,
		&r.ExerciseType,
		&r.ExerciseName,
		&r.ExerciseDescription,
		&r.WorkoutExerciseID,
		&r.WorkoutExercisePosition,
		&r.WorkoutExerciseNotes,
	); err != nil {
		return nil, err
	}
	return &r, nil
}

func workoutExerciseFromRow(workoutID string, r *workoutRow) *WorkoutExercise {
	if !r.ExerciseID.Valid {
		return nil
	}

	exercise := exercise.Exercise{
		ID:        r.ExerciseID.String,
		CreatedAt: r.ExerciseCreatedAt.Time,
		Type:      exercise.ExerciseType(r.ExerciseType.String),
		Name:      r.ExerciseName.String,
	}

	if r.ExerciseDescription.Valid {
		exercise.Description = &r.ExerciseDescription.String
	}

	workoutExercise := &WorkoutExercise{
		ID:        r.WorkoutExerciseID.String,
		WorkoutID: workoutID,
		Exercise:  exercise,
		Position:  int(r.WorkoutExercisePosition.Int16),
	}

	if r.WorkoutExerciseNotes.Valid {
		workoutExercise.Notes = &r.WorkoutExerciseNotes.String
	}
	return workoutExercise
}

func (r *PostgresRepository) FindByID(ctx context.Context, userID, workoutID string) (*Workout, error) {
	rows, err := r.db.Query(baseWorkoutSelectQuery+" AND w.id = $2", userID, workoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workout *Workout

	for rows.Next() {
		wRow, err := scanWorkoutRow(rows)
		if err != nil {
			return nil, err
		}

		if workout == nil {
			workout = &Workout{
				ID:        wRow.WorkoutID,
				CreatedAt: wRow.WorkoutCreatedAt,
				Name:      wRow.WorkoutName,
				Exercises: []*WorkoutExercise{},
			}
		}

		if we := workoutExerciseFromRow(wRow.WorkoutID, wRow); we != nil {
			workout.Exercises = append(workout.Exercises, we)
		}
	}

	if workout == nil {
		return nil, ErrWorkoutNotFound
	}

	return workout, nil
}

func (r *PostgresRepository) FindAll(ctx context.Context, userID string) ([]*Workout, error) {
	rows, err := r.db.Query(baseWorkoutSelectQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workoutsMap := map[string]*Workout{}

	for rows.Next() {
		wRow, err := scanWorkoutRow(rows)
		if err != nil {
			return nil, err
		}

		workout, exists := workoutsMap[wRow.WorkoutID]
		if !exists {
			workout = &Workout{
				ID:        wRow.WorkoutID,
				CreatedAt: wRow.WorkoutCreatedAt,
				Name:      wRow.WorkoutName,
				Exercises: []*WorkoutExercise{},
			}
			workoutsMap[wRow.WorkoutID] = workout
		}
		if we := workoutExerciseFromRow(wRow.WorkoutID, wRow); we != nil {
			workout.Exercises = append(workout.Exercises, we)
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
		Exercises: []*WorkoutExercise{},
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

type CreatedWorkoutExercise struct {
	ID         string
	WorkoutID  string
	ExerciseID string
	Position   int
	Notes      *string
}

func (r *PostgresRepository) CreateWorkoutExercise(
	ctx context.Context,
	userID string,
	workoutID string,
	exerciseID string,
	notes *string,
) (*CreatedWorkoutExercise, error) {
	var insertedNotes sql.NullString

	we := &CreatedWorkoutExercise{}

	if err := r.db.QueryRow(`
		INSERT INTO workout_exercises (workout_id, exercise_id, position, notes)
		SELECT 
			w.id,
			$1,
			COALESCE((SELECT MAX(position) FROM workout_exercises WHERE workout_id = w.id), 0) + 1,
			$2
		FROM workouts w
		WHERE w.id = $3 AND w.user_id = $4
		RETURNING id, workout_id, exercise_id, position, notes
	`, exerciseID, notes, workoutID, userID,
	).Scan(&we.ID, &we.WorkoutID, &we.ExerciseID, &we.Position, &insertedNotes); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrWorkoutNotFound
		}
		return nil, err
	}

	if insertedNotes.Valid {
		we.Notes = &insertedNotes.String
	}

	return we, nil
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
