package lesson

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type IRepository interface {
	Create(ctx context.Context, lesson *models.Lesson) error
	Update(ctx context.Context, lesson *models.Lesson) error
	FindById(ctx context.Context, id int) (*models.Lesson, error)
	FindByFolderId(ctx context.Context, folderId int) ([]models.Lesson, error)
}

type service struct {
	repository    IRepository
	folderService IFolderService
}

type IFolderService interface {
	GetById(ctx context.Context, id int) (*models.Folder, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, input CreateLessonInput, userId int) (*models.Lesson, error) {
	_, err := s.folderService.GetById(ctx, input.FolderId)
	if err != nil {
		return nil, err
	}

	lesson := models.Lesson{
		Name:        input.Name,
		Description: input.Description,
		FolderId:    input.FolderId,
		YoutubeLink: input.YoutubeLink,
		CreatedBy:   userId,
	}

	err = s.repository.Create(ctx, &lesson)
	if err != nil {
		return nil, err
	}
	return &lesson, nil

}

func (s *service) Edit(ctx context.Context, id int, input EditLessonInput) (*models.Lesson, error) {
	lesson, err := s.repository.FindById(ctx, id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app.NotFoundHttpError("not found lesson", err)
	}
	if err != nil {
		return nil, err
	}

	lesson.Name = input.Name
	lesson.Description = input.Description
	lesson.YoutubeLink = input.YoutubeLink
	lesson.FolderId = input.FolderId

	err = s.repository.Update(ctx, lesson)
	if err != nil {
		return nil, err
	}

	return lesson, nil

}

func (s *service) GetByFolderId(ctx context.Context, folderId int) ([]models.Lesson, error) {
	return s.repository.FindByFolderId(ctx, folderId)
}
