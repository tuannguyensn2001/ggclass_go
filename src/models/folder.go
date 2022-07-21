package models

import (
	"gorm.io/gorm"
	"time"
)

type Folder struct {
	Id        int             `gorm:"column:id;" json:"id"`
	Name      string          `gorm:"column:name;" json:"name"`
	ClassId   int             `gorm:"column:class_id;" json:"classId"`
	CreatedBy int             `gorm:"column:created_by;" json:"createdBy"`
	CreatedAt *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt"`
}
