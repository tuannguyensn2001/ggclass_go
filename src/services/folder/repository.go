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

func (r *repository) GetNumberFolderInClass(ctx context.Context, classId int) (int, error) {
	result := 0
	err := r.db.Raw(`select count(id) from folders where class_id = ?`, classId).Scan(&result).Error

	return result, err

}

func (r *repository) GetRootFolderInClass(ctx context.Context, classId int) (*models.Folder, error) {
	var result models.Folder

	err := r.db.Where("class_id = ?", classId).Where("is_root = ?", 1).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil

}

func (r *repository) FindById(ctx context.Context, id int) (*models.Folder, error) {
	var result models.Folder

	err := r.db.Where("id = ?", id).First(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil

}
