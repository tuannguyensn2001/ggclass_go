package user

import (
	"context"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetByIds(ctx context.Context, ids []int) ([]models.User, error) {
	var result []models.User
	err := r.db.Where("id IN ?", ids).Preload("Profile").Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindById(ctx context.Context, userId int) (*models.User, error) {
	var result models.User
	err := r.db.Preload("Profile").Where("id = ?", userId).First(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, err

}

func (r *repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var result models.User
	err := r.db.Preload("Profile").Where("email = ?", email).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
