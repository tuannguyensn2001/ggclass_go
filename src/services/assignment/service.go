package assignment

import (
	"context"
	"encoding/json"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/base"
	"ggclass_go/src/config"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	logsAssignmentpb "ggclass_go/src/pb/logs_assignment"
	"github.com/gookit/event"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
	"time"
)

type IRepository interface {
	base.IRepositoryBase
	Create(ctx context.Context, assignment *models.Assigment) error
	CreateMultipleChoiceAnswer(ctx context.Context, assignment *models.AssigmentMultipleChoice) error
	FindByAssignmentIdAndExerciseMultipleChoiceAnswerCloneId(ctx context.Context, assignmentId int, cloneId int) (*models.AssigmentMultipleChoice, error)
	SaveMultipleChoiceAnswer(ctx context.Context, assignment *models.AssigmentMultipleChoice) error
	FindById(ctx context.Context, id int) (*models.Assigment, error)
	Save(ctx context.Context, assignment *models.Assigment) error
	CreateListAssignmentMultipleChoice(ctx context.Context, list *[]models.AssigmentMultipleChoice) error
	FindMultipleChoiceAnswerByAssignmentId(ctx context.Context, id int) ([]models.AssigmentMultipleChoice, error)
}

type service struct {
	repository                    IRepository
	exerciseClone                 IExerciseClone
	exerciseMultipleChoiceService IExerciseMultipleChoiceService
}

type IExerciseClone interface {
	GetLatestClone(ctx context.Context, exerciseId int) (*models.ExerciseClone, error)
	GetById(ctx context.Context, id int) (*models.ExerciseClone, error)
	GetMultipleChoiceExerciseCloneById(ctx context.Context, id int) (*models.ExerciseMultipleChoiceClone, error)
	GetAnswersByMultipleChoiceCloneId(ctx context.Context, id int) ([]models.ExerciseMultipleChoiceAnswerClone, error)
}

type IExerciseMultipleChoiceService interface {
	GetMark(ctx context.Context, answers []models.ExerciseMultipleChoiceAnswerClone, assignmentAnswers []models.AssigmentMultipleChoice) float64
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetExerciseCloneService(exerciseCloneService IExerciseClone) {
	s.exerciseClone = exerciseCloneService
}

func (s *service) SetExerciseMultipleChoiceService(exerciseMultipleChoiceService IExerciseMultipleChoiceService) {
	s.exerciseMultipleChoiceService = exerciseMultipleChoiceService
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

func (s *service) GetLogs(ctx context.Context, assignmentId int, userId int) ([]models.LogAssignment, error) {
	conn, err := grpc.Dial(config.Cfg.LogService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	c := logsAssignmentpb.NewLogAssignmentServiceClient(conn)

	r, err := c.GetLogAssignmentByAssignment(ctx, &logsAssignmentpb.GetLogAssignmentByAssignmentRequest{AssignmentId: int64(assignmentId), UserId: int64(userId)})
	if err != nil {
		return nil, err
	}

	var result []models.LogAssignment

	for _, item := range r.Data {
		result = append(result, models.LogAssignment{
			AssignmentId: int(item.AssignmentId),
			Action:       item.Action,
			Id:           item.Id,
			UserId:       int(item.UserId),
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
		now := time.Now()
		item := models.AssigmentMultipleChoice{
			AssignmentId:                        assignment.Id,
			ExerciseMultipleChoiceAnswerCloneId: item.Id,
			Answer:                              item.Answer,
			CreatedAt:                           &now,
			UpdatedAt:                           &now,
		}
		answers = append(answers, item)
	}

	s.repository.BeginTransaction()

	err = s.repository.Save(ctx, assignment)
	if err != nil {
		s.repository.Rollback()
		return err
	}
	err = s.repository.CreateListAssignmentMultipleChoice(ctx, &answers)
	if err != nil {
		s.repository.Rollback()
		return err
	}

	mark, err := s.MarkAssignment(ctx, assignment)
	if err != nil {
		s.repository.Rollback()
		return err
	}
	assignment.Mark = mark
	err = s.repository.Save(ctx, assignment)
	if err != nil {
		s.repository.Rollback()
		return err
	}

	s.repository.Commit()

	event.MustFire("user_done_test", event.M{"assignmentId": input.AssignmentId})

	return nil

}

func (s *service) MarkAssignment(ctx context.Context, assignment *models.Assigment) (float64, error) {
	assignmentId := assignment.Id
	exerciseClone, err := s.exerciseClone.GetById(ctx, assignment.ExerciseCloneId)
	if err != nil {
		return -1, err
	}

	if exerciseClone.Type == enums.MultipleChoice {
		exerciseMultipleChoiceClone, err := s.exerciseClone.GetMultipleChoiceExerciseCloneById(ctx, exerciseClone.TypeId)
		if err != nil {
			return -1, err
		}
		answers, err := s.exerciseClone.GetAnswersByMultipleChoiceCloneId(ctx, exerciseMultipleChoiceClone.Id)
		if err != nil {
			return -1, err
		}
		assignmentAnswers, err := s.repository.FindMultipleChoiceAnswerByAssignmentId(ctx, assignmentId)
		if err != nil {
			return -1, err
		}
		mark := s.exerciseMultipleChoiceService.GetMark(ctx, answers, assignmentAnswers)

		return mark, nil
	}

	return -1, errors.New("not valid")

}

func (s *service) GetById(ctx context.Context, assignmentId int) (*models.Assigment, error) {
	return s.repository.FindById(ctx, assignmentId)
}
