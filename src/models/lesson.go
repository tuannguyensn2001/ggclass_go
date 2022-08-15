package models

import (
	"gorm.io/gorm"
	"time"
)

type Lesson struct {
	Id          int             `gorm:"column:id;" json:"id"`
	Name        string          `gorm:"column:name;" json:"name"`
	Description string          `gorm:"column:description;" json:"description"`
	FolderId    int             `gorm:"column:folder_id;" json:"folderId"`
	YoutubeLink string          `gorm:"column:youtube_link;" json:"youtubeLink"`
	CreatedBy   int             `gorm:"column:created_by;" json:"createdBy"`
	CreatedAt   *time.Time      `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt   *time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `json:"deletedAt"`
	Thumbnail   string          `gorm:"column:thumbnail;" json:"thumbnail"`
}
