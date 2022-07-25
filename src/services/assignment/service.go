package assignment

import (
	"context"
	"ggclass_go/src/models"
)

type IRepository interface {
	Create(ctx context.Context, assignment *models.Assigment) error
}

type service struct {
	repository    IRepository
	exerciseClone IExerciseClone
}

type IExerciseClone interface {
	GetLatestClone(ctx context.Context, exerciseId int) (*models.ExerciseClone, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetExerciseCloneService(exerciseCloneService IExerciseClone) {
	s.exerciseClone = exerciseCloneService
}

func (s *service) Start(ctx context.Context, input StartAssignmentInput) (*models.Assigment, error) {

	exerciseClone, err := s.exerciseClone.GetLatestClone(ctx, input.ExerciseId)
	if err != nil {
		return nil, err
	}

	assignment := models.Assigment{
		ExerciseId:      input.ExerciseId,
		UserId:          input.UserId,
		ExerciseCloneId: exerciseClone.Id,
	}
	err = s.repository.Create(ctx, &assignment)
	if err != nil {
		return nil, nil
	}

	return &assignment, nil
}
