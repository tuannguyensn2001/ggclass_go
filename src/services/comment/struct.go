package comment

type CreateCommentInput struct {
	Content string `form:"content" binding:"required"`
	PostId  int    `form:"postId" binding:"required"`
}
