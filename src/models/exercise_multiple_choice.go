package models

import (
	"gorm.io/gorm"
	"time"
)

type ExerciseMultipleChoice struct {
	Id                int             `gorm:"column:id;" json:"id"`
	NumberOfQuestions int             `gorm:"number_of_questions;" json:"numberOfQuestions"`
	Mark              int             `gorm:"column:mark;" json:"mark"`
	FileQuestionUrl   string          `gorm:"column:file_question_url;" json:"fileQuestionUrl"`
	CreatedAt         *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt         *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt         *gorm.DeletedAt `json:"deletedAt"`
}

func (e ExerciseMultipleChoice) TableName() string {
	return "exercise_multiple_choice"
}
