package notification

import "ggclass_go/src/enums"

type createNotificationFromTeacherToClassInput struct {
	ClassId int    `form:"classId" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type createNotificationInput struct {
	OwnerName   string
	OwnerAvatar string
	CreatedBy   int
	HtmlContent string
	ClassId     int
	Content     string
	Type        enums.NotificationType
}

type notifyToUsers struct {
	Id    string `json:"id"`
	Users []int  `json:"users"`
}

type setSeenInput struct {
	UserId         int    `json:"userId"`
	NotificationId string `json:"notificationId" form:"notificationId" binding:"required"`
}
