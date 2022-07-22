package routes

import (
	"ggclass_go/src/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MatchRoutes(r *gin.Engine) {

	authTransport := buildAuthTransport()
	classTransport := buildClassTransport()
	postTransport := buildPostTransport()
	exerciseTransport := buildExerciseTransport()
	commentTransport := buildCommentTransport()
	folderTransport := buildFolderTransport()

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "gg_class_api",
		})
	})

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

		v1.POST("/folders", middlewares.Auth, folderTransport.Create)
		v1.GET("/folders", folderTransport.GetFolders)
	}
}
