package assignment

import (
	"context"
	"encoding/json"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/config"
	"ggclass_go/src/models"
	logsAssignmentpb "ggclass_go/src/pb"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
)

type IRepository interface {
	Create(ctx context.Context, assignment *models.Assigment) error
	CreateMultipleChoiceAnswer(ctx context.Context, assignment *models.AssigmentMultipleChoice) error
	FindByAssignmentIdAndExerciseMultipleChoiceAnswerCloneId(ctx context.Context, assignmentId int, cloneId int) (*models.AssigmentMultipleChoice, error)
	SaveMultipleChoiceAnswer(ctx context.Context, assignment *models.AssigmentMultipleChoice) error
	FindById(ctx context.Context, id int) (*models.Assigment, error)
	Save(ctx context.Context, assignment *models.Assigment) error
	CreateListAssignmentMultipleChoice(ctx context.Context, list *[]models.AssigmentMultipleChoice) error
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

func (s *service) UserCreateAnswerMultipleChoice(ctx context.Context, input userCreateAnswerInput) error {

	check, err := s.repository.FindByAssignmentIdAndExerciseMultipleChoiceAnswerCloneId(ctx, input.AssignmentId, input.ExerciseMultipleChoiceAnswerCloneId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		assignmentMultipleChoice := models.AssigmentMultipleChoice{
			AssignmentId:                        input.AssignmentId,
			ExerciseMultipleChoiceAnswerCloneId: input.ExerciseMultipleChoiceAnswerCloneId,
			Answer:                              input.Answer,
		}

		err = s.repository.CreateMultipleChoiceAnswer(ctx, &assignmentMultipleChoice)

		return err
	}

	check.Answer = input.Answer

	err = s.repository.SaveMultipleChoiceAnswer(ctx, check)

	if err != nil {
		return err
	}

	return nil

}

func (s *service) SubmitMultipleChoiceExercise(ctx context.Context, input submitMultipleChoiceInput) error {
	assignment, err := s.repository.FindById(ctx, input.AssignmentId)
	if err != nil {
		return err
	}

	if assignment.IsSubmit == 1 {
		return app.ConflictHttpError("assignment submitted", errors.New("assignment submitted"))
	}

	assignment.IsSubmit = 1

	answers := make([]models.AssigmentMultipleChoice, 0)

	for _, item := range input.Answers {
		item := models.AssigmentMultipleChoice{
			AssignmentId:                        assignment.Id,
			ExerciseMultipleChoiceAnswerCloneId: item.Id,
			Answer:                              item.Answer,
		}
		answers = append(answers, item)
	}

	err = s.repository.Save(ctx, assignment)
	if err != nil {
		return err
	}
	err = s.repository.CreateListAssignmentMultipleChoice(ctx, &answers)
	if err != nil {
		return err
	}

	return nil

}
