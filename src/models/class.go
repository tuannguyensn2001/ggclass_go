package models

import (
	"gorm.io/gorm"
	"time"
)

type Class struct {
	Id          int             `gorm:"column:id;" json:"id"`
	Name        string          `gorm:"column:name;" json:"name"`
	Description string          `gorm:"column:description;" json:"description"`
	Room        string          `gorm:"column:room;" json:"room"`
	Topic       string          `gorm:"column:topic;" json:"topic"`
	Code        string          `gorm:"column:code;" json:"code"`
	CreatedBy   int             `gorm:"column:created_by;" json:"createdBy"`
	CreatedAt   *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt   *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `json:"deletedAt"`
}
