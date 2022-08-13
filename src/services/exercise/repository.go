package exercise

import (
	"context"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db       *gorm.DB
	originDB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db, originDB: db}
}

func (r *repository) Create(ctx context.Context, exercise *models.Exercise) error {
	return r.db.Create(exercise).Error
}

func (r *repository) CreateMultipleChoice(ctx context.Context, multipleChoice *models.ExerciseMultipleChoice) error {
	return r.db.Create(multipleChoice).Error
}

func (r *repository) GetBeginTransaction() *gorm.DB {
	return r.db.Begin()
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

func (r *repository) SetDB(db *gorm.DB) {
	r.db = db
}

func (r *repository) GetDB() *gorm.DB {
	return r.db
}

func (r *repository) FindById(ctx context.Context, id int) (*models.Exercise, error) {
	var result models.Exercise
	if err := r.db.Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindMultipleChoiceById(ctx context.Context, id int) (*models.ExerciseMultipleChoice, error) {
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

func (r *repository) BeginTransaction() {
	tx := r.db.Begin()
	r.db = tx
}

func (r *repository) Commit() {
	r.db.Commit()
	r.db = r.originDB
}

func (r *repository) Rollback() {
	r.db.Rollback()
	r.db = r.originDB
}

func (r *repository) Save(ctx context.Context, exercise *models.Exercise) error {
	return r.db.Save(exercise).Error
}

func (r *repository) SaveMultipleChoice(ctx context.Context, multipleChoice *models.ExerciseMultipleChoice) error {
	return r.db.Save(multipleChoice).Error
}

func (r *repository) DeleteAnswersByMultipleChoiceId(ctx context.Context, id int) error {
	return r.db.Where("exercise_multiple_choice_id = ?", id).Delete(&models.ExerciseMultipleChoiceAnswer{}).Error
}

func (r *repository) FindByClassId(ctx context.Context, classId int) ([]models.Exercise, error) {
	var result []models.Exercise
	err := r.db.Where("class_id = ?", classId).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindExerciseCloneByExerciseIdAndVersion(ctx context.Context, exerciseId int, version int) (*models.ExerciseClone, error) {
	var result models.ExerciseClone
	err := r.db.Select("id").Where("exercise_id = ?", exerciseId).Where("version = ?", version).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindByIds(ctx context.Context, ids []int) ([]models.Exercise, error) {
	var result []models.Exercise

	err := r.db.Where("id IN ?", ids).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
