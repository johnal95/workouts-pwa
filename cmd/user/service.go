package user

import (
	"errors"
	"log/slog"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUser(userID string) (*User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			slog.Warn("user not found", "user_id", userID)
		} else {
			slog.Error("failed to get user by ID", "user_id", userID, "error", err)
		}
		return nil, err
	}
	return user, nil
}

func (s *Service) CreateUser(u *User) (*User, error) {
	user, err := s.repo.Create(u)
	if err != nil {
		if errors.Is(err, ErrUserEmailAlreadyExists) {
			slog.Warn("user with email already exists", "email", u.Email)
		} else {
			slog.Error("failed to create user", "error", err)
		}
		return nil, err
	}
	return user, nil
}
