package models

import (
	"gorm.io/gorm"
	"time"
)

type Assigment struct {
	Id              int             `gorm:"column:id;" json:"id"`
	ExerciseId      int             `gorm:"column:exercise_id;" json:"exerciseId"`
	ExerciseCloneId int             `gorm:"column:exercise_clone_id;" json:"exerciseCloneId"`
	TimeLate        int             `gorm:"column:time_late;" json:"timeLate"`
	UserId          int             `gorm:"column:user_id;" json:"userId"`
	Mark            float64         `gorm:"column:mark;" json:"mark"`
	IsSubmit        int             `gorm:"column:is_submit;" json:"isSubmit"`
	CreatedAt       *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt       *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt       *gorm.DeletedAt `json:"deletedAt"`
	Exercise        *Exercise       `json:"exercise"`
}

func (e Assigment) TableName() string {
	return "assignments"
}
