package class

import (
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
)

type CreateClassInput struct {
	Name        string `form:"name" binding:"required" validate:"required"`
	Description string `form:"description"`
	Room        string `form:"room"`
	Topic       string `form:"topic"`
}

type InviteMemberInput struct {
	UserId  int             `form:"userId" binding:"required" validate:"required"`
	ClassId int             `form:"classId" binding:"required" validate:"required"`
	Role    enums.ClassRole `form:"role" binding:"required" validate:"required,min=1,max=3"`
}

type DeleteMemberInput struct {
	UserId  int `form:"userId" binding:"required" validate:"required"`
	ClassId int `form:"classId" binding:"required" validate:"required"`
}

type GetMembersOutput struct {
	models.User
	Role        enums.ClassRole `json:"role"`
	StatusClass enums.Status    `json:"statusClass"`
}

type GetMyClassOutput struct {
	models.Class
	StatusClass enums.Status `json:"statusClass"`
}
