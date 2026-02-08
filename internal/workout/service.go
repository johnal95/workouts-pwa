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
	w, err := s.repo.FindByID(ctx, userID, workoutID)

	if err != nil {
		if errors.Is(err, ErrWorkoutNotFound) {
			logging.Logger(ctx).Warn(
				"workout not found",
				"id", workoutID,
			)
		} else {
			logging.Logger(ctx).Error(
				"failed to get workout by ID",
				"error", err,
			)
		}
		return nil, err
	}

	return w, nil
}

func (s *Service) GetWorkouts(ctx context.Context, userID string) ([]*Workout, error) {
	workouts, err := s.repo.FindAll(ctx, userID)
	if err != nil {
		logging.Logger(ctx).Error(
			"failed to retrieve workouts",
			"error", err,
		)
		return nil, err
	}

	return workouts, nil
}

type CreateWorkoutInput struct {
	Name string
}

func (s *Service) CreateWorkout(ctx context.Context, userID string, input *CreateWorkoutInput) (*Workout, error) {
	workout, err := s.repo.Create(ctx, userID, &Workout{Name: input.Name})
	if err != nil {
		if errors.Is(err, ErrWorkoutNameAlreadyExists) {
			logging.Logger(ctx).Warn(
				"duplicate workout name",
				"error", err,
			)
		} else if errors.Is(err, ErrWorkoutLimitReached) {
			logging.Logger(ctx).Warn(
				"workout limit reached",
				"error", err,
			)
		} else {
			logging.Logger(ctx).Error(
				"failed to create workout",
				"error", err,
			)
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
	cwe, err := s.repo.CreateWorkoutExercise(ctx, userID, workoutID, input.ExerciseID, input.Notes)
	if err != nil {
		logging.Logger(ctx).Error(
			"failed to create workout exercise",
			"workout_id", workoutID,
			"exercise_id", input.ExerciseID,
			"error", err,
		)
		return nil, err
	}

	exercise, err := s.exerciseService.GetExercise(ctx, cwe.ExerciseID)
	if err != nil {
		logging.Logger(ctx).Error(
			"failed to retrieve exercise data after creating workout exercise",
			"id", cwe.ID,
			"workout_id", workoutID,
			"exercise_id", input.ExerciseID,
			"error", err,
		)
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

func (s *Service) UpdateWorkoutOrder(ctx context.Context, userID, workoutID string, workoutExerciseIDOrder []string) ([]string, error) {
	updatedWorkoutExerciseIDOrder, err := s.repo.UpdateWorkoutExerciseOrder(ctx, userID, workoutID, workoutExerciseIDOrder)
	if err != nil {
		if errors.Is(err, ErrInvalidWorkoutExerciseIDs) {
			logging.Logger(ctx).Warn(
				"invalid workout exercise IDs",
				"id", workoutID,
				"workout_exercise_ids", workoutExerciseIDOrder,
			)
		} else {
			logging.Logger(ctx).Error(
				"failed to update workout workout exercise order",
				"workout_id", workoutID,
				"error", err,
			)
		}
		return nil, err
	}

	return updatedWorkoutExerciseIDOrder, nil
}

func (s *Service) DeleteWorkout(ctx context.Context, userID, workoutID string) error {
	err := s.repo.Delete(ctx, userID, workoutID)
	if err != nil {
		if errors.Is(err, ErrWorkoutNotFound) {
			logging.Logger(ctx).Warn(
				"workout not found",
				"id", workoutID,
			)
		} else {
			logging.Logger(ctx).Error(
				"failed to delete workout",
				"error", err,
			)
		}
		return err
	}
	return nil
}
