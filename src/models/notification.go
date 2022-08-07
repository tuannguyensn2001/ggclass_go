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

type NotificationV2 struct {
	Id          string                 `json:"id"`
	OwnerName   string                 `bson:"ownerName,omitempty" json:"ownerName"`
	OwnerAvatar string                 `bson:"ownerAvatar,omitempty" json:"ownerAvatar"`
	HtmlContent string                 `bson:"htmlContent,omitempty" json:"htmlContent"`
	ClassId     int                    `bson:"classId,omitempty" json:"classId"`
	CreatedBy   int                    `bson:"createdBy,omitempty" json:"createdBy"`
	Content     string                 `bson:"content,omitempty" json:"content"`
	CreatedAt   *time.Time             `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt   *time.Time             `bson:"updatedAt,omitempty" json:"updatedAt"`
	Type        enums.NotificationType `bson:"type,omitempty" json:"type"`
	Seen        int                    `json:"seen"`
}
