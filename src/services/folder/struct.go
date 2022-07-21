package folder

type CreateFolderInput struct {
	Name    string `form:"name" binding:"required"`
	ClassId int    `form:"classId" binding:"required"`
}

type GetFoldersQuery struct {
	classId int
}