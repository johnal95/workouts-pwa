package exercise

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	FindAll(ctx context.Context) ([]*Exercise, error)
	FindByID(ctx context.Context, exerciseID string) (*Exercise, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

const baseExerciseSelectQuery = `
SELECT e.id, e.created_at, e.type, et.name, et.description
FROM exercises e
LEFT JOIN exercise_translations et
ON et.exercise_id = e.id AND et.locale = 'en-US'
`

type exerciseRow struct {
	ID          string
	CreatedAt   time.Time
	Type        string
	Name        string
	Description sql.NullString
}

func scanExerciseRow(rows *sql.Rows) (*exerciseRow, error) {
	var r exerciseRow
	if err := rows.Scan(
		&r.ID,
		&r.CreatedAt,
		&r.Type,
		&r.Name,
		&r.Description,
	); err != nil {
		return nil, err
	}
	return &r, nil
}

func exerciseFromRow(er *exerciseRow) *Exercise {
	e := &Exercise{
		ID:        er.ID,
		CreatedAt: er.CreatedAt,
		Type:      ExerciseType(er.Type),
		Name:      er.Name,
	}
	if er.Description.Valid {
		e.Description = &er.Description.String
	}
	return e
}

func (r *PostgresRepository) FindAll(ctx context.Context) ([]*Exercise, error) {
	rows, err := r.db.Query(baseExerciseSelectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	exercises := []*Exercise{}

	for rows.Next() {
		er, err := scanExerciseRow(rows)
		if err != nil {
			return nil, err
		}

		exercises = append(exercises, exerciseFromRow(er))
	}

	return exercises, nil
}

func (r *PostgresRepository) FindByID(ctx context.Context, exerciseID string) (*Exercise, error) {
	rows, err := r.db.Query(baseExerciseSelectQuery+` WHERE id = $1`, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercise *Exercise

	for rows.Next() {
		er, err := scanExerciseRow(rows)
		if err != nil {
			return nil, err
		}
		exercise = exerciseFromRow(er)
	}

	if exercise == nil {
		return nil, ErrExerciseNotFound
	}

	return exercise, nil
}
