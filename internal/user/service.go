package user

import (
	"context"
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

func (s *Service) GetUser(ctx context.Context, userID string) (*User, error) {
	user, err := s.repo.FindByID(ctx, userID)
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

type CreateUserInput struct {
	Email        string
	AuthProvider AuthProvider
	AuthID       string
}

func (s *Service) CreateUser(ctx context.Context, input *CreateUserInput) (*User, error) {
	user, err := s.repo.Create(ctx, &User{
		Email:        input.Email,
		AuthProvider: input.AuthProvider,
		AuthID:       input.AuthID,
	})
	if err != nil {
		if errors.Is(err, ErrUserEmailAlreadyExists) {
			slog.Warn("user with email already exists", "email", input.Email)
		} else {
			slog.Error("failed to create user", "error", err)
		}
		return nil, err
	}
	return user, nil
}
