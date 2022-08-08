package exercise

import (
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	"gorm.io/gorm"
	"time"
)

type CreateExerciseMultipleChoiceInput struct {
	Name                string                    `form:"name" validate:"required"`
	Password            string                    `form:"password"   `
	TimeToDo            int                       `form:"timeToDo" `
	TimeStart           string                    `form:"timeStart" `
	TimeEnd             string                    `form:"timeEnd"`
	IsTest              enums.IsTest              `form:"isTest" validate:"min=0,max=1"`
	PreventViewQuestion enums.PreventViewQuestion `form:"preventViewQuestion" validate:"min=0,max=1"`
	RoleStudent         enums.RoleStudent         `form:"roleStudent" validate:"required,min=1,max=4"`
	NumberOfTimeToDo    int                       `form:"numberOfTimeToDo"`
	Mode                enums.ExerciseMode        `form:"mode" validate:"required,min=1,max=3"`
	ClassId             int                       `form:"classId" validate:"required"`
	Type                enums.ExerciseType        `form:"type" `
	MultipleChoice      MultipleChoiceInput       `form:"multipleChoice"`
}

type MultipleChoiceInput struct {
	NumberOfQuestions int                         `form:"numberOfQuestions" validate:"required" json:"numberOfQuestions"`
	Mark              int                         `form:"mark" validate:"required" json:"mark"`
	FileQuestionUrl   string                      `form:"fileQuestionUrl" json:"fileQuestionUrl"`
	Answers           []MultipleChoiceAnswerInput `form:"answers" json:"answers"`
}

type MultipleChoiceAnswerInput struct {
	Order  int     `form:"order" validate:"required" json:"order"`
	Answer string  `form:"answer" validate:"required" json:"answer"`
	Mark   float64 `form:"mark" validate:"required" json:"mark"`
}

type createExerciseEssayInput struct {
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
	CanLate             int                       `form:"canLate" validate:"min=0,max=1"`
}

type editExerciseMultipleChoiceInput struct {
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
	MultipleChoice      MultipleChoiceInput       `form:"multipleChoice"`
}

type getMultipleChoiceOutput struct {
	Exercise       *models.Exercise               `json:"exercise"`
	MultipleChoice *models.ExerciseMultipleChoice `json:"multipleChoice"`
}

type getMultipleChoiceAnswer struct {
	Id    int `json:"id"`
	Order int `json:"order"`
}

type submitMultipleChoiceInput struct {
	AssigmentId int `form:"assignmentId" binding:"required"`
	Answers     []struct {
		Id     int    `form:"id" binding:"required"`
		Answer string `form:"answer" binding:"required"`
	} `form:"answers" binding:"required"`
}

type getMultipleChoiceDetailOutput struct {
	Id                  int                       `gorm:"column:id;" json:"id"`
	Name                string                    `gorm:"column:name;" json:"name"`
	Password            string                    `gorm:"column:password;" json:"password"`
	TimeToDo            int                       `gorm:"column:time_to_do;" json:"timeToDo"`
	TimeStart           *time.Time                `gorm:"column:time_start;" json:"timeStart"`
	TimeEnd             *time.Time                `gorm:"column:time_end;" json:"timeEnd"`
	IsTest              enums.IsTest              `gorm:"column:is_test;" json:"isTest"`
	PreventViewQuestion enums.PreventViewQuestion `gorm:"column:prevent_view_question;" json:"preventViewQuestion"`
	RoleStudent         enums.RoleStudent         `gorm:"column:role_student;" json:"roleStudent"`
	NumberOfTimeToDo    int                       `gorm:"column:number_of_time_to_do;" json:"numberOfTimeToDo"`
	Mode                enums.ExerciseMode        `gorm:"column:mode;" json:"mode"`
	ClassId             int                       `gorm:"column:class_id;" json:"classId"`
	CreatedBy           int                       `gorm:"column:created_by;" json:"createdBy"`
	CreatedAt           *time.Time                `gorm:"column:created_at;" json:"createdAt"`
	UpdatedAt           *time.Time                `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt           `json:"deletedAt"`
	Type                enums.ExerciseType        `gorm:"column:type;" json:"type"`
	TypeId              int                       `gorm:"column:type_id;" json:"typeId"`
	Version             int                       `gorm:"column:version;" json:"version"`
	CanLate             int                       `gorm:"column:can_late;" json:"canLate"`
	ModeSubmit          enums.ModeSubmit          `gorm:"column:mode_submit;" json:"modeSubmit"`
	ExerciseCloneId     int                       ` gorm:"-" json:"exerciseCloneId"`
	MultipleChoice      MultipleChoiceInput       `form:"multipleChoice" json:"multipleChoice"`
}
