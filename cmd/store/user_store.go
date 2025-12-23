package store

import "database/sql"

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
	return nil, nil
}
