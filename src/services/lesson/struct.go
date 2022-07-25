package lesson

type CreateLessonInput struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" `
	FolderId    int    `form:"folderId" binding:"required"`
	YoutubeLink string `form:"youtubeLink" binding:"required"`
}

type EditLessonInput struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description"`
	YoutubeLink string `form:"youtubeLink" binding:"required"`
	FolderId    int    `form:"folderId" binding:"required"`
}
