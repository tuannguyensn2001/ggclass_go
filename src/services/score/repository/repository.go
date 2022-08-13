package repository

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

func (r *repository) GetMarkForFirstTimeToDo(ctx context.Context, exerciseIds []int) ([]models.Assigment, error) {
	var result []models.Assigment

	err := r.db.Where("exercise_id IN ?", exerciseIds).Where("is_submit = ?", 1).Limit(1).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (r *repository) GetMarkForNewest(ctx context.Context, exerciseIds []int) ([]models.Assigment, error) {
	var result []models.Assigment

	err := r.db.Raw(`select * from assignments where exercise_id in ? order by id desc limit 1`, exerciseIds).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) GetMarkForHighest(ctx context.Context, exerciseIds []int) ([]models.Assigment, error) {
	var result []models.Assigment

	err := r.db.Raw(`select max(mark) as mark,user_id,exercise_id from assignments where exercise_id in ? group by user_id,exercise_id`, exerciseIds).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
