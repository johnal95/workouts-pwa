package exercise

import (
	"context"

	"github.com/johnal95/workouts-pwa/internal/logging"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetByID(ctx context.Context, exerciseID string) (*Exercise, error) {
	e, err := s.repo.FindByID(ctx, exerciseID)

	if err != nil {
		logging.Logger(ctx).Error(
			"failed to retrieve exercise",
			"error", err,
		)
		return nil, err
	}

	return e, nil
}
