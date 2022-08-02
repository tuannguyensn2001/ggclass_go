package notification

import (
	"context"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db       *gorm.DB
	originDB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db, originDB: db}
}

func (r *repository) BeginTransaction() {
	r.db = r.originDB.Begin()
}

func (r *repository) Commit() {
	r.db.Commit()
	r.db = r.originDB
}

func (r *repository) Rollback() {
	r.db.Commit()
	r.db = r.originDB
}

func (r *repository) CreateNotification(ctx context.Context, notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *repository) CreateNotificationFromTeacherToClass(ctx context.Context, notification *models.NotificationFromTeacherToClass) error {
	return r.db.Create(notification).Error
}

func (r *repository) CreateNotificationsToUser(ctx context.Context, list *[]models.NotificationToUser) error {
	return r.db.Create(list).Error
}
