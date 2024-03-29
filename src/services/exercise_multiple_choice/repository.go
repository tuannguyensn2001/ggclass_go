package exercise_multiple_choice

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

func (r *repository) SetDB(db *gorm.DB) {
	r.db = db
}

func (r *repository) CreateMultipleChoice(ctx context.Context, multipleChoice *models.ExerciseMultipleChoice) error {
	return r.db.Create(multipleChoice).Error
}

func (r *repository) CreateMultipleChoiceAnswer(ctx context.Context, answers []models.ExerciseMultipleChoiceAnswer) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range answers {
			err := tx.Create(&item).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *repository) GetTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *repository) FindById(ctx context.Context, id int) (*models.ExerciseMultipleChoice, error) {
	var result models.ExerciseMultipleChoice
	if err := r.db.Where("id  = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindAnswersByMultipleChoiceId(ctx context.Context, id int) ([]models.ExerciseMultipleChoiceAnswer, error) {
	var result []models.ExerciseMultipleChoiceAnswer
	if err := r.db.Where("exercise_multiple_choice_id = ?", id).Find(&result).Error; err != nil {
		return nil, err

	}
	return result, nil
}
