//go:generate mockgen --source=service.go --destination=repository.mock.go --package=class

package class

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	"ggclass_go/src/packages/str"
	"ggclass_go/src/packages/validate"
)

type IRepository interface {
	Create(ctx context.Context, class *models.Class) error
	FindByNameAndCreateBy(ctx context.Context, name string, userId int) (*models.Class, error)
	AddMember(ctx context.Context, userClass *models.UserClass) error
	FindByUserAndClass(ctx context.Context, userId int, classId int) (*models.UserClass, error)
	SetMemberInactive(ctx context.Context, userId int, classId int) error
	GetUsersByClass(ctx context.Context, classId int) ([]models.UserClass, error)
	GetActiveUsersByClass(ctx context.Context, classId int) ([]models.UserClass, error)
	SetMemberPending(ctx context.Context, userId int, classId int) error
	SetMemberActive(ctx context.Context, userId int, classId int) error
	GetClassActiveByUser(ctx context.Context, userId int) ([]models.UserClass, error)
	GetClassByIds(ctx context.Context, ids []int) ([]models.Class, error)
}

type IUserService interface {
	GetUsersByIds(ctx context.Context, ids []int) ([]models.User, error)
}

type IPostService interface {
	GetPostsByClass(ctx context.Context, classId int) ([]models.Post, error)
}

type service struct {
	repository  IRepository
	userService IUserService
	postService IPostService
}

func NewService(repository IRepository, userService IUserService) *service {
	return &service{repository: repository, userService: userService}
}

func (s *service) SetPostService(postService IPostService) {
	s.postService = postService
}

func (s *service) Create(ctx context.Context, input CreateClassInput, userId int) (*models.Class, error) {
	err := validate.Exec(input)
	if err != nil {
		return nil, err
	}

	check, err := s.repository.FindByNameAndCreateBy(ctx, input.Name, userId)

	if err != nil {
		return nil, err
	}

	if check != nil {
		return nil, app.ConflictHttpError("class existed", errors.New("class existed"))
	}

	code := str.Random(5)

	class := models.Class{
		Name:        input.Name,
		Description: input.Description,
		Room:        input.Room,
		Topic:       input.Topic,
		Code:        code,
		CreatedBy:   userId,
	}

	err = s.repository.Create(ctx, &class)
	if err != nil {
		return nil, err
	}

	return &class, nil
}

func (s *service) AddMember(ctx context.Context, input InviteMemberInput) error {
	check, err := s.repository.FindByUserAndClass(ctx, input.UserId, input.ClassId)

	if err != nil {
		return nil
	}

	if check != nil {
		if check.Status == enums.PENDING {
			return app.ConflictHttpError("user invited", errors.New("user invited"))
		}
		if check.Status == enums.ACTIVE {
			return app.ConflictHttpError("user existed in class", errors.New("user existed in class"))
		}

		err := s.repository.SetMemberPending(ctx, input.UserId, input.ClassId)
		if err != nil {
			return err
		}

	} else {

		userClass := models.UserClass{
			UserId:  input.UserId,
			ClassId: input.ClassId,
			Role:    input.Role,
			Status:  enums.PENDING,
		}

		err = s.repository.AddMember(ctx, &userClass)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) DeleteMember(ctx context.Context, input DeleteMemberInput, userId int) error {

	check, err := s.repository.FindByUserAndClass(ctx, input.UserId, input.ClassId)

	if err != nil {
		return err
	}

	if check == nil {
		return app.NotFoundHttpError("user not existed in class", errors.New("not found user"))
	}

	if check.Status == enums.INACTIVE {
		return app.BadRequestHttpError("user not existed in class", errors.New("not existed"))
	}

	admin, err := s.repository.FindByUserAndClass(ctx, userId, input.ClassId)
	if err != nil {
		return err
	}

	if admin == nil {
		return app.ConflictHttpError("not permission", errors.New("not permission"))
	}

	if admin.Role != enums.ADMIN {
		return app.ConflictHttpError("not permission", errors.New("not permission"))
	}

	return s.repository.SetMemberInactive(ctx, input.UserId, input.ClassId)
}

func (s *service) GetMembers(ctx context.Context, classId int) ([]GetMembersOutput, error) {
	ids, err := s.repository.GetActiveUsersByClass(ctx, classId)

	if err != nil {
		return nil, err
	}

	mapRole := make(map[int]models.UserClass)

	userIds := make([]int, len(ids))

	for index, item := range ids {
		userIds[index] = item.UserId

		mapRole[item.UserId] = item
	}

	users, err := s.userService.GetUsersByIds(ctx, userIds)

	if err != nil {
		return nil, err
	}

	result := make([]GetMembersOutput, len(users))

	for index, item := range users {
		val, _ := mapRole[item.Id]
		result[index] = GetMembersOutput{
			User:        item,
			Role:        val.Role,
			StatusClass: val.Status,
		}
	}

	return result, nil

}

func (s *service) AcceptInvite(ctx context.Context, userId int, classId int) error {

	check, err := s.repository.FindByUserAndClass(ctx, userId, classId)

	if err != nil {
		return err
	}

	if check == nil {
		return app.ForbiddenHttpError("no permission", errors.New("no permission"))
	}

	if check.Status == enums.INACTIVE {
		return app.ConflictHttpError("not exist user", errors.New("not exist"))
	}

	if check.Status == enums.ACTIVE {
		return app.ConflictHttpError("user joined", errors.New("user joined"))
	}

	return s.repository.SetMemberActive(ctx, userId, classId)
}

func (s *service) GetMyClass(ctx context.Context, userId int) ([]GetMyClassOutput, error) {
	userClass, err := s.repository.GetClassActiveByUser(ctx, userId)

	if err != nil {
		return nil, err
	}

	ids := make([]int, len(userClass))
	mapStatus := make(map[int]models.UserClass)

	for index, item := range userClass {
		ids[index] = item.ClassId
		mapStatus[item.ClassId] = item
	}

	classes, err := s.repository.GetClassByIds(ctx, ids)

	if err != nil {
		return nil, err
	}

	result := make([]GetMyClassOutput, len(classes))

	for index, item := range classes {
		val, _ := mapStatus[item.Id]
		result[index] = GetMyClassOutput{
			Class:       item,
			StatusClass: val.Status,
		}
	}

	return result, nil

}

func (s *service) CheckUserExistedInClass(ctx context.Context, userId int, classId int) bool {
	check, err := s.repository.FindByUserAndClass(ctx, userId, classId)

	if err != nil {
		return false
	}

	if check == nil {
		return false
	}

	return true
}

func (s *service) GetPosts(ctx context.Context, classId int) ([]models.Post, error) {
	return s.postService.GetPostsByClass(ctx, classId)
}
