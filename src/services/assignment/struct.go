package assignment

type StartAssignmentInput struct {
	UserId     int
	ExerciseId int `form:"exerciseId" binding:"required"`
}

type createLogInput struct {
	AssignmentId int    `form:"assignmentId" binding:"required" json:"assignmentId"`
	Action       string `form:"action" binding:"required" json:"action" `
	UserId       int    `json:"userId"`
}

type userCreateAnswerInput struct {
	AssignmentId                        int    `form:"assignmentId" json:"assignmentId" binding:"required"`
	ExerciseMultipleChoiceAnswerCloneId int    `form:"exerciseMultipleChoiceAnswerCloneId" binding:"required"`
	Answer                              string `form:"answer" binding:"required"`
}

type submitMultipleChoiceInput struct {
	AssignmentId int `form:"assignmentId" binding:"required"`
	Answers      []struct {
		Id     int    `form:"id" binding:"required"`
		Answer string `form:"answer" binding:"required"`
	} `form:"answers" binding:"required"`
}

type notifyToUsers struct {
	Id    string `json:"id"`
	Users []int  `json:"users"`
}
