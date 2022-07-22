package post

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"github.com/pusher/pusher-http-go"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"log"
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

type IUserService interface {
	GetById(ctx context.Context, userId int) (*models.User, error)
}

type service struct {
	repository   IRepository
	classService IClassService
	userService  IUserService
	pusher       pusher.Client
	rabbitMQ     *amqp091.Connection
}

func NewService(repository IRepository, classService IClassService, pusher pusher.Client, rabbitMQ *amqp091.Connection) *service {
	return &service{repository: repository, classService: classService, pusher: pusher, rabbitMQ: rabbitMQ}
}

func (s *service) SetUserService(userService IUserService) {
	s.userService = userService
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

	user, err := s.userService.GetById(ctx, post.CreatedBy)
	if err != nil {
		return nil, err
	}
	post.CreatedByUser = *user

	go s.PushPostToRabbitMQ(&post)

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

func (s *service) GetById(ctx context.Context, id int) (*models.Post, error) {
	result, err := s.repository.FindById(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, app.NotFoundHttpError("not found post", err)
	}

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *service) PushPostToRabbitMQ(post *models.Post) {
	err := s.pusher.Trigger(fmt.Sprintf("class-%d-newsfeed", post.ClassId), "create-post", post)
	if err != nil {
		log.Println("failed to pusher", err)
	}

	return
	ch, err := s.rabbitMQ.Channel()
	defer ch.Close()
	if err != nil {
		log.Println("fail to declare channel")
		return
	}

	q, err := ch.QueueDeclare("create-post", false, false, false, false, nil)
	if err != nil {
		log.Println("fail to declare queue")
		return
	}

	data, _ := json.Marshal(post)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		},
	)
	if err != nil {
		log.Println("fail to push to queue", err)
		return
	}
}
