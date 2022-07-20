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
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (*models.Post, error)
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

func (s *service) Create(ctx context.Context, userId int, input CreatePostInput) (*models.Post, error) {

	check := s.classService.CheckUserExistedInClass(ctx, userId, input.ClassId)

	if !check {
		return nil, app.ForbiddenHttpError("forbidden", errors.New("forbidden"))
	}

	post := models.Post{
		Content:   input.Content,
		ClassId:   input.ClassId,
		CreatedBy: userId,
	}

	err := s.repository.Create(ctx, &post)

	if err != nil {
		return nil, err
	}

	return &post, nil

}

func (s *service) GetPostsByClass(ctx context.Context, classId int) ([]models.Post, error) {
	return s.repository.FindPostsByClassId(ctx, classId)
}

func (s *service) Delete(ctx context.Context, id int, userId int) error {
	post, err := s.repository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if post.Id != userId {
		return app.ForbiddenHttpError("don't have permission to delete this post", errors.New("don't have permission"))
	}

	return s.repository.Delete(ctx, id)
}
