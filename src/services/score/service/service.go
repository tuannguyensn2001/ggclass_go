package score_service

import (
	"context"
	"ggclass_go/src/models"
)

type IRepository interface {
}

type service struct {
	repository      IRepository
	exerciseService IExerciseService
}

type IExerciseService interface {
	GetByClassId(ctx context.Context, classId int) ([]models.Exercise, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetExerciseService(service IExerciseService) {
	s.exerciseService = service
}
