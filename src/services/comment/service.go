package comment

import (
	"context"
	"ggclass_go/src/models"
)

type IRepository interface {
	Create(ctx context.Context, comment *models.Comment) error
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, input CreateCommentInput, userId int) (*models.Comment, error) {
	comment := models.Comment{
		Content:   input.Content,
		PostId:    input.PostId,
		CreatedBy: userId,
	}

	err := s.repository.Create(ctx, &comment)

	if err != nil {
		return nil, err
	}

	return &comment, nil

}
