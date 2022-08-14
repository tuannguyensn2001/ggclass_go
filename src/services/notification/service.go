package notification

import (
	"context"
	"encoding/json"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/config"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	notificationpb "ggclass_go/src/pb/notification"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	repository        IRepository
	userService       IUserService
	memberService     IMemberService
	assignmentService IAssignmentService
	classService      IClassService
}

type IUserService interface {
	GetById(ctx context.Context, id int) (*models.User, error)
}

type IMemberService interface {
	GetIdMembers(ctx context.Context, classId int) ([]int, error)
}

type IAssignmentService interface {
	GetById(ctx context.Context, assignmentId int) (*models.Assigment, error)
}

type IClassService interface {
	GetIdTeachersInClass(ctx context.Context, classId int) ([]int, error)
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

func (s *service) SetAssignmentService(service IAssignmentService) {
	s.assignmentService = service
}

func (s *service) SetClassService(service IClassService) {
	s.classService = service
}

func (s *service) CreateNotificationFromTeacherToClass(ctx context.Context, input createNotificationFromTeacherToClassInput, userId int) error {

	//user, err := s.userService.GetById(ctx, userId)
	//if err != nil {
	//	return err
	//}
	//
	//s.repository.BeginTransaction()
	//
	//notificationFromTeacherToClass := models.NotificationFromTeacherToClass{
	//	ClassId:   input.ClassId,
	//	Content:   input.Content,
	//	CreatedBy: userId,
	//}
	//err = s.repository.CreateNotificationFromTeacherToClass(ctx, &notificationFromTeacherToClass)
	//if err != nil {
	//	s.repository.Rollback()
	//	return err
	//}
	//
	//notification := models.Notification{
	//	OwnerAvatar: user.Profile.Avatar,
	//	OwnerName:   user.Username,
	//	CreatedBy:   user.Id,
	//	Type:        enums.NotificationFromTeacherToClass,
	//	TypeId:      notificationFromTeacherToClass.Id,
	//	HtmlContent: "Giao vien vua them thong bao vao lop hoc",
	//}
	//err = s.repository.CreateNotification(ctx, &notification)
	//if err != nil {
	//	s.repository.Rollback()
	//	return err
	//}
	//
	//s.repository.Commit()
	//
	//members, err := s.memberService.GetIdMembers(ctx, input.ClassId)
	//if err != nil {
	//	return err
	//}
	//
	//notifications := make([]models.NotificationToUser, len(members))
	//
	//for index, item := range members {
	//	notifications[index] = models.NotificationToUser{
	//		UserId:         item,
	//		NotificationId: notification.Id,
	//		Seen:           0,
	//	}
	//}
	//
	//go s.repository.CreateNotificationsToUser(ctx, &notifications)
	//
	//return nil

	user, err := s.userService.GetById(ctx, userId)
	if err != nil {
		return err
	}

	request := createNotificationInput{
		OwnerName:   user.Username,
		OwnerAvatar: user.Profile.Avatar,
		CreatedBy:   user.Id,
		HtmlContent: "Giao vien vua them thong bao",
		ClassId:     input.ClassId,
		Content:     input.Content,
		Type:        enums.NotificationFromTeacherToClass,
	}
	id, err := s.Create(ctx, request)
	if err != nil {
		return err
	}
	members, err := s.memberService.GetIdMembers(ctx, input.ClassId)
	if err != nil {
		return err
	}

	notify := notifyToUsers{
		Id:    id,
		Users: members,
	}

	rabbit := config.Cfg.GetRabbitMQ()
	if rabbit == nil {
		return errors.New("init rabbit err")
	}

	ch, err := rabbit.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare("notification", "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(notify)
	if err != nil {
		return err
	}

	err = ch.Publish("notification", "teacher_create", false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Create(ctx context.Context, input createNotificationInput) (string, error) {
	conn, err := grpc.Dial(config.Cfg.RealtimeService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}

	defer conn.Close()

	c := notificationpb.NewNotificationServiceClient(conn)

	r, err := c.Create(ctx, &notificationpb.CreateNotificationRequest{
		OwnerAvatar: input.OwnerAvatar,
		OwnerName:   input.OwnerName,
		CreatedBy:   int64(input.CreatedBy),
		Content:     input.Content,
		HtmlContent: input.HtmlContent,
		ClassId:     int64(input.ClassId),
		Type:        int64(input.Type),
	})
	if err != nil {
		return "", err
	}

	return r.Id, nil
}

func (s *service) GetByClassIdAndType(ctx context.Context, classId int, typeNotification enums.NotificationType) ([]models.NotificationV2, error) {
	conn, err := grpc.Dial(config.Cfg.RealtimeService, grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer conn.Close()
	if err != nil {
		return nil, err
	}

	c := notificationpb.NewNotificationServiceClient(conn)

	r, err := c.GetByClassAndType(ctx, &notificationpb.GetNotificationByClassAndTypeRequest{
		ClassId: int64(classId),
		Type:    int64(typeNotification),
	})
	if err != nil {
		return nil, err
	}

	var result []models.NotificationV2

	if r.Data == nil {
		return nil, nil
	}

	for _, item := range r.Data {
		createdAt := item.CreatedAt.AsTime()
		updatedAt := item.UpdatedAt.AsTime()
		result = append(result, models.NotificationV2{
			Id:          item.Id,
			OwnerName:   item.OwnerName,
			OwnerAvatar: item.OwnerAvatar,
			HtmlContent: item.HtmlContent,
			ClassId:     int(item.ClassId),
			CreatedBy:   int(item.CreatedBy),
			Content:     item.Content,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
		})
	}

	return result, nil

}

func (s *service) GetByUserId(ctx context.Context, userId int) ([]models.NotificationV2, error) {
	conn, err := grpc.Dial(config.Cfg.RealtimeService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	c := notificationpb.NewNotificationServiceClient(conn)

	r, err := c.GetByUserId(ctx, &notificationpb.GetNotificationByUserIdRequest{
		UserId: int64(userId),
	})
	if err != nil {
		return nil, err
	}

	var result []models.NotificationV2

	if r.Data == nil {
		return nil, nil
	}

	for _, item := range r.Data {
		//result = append(result,models.NotificationToUserV2{
		//	Id: item.Id,
		//	UserId: item.
		//})
		createdAt := item.CreatedAt.AsTime()
		updatedAt := item.UpdatedAt.AsTime()
		result = append(result, models.NotificationV2{
			Id:          item.Id,
			OwnerName:   item.OwnerName,
			OwnerAvatar: item.OwnerAvatar,
			HtmlContent: item.HtmlContent,
			ClassId:     int(item.ClassId),
			CreatedBy:   int(item.CreatedBy),
			Content:     item.Content,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
			Seen:        int(item.Seen),
			Type:        enums.NotificationType(item.Type),
		})
	}

	return result, nil

}

func (s *service) SetSeen(ctx context.Context, userId int, notificationId string) error {

	rabbit := config.Cfg.GetRabbitMQ()
	if rabbit == nil {
		return app.InternalHttpError("rabbit not setup", errors.New("rabbit not setup"))
	}

	ch, err := rabbit.Channel()
	if err != nil {
		return err
	}

	input := setSeenInput{
		UserId:         userId,
		NotificationId: notificationId,
	}

	body, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = ch.Publish("notification", "seen", false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateNotificationUserDoneTest(ctx context.Context, assignmentId int) error {

	assignment, err := s.assignmentService.GetById(ctx, assignmentId)
	if err != nil {
		return err
	}

	user, err := s.userService.GetById(ctx, assignment.UserId)
	if err != nil {
		return err
	}

	request := createNotificationInput{
		OwnerName:   user.Username,
		OwnerAvatar: user.Profile.Avatar,
		CreatedBy:   user.Id,
		HtmlContent: "Thi sinh vua lam xong bai thi",
		ClassId:     assignment.Exercise.ClassId,
		Content:     "lam xong bai thi",
		Type:        enums.NotificationUserDoneTest,
	}

	id, err := s.Create(ctx, request)
	if err != nil {
		return err
	}

	teachers, err := s.classService.GetIdTeachersInClass(ctx, assignment.Exercise.ClassId)
	if err != nil {
		return err
	}

	notify := notifyToUsers{
		Id:    id,
		Users: teachers,
	}

	rabbit := config.Cfg.GetRabbitMQ()
	if rabbit == nil {
		return errors.New("init rabbit err")
	}

	ch, err := rabbit.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare("notification", "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(notify)
	if err != nil {
		return err
	}

	err = ch.Publish("notification", "user_done_test", false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}
