package notification

type createNotificationFromTeacherToClassInput struct {
	ClassId int    `form:"classId" binding:"required"`
	Content string `form:"content" binding:"required"`
}
