package models

import (
	"gorm.io/gorm"
	"time"
)

type AssigmentMultipleChoice struct {
	Id                                  int             `gorm:"column:id;" json:"id"`
	AssignmentId                        int             `gorm:"column:assignment_id" json:"assignmentId"`
	ExerciseMultipleChoiceAnswerCloneId int             `gorm:"column:exercise_multiple_choice_answer_clone_id;" json:"exerciseMultipleChoiceAnswerCloneId"`
	Answer                              string          `gorm:"column:answer;" json:"answer"`
	CreatedAt                           *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt                           *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt                           *gorm.DeletedAt `json:"deletedAt"`
}

func (AssigmentMultipleChoice) TableName() string {
	return "assignment_multiple_choice"
}
