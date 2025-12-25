package store

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type UserStore interface {
	FindById(userId string) (*User, error)
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

func (s *PostgresUserStore) FindById(userId string) (*User, error) {
	var user User
	err := s.db.QueryRow(`
		SELECT
			id,
			created_at,
			email,
			auth_id,
			auth_provider
		FROM users
		WHERE id = $1
	`,
		userId,
	).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.Email,
		&user.AuthId,
		&user.AuthProvider,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *PostgresUserStore) Create(u *User) (*User, error) {
	err := s.db.QueryRow(`
		INSERT INTO users
			(email, auth_provider, auth_id)
		VALUES
			($1, $2 ,$3)
		RETURNING
			id, created_at, email, auth_id, auth_provider
	`,
		u.Email,
		u.AuthProvider,
		u.AuthId,
	).Scan(
		&u.Id,
		&u.CreatedAt,
		&u.Email,
		&u.AuthId,
		&u.AuthProvider,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, errors.Join(ErrUniqueConstraint, err)
		}
		return nil, err
	}
	return u, nil
}
