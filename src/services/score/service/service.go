package score_service

import (
	"context"
	"ggclass_go/src/models"
)

type IRepository interface {
	GetMarkForFirstTimeToDo(ctx context.Context, exerciseIds []int) ([]models.Assigment, error)
	GetMarkForNewest(ctx context.Context, exerciseIds []int) ([]models.Assigment, error)
	GetMarkForHighest(ctx context.Context, exerciseIds []int) ([]models.Assigment, error)
}

type service struct {
	repository      IRepository
	exerciseService IExerciseService
	userService     IUserService
}

type IExerciseService interface {
	GetByClassId(ctx context.Context, classId int) ([]models.Exercise, error)
}

type IUserService interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]models.User, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetExerciseService(service IExerciseService) {
	s.exerciseService = service
}

func (s *service) SetUserService(service IUserService) {
	s.userService = service
}
