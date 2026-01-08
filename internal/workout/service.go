package workout

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

func (s *Service) GetWorkout(userID, workoutID string) (*Workout, error) {
	return s.repo.FindByID(userID, workoutID)
}

func (s *Service) GetWorkouts(userID string) ([]*Workout, error) {
	return s.repo.FindAll(userID)
}

func (s *Service) CreateWorkout(w *Workout) (*Workout, error) {
	if len(w.Name) < 1 || len(w.Name) > 40 {
		return nil, ErrWorkoutNameInvalid
	}

	workout, err := s.repo.Create(w)
	if err != nil {
		if errors.Is(err, ErrWorkoutNameAlreadyExists) {
			slog.Warn("duplicate workout name", "user_id", w.UserID, "name", w.Name)
		} else {
			slog.Error("failed to create workout", "error", err)
		}
		return nil, err
	}

	return workout, nil
}

func (s *Service) DeleteWorkout(userID, workoutID string) error {
	return s.repo.Delete(userID, workoutID)
}
