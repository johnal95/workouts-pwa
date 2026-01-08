package user

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type Repository interface {
	FindByID(userID string) (*User, error)
	Create(u *User) (*User, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) FindByID(userID string) (*User, error) {
	var user User

	if err := r.db.QueryRow(`
		SELECT id, created_at, email, auth_id, auth_provider
		FROM users
		WHERE id = $1
	`, userID,
	).Scan(&user.ID, &user.CreatedAt, &user.Email, &user.AuthID, &user.AuthProvider); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (r *PostgresRepository) Create(u *User) (*User, error) {
	var user User
	if err := r.db.QueryRow(`
		INSERT INTO users (email, auth_provider, auth_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, email, auth_provider, auth_id
	`, u.Email, u.AuthProvider, u.AuthID,
	).Scan(&user.ID, &user.CreatedAt, &user.Email, &user.AuthProvider, &user.AuthID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
			return nil, errors.Join(ErrUserEmailAlreadyExists, err)
		}
		return nil, err
	}

	return &user, nil
}
