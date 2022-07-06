package post

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

func (r *repository) Create(ctx context.Context, post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *repository) FindPostsByClassId(ctx context.Context, classId int) ([]models.Post, error) {
	var result []models.Post

	err := r.db.Where("class_id = ?", classId).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil

}