package lesson

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

func (r *repository) Create(ctx context.Context, lesson *models.Lesson) error {
	return r.db.Create(lesson).Error
}

func (r *repository) Update(ctx context.Context, lesson *models.Lesson) error {
	return r.db.Save(lesson).Error
}

func (r *repository) FindById(ctx context.Context, id int) (*models.Lesson, error) {
	var result models.Lesson
	err := r.db.Where("id = ?", id).First(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repository) FindByFolderId(ctx context.Context, folderId int) ([]models.Lesson, error) {
	var result []models.Lesson
	err := r.db.Where("folder_id = ?", folderId).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}