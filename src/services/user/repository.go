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
	err := r.db.Where("id IN ?", ids).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}