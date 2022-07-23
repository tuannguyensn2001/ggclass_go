package user

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type IRepository interface {
	GetByIds(ctx context.Context, ids []int) ([]models.User, error)
	FindById(ctx context.Context, id int) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
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

func (s *service) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.repository.FindByEmail(ctx, email)

	if (err != nil && errors.Is(err, gorm.ErrRecordNotFound)) || user == nil {
		return nil, app.NotFoundHttpError("not found user", err)
	}
	if err != nil {
		return nil, err
	}

	return user, err
}
