package notification

import (
	"context"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
)

type IRepository interface {
	BeginTransaction()
	Rollback()
	Commit()
	CreateNotification(ctx context.Context, notification *models.Notification) error
	CreateNotificationFromTeacherToClass(ctx context.Context, notification *models.NotificationFromTeacherToClass) error
	CreateNotificationsToUser(ctx context.Context, list *[]models.NotificationToUser) error
}

type service struct {
	repository    IRepository
	userService   IUserService
	memberService IMemberService
}

type IUserService interface {
	GetById(ctx context.Context, id int) (*models.User, error)
}

type IMemberService interface {
	GetIdMembers(ctx context.Context, classId int) ([]int, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetUserService(userService IUserService) {
	s.userService = userService
}

func (s *service) SetMemberService(memberService IMemberService) {
	s.memberService = memberService
}

func (s *service) CreateNotificationFromTeacherToClass(ctx context.Context, input createNotificationFromTeacherToClassInput, userId int) error {

	user, err := s.userService.GetById(ctx, userId)
	if err != nil {
		return err
	}

	s.repository.BeginTransaction()

	notificationFromTeacherToClass := models.NotificationFromTeacherToClass{
		ClassId:   input.ClassId,
		Content:   input.Content,
		CreatedBy: userId,
	}
	err = s.repository.CreateNotificationFromTeacherToClass(ctx, &notificationFromTeacherToClass)
	if err != nil {
		s.repository.Rollback()
		return err
	}

	notification := models.Notification{
		OwnerAvatar: user.Profile.Avatar,
		OwnerName:   user.Username,
		CreatedBy:   user.Id,
		Type:        enums.NotificationFromTeacherToClass,
		TypeId:      notificationFromTeacherToClass.Id,
		HtmlContent: "Giao vien vua them thong bao vao lop hoc",
	}
	err = s.repository.CreateNotification(ctx, &notification)
	if err != nil {
		s.repository.Rollback()
		return err
	}

	s.repository.Commit()

	members, err := s.memberService.GetIdMembers(ctx, input.ClassId)
	if err != nil {
		return err
	}

	notifications := make([]models.NotificationToUser, len(members))

	for index, item := range members {
		notifications[index] = models.NotificationToUser{
			UserId:         item,
			NotificationId: notification.Id,
			Seen:           0,
		}
	}

	go s.repository.CreateNotificationsToUser(ctx, &notifications)

	return nil
}
