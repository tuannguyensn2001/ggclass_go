package models

import (
	"ggclass_go/src/enums"
	"time"
)

type Notification struct {
	Id          int                    `gorm:"column:id;" json:"id"`
	OwnerName   string                 `gorm:"column:owner_name;" json:"ownerName"`
	OwnerAvatar string                 `gorm:"column:owner_avatar;" json:"ownerAvatar"`
	CreatedBy   int                    `gorm:"column:created_by;" json:"createdBy"`
	Type        enums.NotificationType `gorm:"column:type;" json:"type"`
	TypeId      int                    `gorm:"column:type_id" json:"typeId"`
	HtmlContent string                 `gorm:"column:html_content" json:"htmlContent"`
	CreatedAt   *time.Time             `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt   *time.Time             `gorm:"column:updated_at" json:"updatedAt"`
}

func (Notification) TableName() string {
	return "notifications"
}
