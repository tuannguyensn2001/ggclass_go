package assignment

type StartAssignmentInput struct {
	UserId     int
	ExerciseId int `form:"exerciseId" binding:"required"`
}

type createLogInput struct {
	AssignmentId int    `form:"assignmentId" binding:"required" json:"assignmentId"`
	Action       string `form:"action" binding:"required" json:"action" `
}
