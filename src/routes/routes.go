package routes

import (
	"ggclass_go/src/middlewares"
	"ggclass_go/src/services/auth"
	"ggclass_go/src/services/class"
	"ggclass_go/src/services/comment"
	"ggclass_go/src/services/exercise"
	"ggclass_go/src/services/post"
	"github.com/gin-gonic/gin"
)

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

func MatchRoutes(r *gin.Engine) {

	authTransport := buildAuthTransport()
	classTransport := buildClassTransport()
	postTransport := buildPostTransport()
	exerciseTransport := buildExerciseTransport()
	commentTransport := buildCommentTransport()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authTransport.Register)
		v1.POST("/auth/login", authTransport.Login)
		v1.GET("/auth/me", middlewares.Auth, authTransport.GetMe)

		v1.POST("/classes", middlewares.Auth, classTransport.Create)
		v1.POST("/classes/members", middlewares.Auth, classTransport.InviteMember)
		v1.DELETE("/classes/members", middlewares.Auth, classTransport.DeleteMember)
		v1.GET("/classes/members/:id", classTransport.GetMembers)
		v1.PUT("/classes/members/accept/:id", middlewares.Auth, classTransport.AcceptInvite)
		v1.GET("/classes", middlewares.Auth, classTransport.GetMyClass)
		v1.GET("/classes/:id/posts", classTransport.GetPosts)

		v1.POST("/posts", middlewares.Auth, postTransport.Create)
		v1.DELETE("/posts/:id", middlewares.Auth, postTransport.Delete)

		v1.POST("/exercises/multiple-choice", middlewares.Auth, exerciseTransport.CreateMultipleChoice)

		v1.POST("/comments", middlewares.Auth, commentTransport.Create)
	}
}

func buildAuthTransport() AuthHttpTransport {
	service := auth.BuildService()
	transport := auth.NewHttpTransport(service)

	return transport
}

func buildClassTransport() ClassHttpTransport {

	service := class.BuildService()
	service.SetPostService(post.BuildService())
	transport := class.NewHttpTransport(service)
	return transport
}

func buildPostTransport() PostHttpTransport {

	service := post.BuildService()
	transport := post.NewHttpTransport(service)

	return transport
}

func buildExerciseTransport() ExerciseHttpTransport {
	service := exercise.BuildService()
	transport := exercise.NewHttpTransport(service)
	return transport
}

func buildCommentTransport() CommentHttpTransport {
	service := comment.BuildService()
	service.SetPostService(post.BuildService())
	transport := comment.NewHttpTransport(service)
	return transport
}
