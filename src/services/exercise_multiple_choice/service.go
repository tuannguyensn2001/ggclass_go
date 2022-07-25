package exercise_multiple_choice

import (
	"context"
	"ggclass_go/src/models"
	"gorm.io/gorm"
)

type IRepository interface {
	CreateMultipleChoice(ctx context.Context, multipleChoice *models.ExerciseMultipleChoice) error
	CreateMultipleChoiceAnswer(ctx context.Context, answers []models.ExerciseMultipleChoiceAnswer) error
	GetTransaction() *gorm.DB
	SetDB(db *gorm.DB)
	FindById(ctx context.Context, id int) (*models.ExerciseMultipleChoice, error)
	FindAnswersByMultipleChoiceId(ctx context.Context, id int) ([]models.ExerciseMultipleChoiceAnswer, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, input CreateMultipleChoiceInput, exerciseId int) error {

	//tx := s.repository.GetTransaction()
	//
	//s.repository.SetDB(tx)
	//
	//defer func() {
	//	s.repository.SetDB(config.Cfg.GetDB())
	//}()
	//
	//multipleChoice := models.ExerciseMultipleChoice{
	//	ExerciseId:        exerciseId,
	//	NumberOfQuestions: input.NumberOfQuestions,
	//	Mark:              input.Mark,
	//	FileQuestionUrl:   input.FileQuestionUrl,
	//}
	//
	//err := s.repository.CreateMultipleChoice(ctx, &multipleChoice)
	//
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//answers := make([]models.ExerciseMultipleChoiceAnswer, len(input.Answers))
	//
	//for index, item := range input.Answers {
	//	answers[index] = models.ExerciseMultipleChoiceAnswer{
	//		ExerciseMultipleChoiceId: multipleChoice.Id,
	//		Order:                    item.Order,
	//		Type:                     enums.ExerciseMultipleChoiceAnswerPick,
	//		Answer:                   item.Answer,
	//		Mark:                     item.Mark,
	//	}
	//}
	//
	//err = s.repository.CreateMultipleChoiceAnswer(ctx, answers)
	//
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//tx.Commit()

	return nil

}

func (s *service) GetById(ctx context.Context, id int) (*models.ExerciseMultipleChoice, error) {
	return s.repository.FindById(ctx, id)
}

func (s *service) GetAnswers(ctx context.Context, exerciseMultipleChoiceId int) ([]models.ExerciseMultipleChoiceAnswer, error) {
	return s.repository.FindAnswersByMultipleChoiceId(ctx, exerciseMultipleChoiceId)
}
