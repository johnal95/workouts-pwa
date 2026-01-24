package exercise

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
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

type exerciseRow struct {
	ID          string
	CreatedAt   time.Time
	Type        string
	Name        string
	Description sql.NullString
}

func (r *PostgresRepository) FindByID(ctx context.Context, exerciseID string) (*Exercise, error) {
	rows, err := r.db.Query(`
		SELECT e.id, e.created_at, e.type, et.name, et.description
		FROM exercises e
		LEFT JOIN exercise_translations et
		ON et.exercise_id = e.id AND et.locale = 'en-US'
		WHERE id = $1
	`, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercise *Exercise

	for rows.Next() {
		var er exerciseRow

		if err := rows.Scan(
			&er.ID,
			&er.CreatedAt,
			&er.Type,
			&er.Name,
			&er.Description,
		); err != nil {

			return nil, err
		}

		exercise = &Exercise{
			ID:        er.ID,
			CreatedAt: er.CreatedAt,
			Type:      ExerciseType(er.Type),
			Name:      er.Name,
		}

		if er.Description.Valid {
			exercise.Description = &er.Description.String
		}

	}

	if exercise == nil {
		return nil, ErrExerciseNotFound
	}

	return exercise, nil
}
