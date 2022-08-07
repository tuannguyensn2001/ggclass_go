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

type NotificationToUserV2 struct {
	Id             string     `json:"id"`
	UserId         int        `bson:"userId,omitempty" json:"userId"`
	NotificationId string     `bson:"notificationId,omitempty" json:"notificationId"`
	Seen           int        `bson:"seen,omitempty" json:"seen"`
	CreatedAt      *time.Time `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt      *time.Time `bson:"updatedAt,omitempty" json:"updatedAt"`
}
