package members

import (
	"context"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) AddMember(ctx context.Context, userClass *models.UserClass) error {
	return r.db.Create(userClass).Error
}

func (r *repository) GetByUserAndClass(ctx context.Context, userId int, classId int) (*models.UserClass, error) {
	var result models.UserClass
	err := r.db.Where("user_id = ?", userId).Where("class_id = ?", classId).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repository) Update(ctx context.Context, userClass *models.UserClass) error {
	return r.db.Save(userClass).Error
}

func (r *repository) GetStudentsPendingByClass(ctx context.Context, classId int) ([]models.UserClass, error) {
	var result []models.UserClass
	err := r.db.Where("class_id = ?", classId).Where("status = ?", enums.PENDING).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateActiveByClass(ctx context.Context, classId int) error {
	return r.db.Model(&models.UserClass{}).Where("class_id = ?", classId).Where("role = ?", enums.STUDENT).Where("status = ?", enums.PENDING).Update("status", enums.ACTIVE).Error
}
