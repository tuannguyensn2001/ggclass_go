package exercise_clone

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type IRepository interface {
	CloneExercise(ctx context.Context, exerciseClone *models.ExerciseClone) error
	CloneExerciseMultipleChoice(ctx context.Context, clone *models.ExerciseMultipleChoiceClone) error
	CloneExerciseMultipleChoiceAnswers(ctx context.Context, answers []models.ExerciseMultipleChoiceAnswerClone) error
	BeginTransaction(ctx context.Context)
	Commit(ctx context.Context)
	Rollback(ctx context.Context)
	GetLatestByExerciseId(ctx context.Context, exerciseId int) (*models.ExerciseClone, error)
	FindById(ctx context.Context, id int) (*models.ExerciseClone, error)
	FindMultipleChoiceCloneById(ctx context.Context, id int) (*models.ExerciseMultipleChoiceClone, error)
	FindAnswersByMultipleChoiceCloneId(ctx context.Context, id int) ([]models.ExerciseMultipleChoiceAnswerClone, error)
}

type service struct {
	repository                    IRepository
	exerciseService               IExerciseService
	exerciseMultipleChoiceService IExerciseMultipleChoiceService
}

type IExerciseService interface {
	GetById(ctx context.Context, id int) (*models.Exercise, error)
}

type IExerciseMultipleChoiceService interface {
	GetById(ctx context.Context, id int) (*models.ExerciseMultipleChoice, error)
	GetAnswers(ctx context.Context, exerciseMultipleChoiceId int) ([]models.ExerciseMultipleChoiceAnswer, error)
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) SetExerciseService(exerciseService IExerciseService) {
	s.exerciseService = exerciseService
}

func (s *service) SetExerciseMultipleChoiceService(exerciseMultipleChoiceService IExerciseMultipleChoiceService) {
	s.exerciseMultipleChoiceService = exerciseMultipleChoiceService
}

func (s *service) StartClone(ctx context.Context, exerciseId int) (int, error) {
	exercise, err := s.exerciseService.GetById(ctx, exerciseId)
	if err != nil {
		return -1, err
	}

	if exercise.Type == enums.MultipleChoice {
		exerciseMultipleChoice, err := s.exerciseMultipleChoiceService.GetById(ctx, exercise.TypeId)
		if err != nil {
			return -1, err
		}

		answers, err := s.exerciseMultipleChoiceService.GetAnswers(ctx, exerciseMultipleChoice.Id)

		s.repository.BeginTransaction(ctx)
		exerciseMultipleChoiceClone := models.ExerciseMultipleChoiceClone{
			NumberOfQuestions: exerciseMultipleChoice.NumberOfQuestions,
			Mark:              exerciseMultipleChoice.Mark,
			FileQuestionUrl:   exerciseMultipleChoice.FileQuestionUrl,
		}
		err = s.repository.CloneExerciseMultipleChoice(ctx, &exerciseMultipleChoiceClone)
		if err != nil {
			s.repository.Rollback(ctx)
			return -1, err
		}
		exerciseClone := models.ExerciseClone{
			ExerciseId:          exerciseId,
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
			Type:                exercise.Type,
			TypeId:              exerciseMultipleChoiceClone.Id,
			Version:             exercise.Version,
		}

		err = s.repository.CloneExercise(ctx, &exerciseClone)
		if err != nil {
			s.repository.Rollback(ctx)
			return -1, err
		}

		answersClone := make([]models.ExerciseMultipleChoiceAnswerClone, len(answers))

		for index, item := range answers {
			answersClone[index] = models.ExerciseMultipleChoiceAnswerClone{
				ExerciseMultipleChoiceId: exerciseClone.Id,
				Order:                    item.Order,
				Type:                     item.Type,
				Answer:                   item.Answer,
				Mark:                     item.Mark,
			}
		}

		err = s.repository.CloneExerciseMultipleChoiceAnswers(ctx, answersClone)
		if err != nil {
			s.repository.Rollback(ctx)
			return -1, err
		}

		s.repository.Commit(ctx)
		return exerciseClone.Id, nil
	}

	return -1, nil
}

func (s *service) GetLatestClone(ctx context.Context, exerciseId int) (*models.ExerciseClone, error) {
	clone, err := s.repository.GetLatestByExerciseId(ctx, exerciseId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	exercise, err := s.exerciseService.GetById(ctx, exerciseId)
	if err != nil {
		return nil, err
	}

	if clone == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		id, err := s.StartClone(ctx, exerciseId)
		if err != nil {
			return nil, err
		}
		return s.repository.FindById(ctx, id)
	}
	if clone.Version != exercise.Version {
		id, err := s.StartClone(ctx, exerciseId)
		if err != nil {
			return nil, err
		}
		return s.repository.FindById(ctx, id)
	}

	return clone, nil

}

func (s *service) GetMultipleChoiceExerciseClone(ctx context.Context, id int) (*getMultipleChoiceOutput, error) {
	exerciseClone, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if exerciseClone.Type != enums.MultipleChoice {
		return nil, app.ConflictHttpError("type not valid", errors.New("type not valid"))
	}

	multipleChoiceClone, err := s.repository.FindMultipleChoiceCloneById(ctx, exerciseClone.TypeId)
	if err != nil {
		return nil, err
	}

	answers, err := s.repository.FindAnswersByMultipleChoiceCloneId(ctx, multipleChoiceClone.Id)
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(answers))

	for index, item := range answers {
		ids[index] = item.Id
	}

	output := getMultipleChoiceOutput{
		Exercise:       exerciseClone,
		MultipleChoice: multipleChoiceClone,
		AnswerIds:      ids,
	}

	return &output, nil
}

func (s *service) GetById(ctx context.Context, id int) (*models.ExerciseClone, error) {
	return s.repository.FindById(ctx, id)
}

func (s *service) GetMultipleChoiceExerciseCloneById(ctx context.Context, id int) (*models.ExerciseMultipleChoiceClone, error) {
	return s.repository.FindMultipleChoiceCloneById(ctx, id)
}

func (s *service) GetAnswersByMultipleChoiceCloneId(ctx context.Context, id int) ([]models.ExerciseMultipleChoiceAnswerClone, error) {
	return s.repository.FindAnswersByMultipleChoiceCloneId(ctx, id)
}
