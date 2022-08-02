package models

import "time"

type NotificationToUser struct {
	Id             int        `gorm:"column:id;" json:"id"`
	NotificationId int        `gorm:"column:notification_id" json:"notificationId"`
	UserId         int        `gorm:"column:user_id;" json:"userId"`
	Seen           int        `gorm:"column:seen;" json:"seen"`
	CreatedAt      *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (NotificationToUser) TableName() string {
	return "notification_to_user"
}
