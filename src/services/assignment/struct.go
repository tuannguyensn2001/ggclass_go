package assignment

type StartAssignmentInput struct {
	UserId     int
	ExerciseId int `form:"exerciseId" binding:"required"`
}

type createLogInput struct {
	AssignmentId int    `form:"assignmentId" binding:"required" json:"assignmentId"`
	Action       string `form:"action" binding:"required" json:"action" `
}

type userCreateAnswerInput struct {
	AssignmentId                        int    `form:"assignmentId" json:"assignmentId" binding:"required"`
	ExerciseMultipleChoiceAnswerCloneId int    `form:"exerciseMultipleChoiceAnswerCloneId" binding:"required"`
	Answer                              string `form:"answer" binding:"required"`
}
