package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id            int             `gorm:"column:id;" json:"id"`
	Content       string          `gorm:"column:content;" json:"content"`
	PostId        int             `gorm:"column:post_id;" json:"postId"`
	CreatedBy     int             `gorm:"column:created_by;" json:"createdBy"`
	CreatedAt     *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt     *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt     *gorm.DeletedAt `json:"deletedAt"`
	CreatedByUser User            `gorm:"foreignKey:CreatedBy" json:"createdByUser"`
}
