package routes

import "github.com/gin-gonic/gin"

type AuthHttpTransport interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetMe(ctx *gin.Context)
}

type ClassHttpTransport interface {
	Create(ctx *gin.Context)
	InviteMember(ctx *gin.Context)
	DeleteMember(ctx *gin.Context)
	GetMembers(ctx *gin.Context)
	AcceptInvite(ctx *gin.Context)
	GetMyClass(ctx *gin.Context)
	GetPosts(ctx *gin.Context)
	Show(ctx *gin.Context)
	GetRoles(ctx *gin.Context)
}

type PostHttpTransport interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ExerciseHttpTransport interface {
	CreateMultipleChoice(ctx *gin.Context)
}

type CommentHttpTransport interface {
	Create(ctx *gin.Context)
}

type FolderHttpTransport interface {
	Create(ctx *gin.Context)
	GetFolders(ctx *gin.Context)
}

type MemberHttpTransport interface {
	JoinClass(ctx *gin.Context)
	AcceptInvite(ctx *gin.Context)
	GetStudentsPendingByClass(ctx *gin.Context)
	AcceptAll(ctx *gin.Context)
}

type AssignmentHttpTransport interface {
	Start(ctx *gin.Context)
}
