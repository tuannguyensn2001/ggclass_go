package assignment

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

func (r *repository) Create(ctx context.Context, assignment *models.Assigment) error {
	return r.db.Create(assignment).Error
}

func (r *repository) CreateMultipleChoiceAnswer(ctx context.Context, assignment *models.AssigmentMultipleChoice) error {
	return r.db.Create(assignment).Error
}

func (r *repository) SaveMultipleChoiceAnswer(ctx context.Context, assignment *models.AssigmentMultipleChoice) error {
	return r.db.Save(assignment).Error
}

func (r *repository) FindByAssignmentIdAndExerciseMultipleChoiceAnswerCloneId(ctx context.Context, assignmentId int, cloneId int) (*models.AssigmentMultipleChoice, error) {
	var result models.AssigmentMultipleChoice

	err := r.db.Where("assignment_id = ?", assignmentId).Where("exercise_multiple_choice_answer_clone_id = ?", cloneId).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
