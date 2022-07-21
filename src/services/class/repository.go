package class

import (
	"context"
	"errors"
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

func (r *repository) Create(ctx context.Context, class *models.Class) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(class).Error; err != nil {
			return err
		}

		userClass := models.UserClass{
			UserId:  class.CreatedBy,
			ClassId: class.Id,
			Role:    enums.ADMIN,
			Status:  enums.ACTIVE,
		}

		if err := tx.Create(&userClass).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *repository) FindByNameAndCreateBy(ctx context.Context, name string, userId int) (*models.Class, error) {
	var result models.Class

	err := r.db.Where("name = ?", name).Where("created_by = ?", userId).First(&result).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &result, nil

}

func (r *repository) AddMember(ctx context.Context, userClass *models.UserClass) error {
	return r.db.Create(userClass).Error
}

func (r *repository) FindByUserAndClass(ctx context.Context, userId int, classId int) (*models.UserClass, error) {
	var result models.UserClass

	err := r.db.Where("user_id = ?", userId).Where("class_id = ?", classId).First(&result).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &result, nil

}

func (r *repository) SetMemberInactive(ctx context.Context, userId int, classId int) error {
	return r.db.Model(models.UserClass{}).Where("user_id = ?", userId).Where("class_id = ?", classId).Update("status", enums.INACTIVE).Error
}

func (r *repository) GetUsersByClass(ctx context.Context, classId int) ([]models.UserClass, error) {
	var ids []models.UserClass

	err := r.db.Model(models.UserClass{}).Where("class_id = ?", classId).Find(&ids).Error

	if err != nil {
		return nil, err
	}

	return ids, nil

}

func (r *repository) GetActiveUsersByClass(ctx context.Context, classId int) ([]models.UserClass, error) {
	var ids []models.UserClass

	err := r.db.Model(models.UserClass{}).Where("class_id = ?", classId).Where("status != ?", enums.INACTIVE).Find(&ids).Error

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *repository) SetMemberPending(ctx context.Context, userId int, classId int) error {
	return r.db.Model(models.UserClass{}).Where("user_id = ?", userId).Where("class_id = ?", classId).Update("status", enums.PENDING).Error
}

func (r *repository) SetMemberActive(ctx context.Context, userId int, classId int) error {
	return r.db.Model(models.UserClass{}).Where("user_id = ?", userId).Where("class_id = ?", classId).Update("status", enums.ACTIVE).Error
}

func (r *repository) GetClassActiveByUser(ctx context.Context, userId int) ([]models.UserClass, error) {
	var result []models.UserClass

	err := r.db.Where("user_id = ?", userId).Where("status != ?", enums.INACTIVE).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) GetClassByIds(ctx context.Context, ids []int) ([]models.Class, error) {
	var result []models.Class
	err := r.db.Where("id IN ?", ids).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindById(ctx context.Context, id int) (*models.Class, error) {
	var result models.Class
	if err := r.db.Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}