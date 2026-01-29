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

func (s *Service) GetExercise(ctx context.Context, exerciseID string) (*Exercise, error) {
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

func (s *Service) GetExercises(ctx context.Context) ([]*Exercise, error) {
	exercises, err := s.repo.FindAll(ctx)
	if err != nil {
		logging.Logger(ctx).Error(
			"failed to retrieve exercises",
			"error", err,
		)
		return nil, err
	}
	return exercises, nil
}
