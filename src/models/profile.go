package models

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	Id        int             `gorm:"column:id;" json:"id"`
	UserId    int             `gorm:"column:user_id;" json:"userId"`
	Avatar    string          `gorm:"column:avatar;" json:"avatar"`
	CreatedAt *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt"`
}
