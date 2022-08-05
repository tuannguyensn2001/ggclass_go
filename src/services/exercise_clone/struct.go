package exercise_clone

import "ggclass_go/src/models"

type getMultipleChoiceOutput struct {
	Exercise       *models.ExerciseClone               `json:"exercise"`
	MultipleChoice *models.ExerciseMultipleChoiceClone `json:"multipleChoice"`
	AnswerIds      []int                               `json:"answerIds"`
}
