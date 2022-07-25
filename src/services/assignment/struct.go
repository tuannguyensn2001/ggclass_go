package assignment

type StartAssignmentInput struct {
	UserId     int
	ExerciseId int `form:"exerciseId" binding:"required"`
}
