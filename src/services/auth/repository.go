package auth

import (
	"context"
	"errors"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *models.User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}

func (r *repository) FindById(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
