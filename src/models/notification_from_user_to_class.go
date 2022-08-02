package models

import "time"

type NotificationFromTeacherToClass struct {
	Id        int        `gorm:"column:id;" json:"id"`
	ClassId   int        `gorm:"column:class_id;" json:"classId"`
	Content   string     `gorm:"column:content" json:"content"`
	CreatedBy int        `gorm:"column:created_by" json:"createdBy"`
	CreatedAt *time.Time `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (NotificationFromTeacherToClass) TableName() string {
	return "notification_from_teacher_to_class"
}
