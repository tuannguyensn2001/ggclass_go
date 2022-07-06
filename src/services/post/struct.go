package post

type CreatePostInput struct {
	Content string `form:"content" binding:"required" validate:"required"`
	ClassId int    `form:"classId" binding:"required" validate:"required"`
}
