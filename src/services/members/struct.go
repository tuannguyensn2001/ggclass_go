package members

type JoinClassInput struct {
	UserId int
	Code   string `form:"code" binding:"required"`
}

type AcceptInviteInput struct {
	UserId  int `form:"userId" binding:"required"`
	ClassId int `form:"classId" binding:"required"`
}
