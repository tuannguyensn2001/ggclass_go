package assignment

import (
	"context"
	"encoding/json"
	"errors"
	"ggclass_go/src/config"
	"ggclass_go/src/models"
	logsAssignmentpb "ggclass_go/src/pb"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IRepository interface {
	Create(ctx context.Context, assignment *models.Assigment) error
}

type service struct {
	repository    IRepository
	exerciseClone IExerciseClone
}

type IExerciseClone interface {
	GetLatestClone(ctx context.Context, exerciseId int) (*models.ExerciseClone, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetExerciseCloneService(exerciseCloneService IExerciseClone) {
	s.exerciseClone = exerciseCloneService
}

func (s *service) Start(ctx context.Context, input StartAssignmentInput) (*models.Assigment, error) {

	exerciseClone, err := s.exerciseClone.GetLatestClone(ctx, input.ExerciseId)
	if err != nil {
		return nil, err
	}

	assignment := models.Assigment{
		ExerciseId:      input.ExerciseId,
		UserId:          input.UserId,
		ExerciseCloneId: exerciseClone.Id,
	}
	err = s.repository.Create(ctx, &assignment)
	if err != nil {
		return nil, nil
	}

	return &assignment, nil
}

func (s *service) CreateLog(ctx context.Context, input createLogInput) error {
	rabbit := config.Cfg.GetRabbitMQ()
	if rabbit == nil {
		return errors.New("init rabbit err")
	}

	ch, err := rabbit.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare("logs", "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = ch.Publish("logs", "assignment", false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetLogs(ctx context.Context, assignmentId int) ([]models.LogAssignment, error) {
	conn, err := grpc.Dial(config.Cfg.LogService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	c := logsAssignmentpb.NewLogAssignmentServiceClient(conn)

	r, err := c.GetLogAssignmentByAssignment(ctx, &logsAssignmentpb.GetLogAssignmentByAssignmentRequest{AssignmentId: int64(assignmentId)})
	if err != nil {
		return nil, err
	}

	var result []models.LogAssignment

	for _, item := range r.Data {
		result = append(result, models.LogAssignment{
			AssignmentId: int(item.AssignmentId),
			Action:       item.Action,
			Id:           item.Id,
		})
	}

	return result, nil
}
