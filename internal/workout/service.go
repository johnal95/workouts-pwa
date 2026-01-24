package workout

import (
	"context"
	"errors"

	"github.com/johnal95/workouts-pwa/internal/exercise"
	"github.com/johnal95/workouts-pwa/internal/logging"
)

type Service struct {
	repo            Repository
	exerciseService *exercise.Service
}

func NewService(repo Repository, exerciseService *exercise.Service) *Service {
	return &Service{
		repo:            repo,
		exerciseService: exerciseService,
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
	logger := logging.Logger(ctx)

	workouts, err := s.repo.FindAll(ctx, userID)
	if err != nil {
		logger.Error("failed to retrieve workouts", "error", err)
		return nil, err
	}

	return workouts, nil
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

type CreateWorkoutExerciseInput struct {
	ExerciseID string
	Notes      *string
}

func (s *Service) CreateWorkoutExercise(
	ctx context.Context,
	userID string,
	workoutID string,
	input *CreateWorkoutExerciseInput,
) (*WorkoutExercise, error) {
	logger := logging.Logger(ctx)

	cwe, err := s.repo.CreateWorkoutExercise(ctx, userID, workoutID, input.ExerciseID, input.Notes)
	if err != nil {
		logger.Error("failed to create workout exercise",
			"user_id", userID,
			"workout_id", workoutID,
			"exercise_id", input.ExerciseID,
			"error", err)
		return nil, err
	}

	exercise, err := s.exerciseService.GetByID(ctx, cwe.ExerciseID)
	if err != nil {
		logger.Error("failed to retrieve exercise data after creating workout exercise",
			"id", cwe.ID,
			"user_id", userID,
			"workout_id", workoutID,
			"exercise_id", input.ExerciseID,
			"error", err)
		return nil, err
	}

	workoutExercise := &WorkoutExercise{
		ID:        cwe.ID,
		WorkoutID: cwe.WorkoutID,
		Exercise:  *exercise,
		Position:  cwe.Position,
		Notes:     cwe.Notes,
	}

	return workoutExercise, nil
}

func (s *Service) DeleteWorkout(ctx context.Context, userID, workoutID string) error {
	return s.repo.Delete(ctx, userID, workoutID)
}
