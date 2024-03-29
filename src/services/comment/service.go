package comment

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
)

type IRepository interface {
	Create(ctx context.Context, comment *models.Comment) error
}

type service struct {
	repository  IRepository
	postService IPostService
	userService IUserService
}

type IPostService interface {
	GetById(ctx context.Context, id int) (*models.Post, error)
}

type IUserService interface {
	GetById(ctx context.Context, id int) (*models.User, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetPostService(postService IPostService) {
	s.postService = postService
}

func (s *service) SetUserService(userService IUserService) {
	s.userService = userService
}

func (s *service) Create(ctx context.Context, input CreateCommentInput, userId int) (*models.Comment, error) {

	post, err := s.postService.GetById(ctx, input.PostId)

	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, app.NotFoundHttpError("not found", errors.New("not found"))
	}

	comment := models.Comment{
		Content:   input.Content,
		PostId:    input.PostId,
		CreatedBy: userId,
	}

	err = s.repository.Create(ctx, &comment)

	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	comment.CreatedByUser = *user

	return &comment, nil

}
