package models

import (
	"ggclass_go/src/enums"
	"time"
)

type UserClass struct {
	Id        int             `gorm:"column:id;"`
	UserId    int             `gorm:"column:user_id;"`
	ClassId   int             `gorm:"column:class_id;"`
	Role      enums.ClassRole `gorm:"column:role;"`
	Status    enums.Status    `gorm:"column:status;"`
	CreatedAt *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
}

func (UserClass) TableName() string {
	return "user_class"
}
