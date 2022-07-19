package models

import (
	"ggclass_go/src/enums"
	"gorm.io/gorm"
	"time"
)

type Exercise struct {
	Id                  int                       `gorm:"column:id;" json:"id"`
	Name                string                    `gorm:"column:name;" json:"name"`
	Password            string                    `gorm:"column:password;" json:"password"`
	TimeToDo            int                       `gorm:"column:time_to_do;" json:"timeToDo"`
	TimeStart           *time.Time                `gorm:"column:time_start;" json:"timeStart"`
	TimeEnd             *time.Time                `gorm:"column:time_end;" json:"timeEnd"`
	IsTest              enums.IsTest              `gorm:"column:is_test;" json:"isTest"`
	PreventViewQuestion enums.PreventViewQuestion `gorm:"column:prevent_view_question;" json:"preventViewQuestion"`
	RoleStudent         enums.RoleStudent         `gorm:"column:role_student;" json:"roleStudent"`
	NumberOfTimeToDo    int                       `gorm:"column:number_of_time_to_do;" json:"numberOfTimeToDo"`
	Mode                enums.ExerciseMode        `gorm:"column:mode;" json:"mode"`
	ClassId             int                       `gorm:"column:class_id;" json:"classId"`
	CreatedBy           int                       `gorm:"column:created_by;" json:"createdBy"`
	CreatedAt           *time.Time                `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt           *time.Time                `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt           `json:"deletedAt"`
	Type                enums.ExerciseType        `gorm:"column:type;" json:"type"`
	TypeId              int                       `gorm:"column:type_id;" json:"typeId"`
}

func (e Exercise) TableName() string {
	return "exercise"
}
