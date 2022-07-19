package exercise

import "ggclass_go/src/enums"

type CreateExerciseMultipleChoiceInput struct {
	Name                string                    `form:"name" validate:"required"`
	Password            string                    `form:"password"  `
	TimeToDo            int                       `form:"timeToDo" `
	TimeStart           string                    `form:"timeStart" `
	TimeEnd             string                    `form:"timeEnd"`
	IsTest              enums.IsTest              `form:"isTest" validate:"min=0,max=1"`
	PreventViewQuestion enums.PreventViewQuestion `form:"preventViewQuestion" validate:"min=0,max=1"`
	RoleStudent         enums.RoleStudent         `form:"roleStudent" validate:"required,min=1,max=4"`
	NumberOfTimeToDo    int                       `form:"numberOfTimeToDo"`
	Mode                enums.ExerciseMode        `form:"mode" validate:"required,min=1,max=3"`
	ClassId             int                       `form:"classId" validate:"required"`
	Type                enums.ExerciseType        `form:"type" validate:"required,min=1"`
	MultipleChoice      MultipleChoiceInput       `form:"multipleChoice"`
}

type MultipleChoiceInput struct {
	NumberOfQuestions int                         `form:"numberOfQuestions" validate:"required"`
	Mark              int                         `form:"mark" validate:"required"`
	FileQuestionUrl   string                      `form:"fileQuestionUrl"`
	Answers           []MultipleChoiceAnswerInput `form:"answers"`
}

type MultipleChoiceAnswerInput struct {
	Order  int     `form:"order" validate:"required"`
	Answer string  `form:"answer" validate:"required"`
	Mark   float64 `form:"mark" validate:"required"`
}
