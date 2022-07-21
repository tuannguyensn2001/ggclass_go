package folder

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

func (r *repository) Create(ctx context.Context, folder *models.Folder) error {
	return r.db.Create(folder).Error
}

func (r *repository) FindByNameAndClassId(ctx context.Context, name string, classId int) (*models.Folder, error) {
	var result models.Folder

	err := r.db.Where("name = ?", name).Where("class_id = ?", classId).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil

}

func (r *repository) FindByQuery(ctx context.Context, query GetFoldersQuery) ([]models.Folder, error) {
	var result []models.Folder

	builder := r.db

	if query.classId > 0 {
		builder = builder.Where("class_id = ?", query.classId)
	}

	err := builder.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil

}
