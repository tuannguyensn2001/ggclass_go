package folder

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type IRepository interface {
	Create(ctx context.Context, folder *models.Folder) error
	FindByNameAndClassId(ctx context.Context, name string, classId int) (*models.Folder, error)
	FindByQuery(ctx context.Context, query GetFoldersQuery) ([]models.Folder, error)
}

type service struct {
	repository   IRepository
	classService IClassService
}

type IClassService interface {
	CheckClassExisted(ctx context.Context, classId int) bool
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetClassService(classService IClassService) {
	s.classService = classService
}

func (s *service) Create(ctx context.Context, input CreateFolderInput, userId int) (*models.Folder, error) {

	checkExisted := s.classService.CheckClassExisted(ctx, input.ClassId)

	if !checkExisted {
		return nil, app.NotFoundHttpError("not found class", errors.New("not found class"))
	}

	check, err := s.repository.FindByNameAndClassId(ctx, input.Name, input.ClassId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if check != nil {
		return nil, app.ConflictHttpError("folder existed in class", errors.New("existed"))
	}

	folder := models.Folder{
		Name:      input.Name,
		ClassId:   input.ClassId,
		CreatedBy: userId,
	}

	err = s.repository.Create(ctx, &folder)
	if err != nil {
		return nil, err
	}

	return &folder, nil

}

func (s *service) GetFolders(ctx context.Context, query GetFoldersQuery) ([]models.Folder, error) {
	return s.repository.FindByQuery(ctx, query)
}
