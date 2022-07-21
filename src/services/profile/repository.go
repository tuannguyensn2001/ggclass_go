package profile

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

func (r *repository) FindByUserId(ctx context.Context, userId int) (*models.Profile, error) {
	var result models.Profile
	if err := r.db.Where("user_id = ?", userId).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repository) Create(ctx context.Context, profile *models.Profile) error {
	return r.db.Create(profile).Error
}
