package models

import (
	"ggclass_go/src/enums"
	"gorm.io/gorm"
	"time"
)

type ExerciseMultipleChoiceAnswerClone struct {
	Id                       int                                    `gorm:"column_id;" json:"id"`
	ExerciseMultipleChoiceId int                                    `gorm:"column:exercise_multiple_choice_id;" json:"exerciseMultipleChoiceId"`
	Order                    int                                    `gorm:"column:order;" json:"order"`
	Type                     enums.ExerciseMultipleChoiceAnswerType `gorm:"column:type;" json:"type"`
	Answer                   string                                 `gorm:"column:answer;" json:"answer"`
	Mark                     float64                                `gorm:"column:mark;" json:"mark"`
	CreatedAt                *time.Time                             `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt                *time.Time                             `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt                *gorm.DeletedAt                        `json:"deletedAt"`
}

func (e ExerciseMultipleChoiceAnswerClone) TableName() string {
	return "exercises_multiple_choice_answers_clone"
}
