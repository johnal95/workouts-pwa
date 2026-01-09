package workout

import (
	"context"
	"errors"

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

func (s *Service) GetWorkout(ctx context.Context, userID, workoutID string) (*Workout, error) {
	logger := logging.Logger(ctx)

	w, err := s.repo.FindByID(ctx, userID, workoutID)

	if err != nil {
		if errors.Is(err, ErrWorkoutNotFound) {
			logger.Warn("workout not found",
				"user_id", userID,
				"id", workoutID)
		} else {
			logger.Error("failed to get workout by ID",
				"error", err)
		}
		return nil, err
	}

	return w, nil
}

func (s *Service) GetWorkouts(ctx context.Context, userID string) ([]*Workout, error) {
	return s.repo.FindAll(ctx, userID)
}

type CreateWorkoutInput struct {
	Name string
}

func (s *Service) CreateWorkout(ctx context.Context, userID string, input *CreateWorkoutInput) (*Workout, error) {
	logger := logging.Logger(ctx)

	workout, err := s.repo.Create(ctx, userID, &Workout{Name: input.Name})
	if err != nil {
		if errors.Is(err, ErrWorkoutNameAlreadyExists) {
			logger.Warn("duplicate workout name",
				"user_id", userID,
				"name", input.Name)
		} else {
			logger.Error("failed to create workout", "error", err)
		}
		return nil, err
	}

	return workout, nil
}

type CreateExerciseInput struct {
	Name            string
	DefaultSetCount uint
	MinReps         uint
	MaxReps         uint
}

func (s *Service) CreateExercise(ctx context.Context, userID, workoutID string, input *CreateExerciseInput) (*Exercise, error) {
	logger := logging.Logger(ctx)

	exercise, err := s.repo.CreateExercise(ctx, userID, workoutID, &Exercise{
		Name:            input.Name,
		DefaultSetCount: input.DefaultSetCount,
		MinReps:         input.MinReps,
		MaxReps:         input.MaxReps,
	})
	if err != nil {
		if errors.Is(err, ErrExerciseNameAlreadyExists) {
			logger.Warn("duplicate exercise name",
				"user_id", userID,
				"workout_id", workoutID,
				"name", input.Name)
		} else if errors.Is(err, ErrWorkoutNotFound) {
			logger.Warn("workout not found",
				"user_id", userID,
				"id", workoutID)
		} else {
			logger.Error("failed to create workout",
				"error", err)
		}
		return nil, err
	}

	return exercise, nil
}

func (s *Service) DeleteWorkout(ctx context.Context, userID, workoutID string) error {
	return s.repo.Delete(ctx, userID, workoutID)
}
