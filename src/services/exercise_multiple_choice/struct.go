package exercise_multiple_choice

type CreateMultipleChoiceInput struct {
	NumberOfQuestions int ` validate:"required"`
	Mark              int ` validate:"required"`
	FileQuestionUrl   string
	Answers           []MultipleChoiceAnswerInput
}

type MultipleChoiceAnswerInput struct {
	Order  int     ` validate:"required"`
	Answer string  ` validate:"required"`
	Mark   float64 `validate:"required"`
}
