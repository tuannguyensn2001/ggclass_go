package exercise_clone

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

func (r *repository) BeginTransaction(ctx context.Context) {
	tx := r.db.Begin()
	r.db = tx
}

func (r *repository) Commit(ctx context.Context) {
	r.db.Commit()
	r.db = r.originDB
}

func (r *repository) Rollback(ctx context.Context) {
	r.db.Rollback()
	r.db = r.originDB
}

func (r *repository) CloneExercise(ctx context.Context, exerciseClone *models.ExerciseClone) error {
	return r.db.Create(exerciseClone).Error
}

func (r *repository) CloneExerciseMultipleChoice(ctx context.Context, clone *models.ExerciseMultipleChoiceClone) error {
	return r.db.Create(clone).Error
}

func (r *repository) CloneExerciseMultipleChoiceAnswers(ctx context.Context, answers []models.ExerciseMultipleChoiceAnswerClone) error {
	return r.db.Create(answers).Error
}

func (r *repository) GetLatestByExerciseId(ctx context.Context, exerciseId int) (*models.ExerciseClone, error) {
	var result models.ExerciseClone
	if err := r.db.Where("exercise_id = ?", exerciseId).Order("id desc").First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) FindById(ctx context.Context, id int) (*models.ExerciseClone, error) {
	var result models.ExerciseClone
	if err := r.db.Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
