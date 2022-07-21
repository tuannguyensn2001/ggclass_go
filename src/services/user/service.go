package user

import (
	"context"
	"ggclass_go/src/models"
)

type IRepository interface {
	GetByIds(ctx context.Context, ids []int) ([]models.User, error)
	FindById(ctx context.Context, id int) (*models.User, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) GetUsersByIds(ctx context.Context, ids []int) ([]models.User, error) {
	return s.repository.GetByIds(ctx, ids)
}

func (s *service) GetById(ctx context.Context, userId int) (*models.User, error) {
	return s.repository.FindById(ctx, userId)
}
