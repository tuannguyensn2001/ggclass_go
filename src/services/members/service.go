package members

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type IRepository interface {
	AddMember(ctx context.Context, userClass *models.UserClass) error
	GetByUserAndClass(ctx context.Context, userId int, classId int) (*models.UserClass, error)
	Update(ctx context.Context, userClass *models.UserClass) error
	GetStudentsPendingByClass(ctx context.Context, classId int) ([]models.UserClass, error)
	UpdateActiveByClass(ctx context.Context, classId int) error
}

type service struct {
	repository   IRepository
	classService IClassService
	userService  IUserService
}

type IClassService interface {
	GetByCode(ctx context.Context, code string) (*models.Class, error)
}
type IUserService interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]models.User, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetClassService(classService IClassService) {
	s.classService = classService
}
func (s *service) SetUserService(userService IUserService) {
	s.userService = userService
}

func (s *service) JoinClass(ctx context.Context, input JoinClassInput) error {
	class, err := s.classService.GetByCode(ctx, input.Code)
	if err != nil {
		return err
	}
	if class == nil {
		return app.NotFoundHttpError("not found class", errors.New("not found class"))
	}

	check, err := s.repository.GetByUserAndClass(ctx, input.UserId, class.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if check != nil {
		return app.ConflictHttpError("pending in class", errors.New("pending in class"))
	}

	userClass := models.UserClass{
		UserId:  input.UserId,
		ClassId: class.Id,
		Role:    enums.STUDENT,
		Status:  enums.PENDING,
	}

	err = s.repository.AddMember(ctx, &userClass)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) AcceptInvite(ctx context.Context, input AcceptInviteInput, userId int) error {
	checkIsAdmin, err := s.repository.GetByUserAndClass(ctx, userId, input.ClassId)
	if err != nil {
		return err
	}

	if checkIsAdmin.Role != enums.ADMIN {
		return app.ForbiddenHttpError("not teacher", errors.New("not teacher"))
	}

	check, err := s.repository.GetByUserAndClass(ctx, input.UserId, input.ClassId)
	if err != nil {
		return err
	}

	if check.Status == enums.ACTIVE {
		return app.ConflictHttpError("user existed", errors.New("user existed"))
	}
	if check.Status == enums.INACTIVE {
		return app.ConflictHttpError("user not in class", errors.New("user not in class"))
	}

	check.Status = enums.ACTIVE
	err = s.repository.Update(ctx, check)
	if err != nil {
		return err
	}

	return nil

}

func (s *service) GetStudentsPendingByClass(ctx context.Context, classId int) ([]models.User, error) {
	lists, err := s.repository.GetStudentsPendingByClass(ctx, classId)
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(lists))

	for index, item := range lists {
		ids[index] = item.UserId
	}

	users, err := s.userService.GetUsersByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (s *service) AcceptAll(ctx context.Context, classId int) error {
	return s.repository.UpdateActiveByClass(ctx, classId)
}
