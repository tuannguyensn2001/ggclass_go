package assignment

import (
	"context"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	origin *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db, origin: db}
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

func (r *repository) FindById(ctx context.Context, id int) (*models.Assigment, error) {
	var result models.Assigment
	if err := r.db.Preload("Exercise.Class").Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) Save(ctx context.Context, assignment *models.Assigment) error {
	return r.db.Save(assignment).Error
}

func (r *repository) CreateListAssignmentMultipleChoice(ctx context.Context, list *[]models.AssigmentMultipleChoice) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range *list {
			err := tx.Create(&item).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *repository) FindMultipleChoiceAnswerByAssignmentId(ctx context.Context, id int) ([]models.AssigmentMultipleChoice, error) {
	var result []models.AssigmentMultipleChoice
	err := r.db.Where("assignment_id = ?", id).Find(&result).Error
	if err != nil {
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
	r.db = r.origin
}

func (r *repository) Rollback() {
	r.db.Rollback()
	r.db = r.origin
}
