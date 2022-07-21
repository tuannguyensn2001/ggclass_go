package profile

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type IRepository interface {
	FindByUserId(ctx context.Context, userId int) (*models.Profile, error)
	Create(ctx context.Context, profile *models.Profile) error
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) GetByUserId(ctx context.Context, userId int) (*models.Profile, error) {
	profile, err := s.repository.FindByUserId(ctx, userId)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *service) Create(ctx context.Context, input CreateProfileInput) (*models.Profile, error) {

	check, err := s.GetByUserId(ctx, input.UserId)

	if err != nil {
		return nil, err
	}

	if check != nil {
		return nil, app.ConflictHttpError("conflict", errors.New("conflict"))
	}

	profile := models.Profile{
		Avatar: input.Avatar,
		UserId: input.UserId,
	}

	err = s.repository.Create(ctx, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, err
}
