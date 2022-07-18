package comment

import (
	"context"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, comment *models.Comment) error {
	return r.db.Create(comment).Error
}
