package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	Id        int             `gorm:"column:id;" json:"id"`
	Content   string          `gorm:"column:content;" json:"content"`
	CreatedBy int             `gorm:"column:created_by;" json:"createdBy"`
	ClassId   int             `gorm:"column:class_id" json:"classId"`
	CreatedAt *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt"`
}
