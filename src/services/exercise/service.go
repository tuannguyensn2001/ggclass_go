package exercise

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/base"
	"ggclass_go/src/config"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	"ggclass_go/src/packages/validate"
	exercise_struct "ggclass_go/src/services/exercise/struct"
	"github.com/gookit/event"
	"gorm.io/gorm"
	"time"
)

type IRepository interface {
	base.IRepositoryBase
	Create(ctx context.Context, exercise *models.Exercise) error
	CreateMultipleChoice(ctx context.Context, multipleChoice *models.ExerciseMultipleChoice) error
	CreateMultipleChoiceAnswer(ctx context.Context, answers []models.ExerciseMultipleChoiceAnswer) error
	GetBeginTransaction() *gorm.DB
	GetDB() *gorm.DB
	SetDB(db *gorm.DB)
	FindById(ctx context.Context, id int) (*models.Exercise, error)
	FindMultipleChoiceById(ctx context.Context, id int) (*models.ExerciseMultipleChoice, error)
	FindAnswersByMultipleChoiceId(ctx context.Context, id int) ([]models.ExerciseMultipleChoiceAnswer, error)
	Save(ctx context.Context, exercise *models.Exercise) error
	SaveMultipleChoice(ctx context.Context, multipleChoice *models.ExerciseMultipleChoice) error
	DeleteAnswersByMultipleChoiceId(ctx context.Context, id int) error
	FindByClassId(ctx context.Context, classId int) ([]models.Exercise, error)
	FindExerciseCloneByExerciseIdAndVersion(ctx context.Context, exerciseId int, version int) (*models.ExerciseClone, error)
	FindByIds(ctx context.Context, ids []int) ([]models.Exercise, error)
	CountStudentsDoExercises(ctx context.Context, exerciseIds []int) ([]exercise_struct.CountMemberDoExercise, error)
}

type service struct {
	repository           IRepository
	exerciseCloneService IExerciseCloneService
	classService         IClassService
}

type IExerciseCloneService interface {
	StartClone(ctx context.Context, exerciseId int) (int, error)
}

type IClassService interface {
	CountStudentsInClass(ctx context.Context, classId int) (int, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetExerciseCloneService(exerciseCloneService IExerciseCloneService) {
	s.exerciseCloneService = exerciseCloneService
}

func (s *service) SetClassService(service IClassService) {
	s.classService = service
}

func (s *service) CreateMultipleChoice(ctx context.Context, input CreateExerciseMultipleChoiceInput, userId int) (*models.Exercise, error) {

	err := validate.Exec(input)
	if err != nil {
		return nil, err
	}

	exercise := models.Exercise{
		Name:                input.Name,
		Password:            input.Password,
		TimeToDo:            input.TimeToDo,
		IsTest:              input.IsTest,
		PreventViewQuestion: input.PreventViewQuestion,
		RoleStudent:         input.RoleStudent,
		NumberOfTimeToDo:    input.NumberOfTimeToDo,
		Mode:                input.Mode,
		ClassId:             input.ClassId,
		CreatedBy:           userId,
		Type:                enums.MultipleChoice,
		Version:             1,
		CanLate:             1,
		ModeSubmit:          enums.Capture,
	}

	if len(input.TimeStart) > 0 {
		value, err := time.Parse("15:04:05 02/01/2006", input.TimeStart)
		if err != nil {
			return nil, err
		}
		exercise.TimeStart = &value
	}

	if len(input.TimeEnd) > 0 {
		value, err := time.Parse("15:04:05 02/01/2006", input.TimeStart)
		if err != nil {
			return nil, err
		}
		exercise.TimeEnd = &value
	}

	db := s.repository.GetDB()
	tx := db.Begin()
	s.repository.SetDB(tx)
	defer func() {
		s.repository.SetDB(config.Cfg.GetDB())
	}()

	// xu ly create exercise multiple choice
	multipleChoice := models.ExerciseMultipleChoice{
		NumberOfQuestions: input.MultipleChoice.NumberOfQuestions,
		Mark:              input.MultipleChoice.Mark,
		FileQuestionUrl:   input.MultipleChoice.FileQuestionUrl,
	}
	err = s.repository.CreateMultipleChoice(ctx, &multipleChoice)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// xu ly create exercise
	exercise.TypeId = multipleChoice.Id
	err = s.repository.Create(ctx, &exercise)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//xu ly create answer
	answers := make([]models.ExerciseMultipleChoiceAnswer, len(input.MultipleChoice.Answers))

	for index, item := range input.MultipleChoice.Answers {
		answers[index] = models.ExerciseMultipleChoiceAnswer{
			ExerciseMultipleChoiceId: multipleChoice.Id,
			Order:                    item.Order,
			Type:                     enums.ExerciseMultipleChoiceAnswerPick,
			Answer:                   item.Answer,
			Mark:                     item.Mark,
		}
	}
	err = s.repository.CreateMultipleChoiceAnswer(ctx, answers)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	go s.exerciseCloneService.StartClone(ctx, exercise.Id)
	return &exercise, nil
}

func (s *service) GetById(ctx context.Context, id int) (*models.Exercise, error) {
	return s.repository.FindById(ctx, id)
}

func (s *service) EditMultipleChoice(ctx context.Context, id int, input editExerciseMultipleChoiceInput) error {
	err := validate.Exec(input)
	if err != nil {
		return err
	}

	exercise, err := s.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	exercise.Name = input.Name
	exercise.Password = input.Password
	exercise.TimeToDo = input.TimeToDo
	exercise.IsTest = input.IsTest
	exercise.PreventViewQuestion = input.PreventViewQuestion
	exercise.RoleStudent = input.RoleStudent
	exercise.NumberOfTimeToDo = input.NumberOfTimeToDo
	exercise.Mode = input.Mode
	exercise.Version++

	multipleChoice, err := s.repository.FindMultipleChoiceById(ctx, exercise.TypeId)
	if err != nil {
		return err
	}

	multipleChoice.FileQuestionUrl = input.MultipleChoice.FileQuestionUrl
	multipleChoice.NumberOfQuestions = input.MultipleChoice.NumberOfQuestions
	multipleChoice.Mark = input.MultipleChoice.Mark

	answers := make([]models.ExerciseMultipleChoiceAnswer, len(input.MultipleChoice.Answers))

	for index, item := range input.MultipleChoice.Answers {
		answers[index] = models.ExerciseMultipleChoiceAnswer{
			ExerciseMultipleChoiceId: multipleChoice.Id,
			Order:                    item.Order,
			Type:                     enums.ExerciseMultipleChoiceAnswerPick,
			Answer:                   item.Answer,
			Mark:                     item.Mark,
		}
	}

	s.repository.BeginTransaction()

	err = s.repository.Save(ctx, exercise)
	if err != nil {
		s.repository.Rollback()
		return err
	}

	err = s.repository.SaveMultipleChoice(ctx, multipleChoice)
	if err != nil {
		s.repository.Rollback()
		return err
	}

	err = s.repository.DeleteAnswersByMultipleChoiceId(ctx, multipleChoice.Id)
	if err != nil {
		s.repository.Rollback()
		return err
	}

	err = s.repository.CreateMultipleChoiceAnswer(ctx, answers)
	if err != nil {
		s.repository.Rollback()
		return err
	}

	s.repository.Commit()

	go s.exerciseCloneService.StartClone(ctx, exercise.Id)

	event.MustFire("update-exercise", event.M{"exerciseId": id})

	return nil
}

func (s *service) GetByClassId(ctx context.Context, classId int) ([]models.Exercise, error) {
	list, err := s.repository.FindByClassId(ctx, classId)
	if err != nil {
		return nil, err
	}

	count, err := s.classService.CountStudentsInClass(ctx, classId)
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(list))

	for _, item := range list {
		ids = append(ids, item.Id)
	}

	queryCountExercise, err := s.repository.CountStudentsDoExercises(ctx, ids)
	if err != nil {
		return nil, err
	}

	mapQueryCountExercise := make(map[int]int)

	for _, item := range queryCountExercise {
		mapQueryCountExercise[item.ExerciseId] = item.Count
	}

	for index, item := range list {
		clone, err := s.repository.FindExerciseCloneByExerciseIdAndVersion(ctx, item.Id, item.Version)
		if err != nil {
			return nil, err
		}

		list[index].ExerciseCloneId = clone.Id
		list[index].TotalMembersInClass = count
		val, ok := mapQueryCountExercise[item.Id]
		if !ok {
			list[index].TotalMembersDoExercise = 0
		} else {
			list[index].TotalMembersDoExercise = val
		}

	}

	return list, nil
}

func (s *service) GetMultipleChoiceExercise(ctx context.Context, id int) (*getMultipleChoiceOutput, error) {
	exercise, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if exercise.Type != enums.MultipleChoice {
		return nil, app.ConflictHttpError("type is not valid", errors.New("type is not valid"))
	}

	multipleChoice, err := s.repository.FindMultipleChoiceById(ctx, exercise.TypeId)
	if err != nil {
		return nil, err
	}

	output := getMultipleChoiceOutput{
		Exercise:       exercise,
		MultipleChoice: multipleChoice,
	}

	return &output, nil
}

func (s *service) GetDetailMultipleChoice(ctx context.Context, id int) (*getMultipleChoiceDetailOutput, error) {
	exercise, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if exercise.Type != enums.MultipleChoice {
		return nil, app.ConflictHttpError("type is not valid", errors.New("type is not valid"))
	}
	multipleChoice, err := s.repository.FindMultipleChoiceById(ctx, exercise.TypeId)
	if err != nil {
		return nil, err
	}

	answers, err := s.repository.FindAnswersByMultipleChoiceId(ctx, multipleChoice.Id)
	if err != nil {
		return nil, err
	}

	responseAnswer := make([]MultipleChoiceAnswerInput, len(answers))

	for index, item := range answers {
		responseAnswer[index] = MultipleChoiceAnswerInput{
			Order:  item.Order,
			Mark:   item.Mark,
			Answer: item.Answer,
		}
	}

	output := getMultipleChoiceDetailOutput{
		MultipleChoice: MultipleChoiceInput{
			NumberOfQuestions: multipleChoice.NumberOfQuestions,
			Mark:              multipleChoice.Mark,
			FileQuestionUrl:   multipleChoice.FileQuestionUrl,
			Answers:           responseAnswer,
		},
		Id:                  exercise.Id,
		Name:                exercise.Name,
		Password:            exercise.Password,
		TimeToDo:            exercise.TimeToDo,
		TimeStart:           exercise.TimeStart,
		TimeEnd:             exercise.TimeEnd,
		IsTest:              exercise.IsTest,
		PreventViewQuestion: exercise.PreventViewQuestion,
		RoleStudent:         exercise.RoleStudent,
		NumberOfTimeToDo:    exercise.NumberOfTimeToDo,
		Mode:                exercise.Mode,
		ClassId:             exercise.ClassId,
		CreatedBy:           exercise.CreatedBy,
		CreatedAt:           exercise.CreatedAt,
		UpdatedAt:           exercise.UpdatedAt,
		Type:                exercise.Type,
		TypeId:              exercise.TypeId,
	}

	return &output, nil

}

func (s *service) GetByIds(ctx context.Context, ids []int) ([]models.Exercise, error) {
	return s.repository.FindByIds(ctx, ids)
}
