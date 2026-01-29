package http

import "github.com/johnal95/workouts-pwa/internal/exercise"

func ToExerciseResponse(e *exercise.Exercise) *ExerciseResponse {
	return &ExerciseResponse{
		ID:          e.ID,
		Type:        e.Type,
		Name:        e.Name,
		Description: e.Description,
	}
}

func ToExerciseListResponse(e []*exercise.Exercise) []*ExerciseResponse {
	list := []*ExerciseResponse{}
	for _, e := range e {
		list = append(list, ToExerciseResponse(e))
	}
	return list
}
