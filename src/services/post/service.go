package post

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
)

type IRepository interface {
	Create(ctx context.Context, post *models.Post) error
	FindPostsByClassId(ctx context.Context, classId int) ([]models.Post, error)
}

type IClassService interface {
	CheckUserExistedInClass(ctx context.Context, userId int, classId int) bool
}

type service struct {
	repository   IRepository
	classService IClassService
}

func NewService(repository IRepository, classService IClassService) *service {
	return &service{repository: repository, classService: classService}
}

func (s *service) Create(ctx context.Context, userId int, input CreatePostInput) error {

	check := s.classService.CheckUserExistedInClass(ctx, userId, input.ClassId)

	if !check {
		return app.ForbiddenHttpError("forbidden", errors.New("forbidden"))
	}

	post := models.Post{
		Content:   input.Content,
		ClassId:   input.ClassId,
		CreatedBy: userId,
	}

	err := s.repository.Create(ctx, &post)

	if err != nil {
		return err
	}

	return nil

}

func (s *service) GetPostsByClass(ctx context.Context, classId int) ([]models.Post, error) {
	return s.repository.FindPostsByClassId(ctx, classId)
}