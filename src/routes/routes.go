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
	memberTransport := buildMemberTransport()
	assignmentTransport := buildAssignmentTransport()
	notificationTransport := buildNotificationTransport()
	exerciseCloneTransport := buildExerciseCloneTransport()
	lessonTransport := buildLessonTransport()

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
		v1.GET("/classes/roles", middlewares.Auth, classTransport.GetRoles)
		v1.GET("/classes/:id/role", middlewares.Auth, classTransport.GetRole)
		v1.GET("/classes/:id", classTransport.Show)

		v1.POST("/posts", middlewares.Auth, postTransport.Create)
		v1.DELETE("/posts/:id", middlewares.Auth, postTransport.Delete)

		v1.POST("/exercises/multiple-choice", middlewares.Auth, exerciseTransport.CreateMultipleChoice)
		v1.PUT("/exercises/multiple-choice/:id", middlewares.Auth, exerciseTransport.EditMultipleChoice)
		v1.GET("/exercises/multiple-choice/:id", exerciseTransport.GetMultipleChoice)
		v1.GET("/exercises/multiple-choice/:id/edit", middlewares.Auth, exerciseTransport.GetDetailMultipleChoice)
		v1.GET("/exercises", exerciseTransport.GetByClassId)

		v1.GET("/exercises/clone/multiple-choice/:id", exerciseCloneTransport.GetMultipleChoice)

		v1.POST("/comments", middlewares.Auth, commentTransport.Create)

		v1.POST("/folders", middlewares.Auth, folderTransport.Create)
		v1.GET("/folders", folderTransport.GetFolders)

		v1.POST("/members", middlewares.Auth, memberTransport.JoinClass)
		v1.PUT("/members", middlewares.Auth, memberTransport.AcceptInvite)
		v1.GET("/members/class/:id/pending", memberTransport.GetStudentsPendingByClass)
		v1.POST("/members/class/:id/accept", middlewares.Auth, memberTransport.AcceptAll)

		v1.POST("/assignments/start", middlewares.Auth, assignmentTransport.Start)
		v1.POST("/assignments/logs", assignmentTransport.CreateLog)
		v1.GET("/assignments/:id/logs", assignmentTransport.GetLogs)
		v1.POST("/assignments/submit/multiple-choice", assignmentTransport.SubmitMultipleChoiceExercise)

		v1.POST("/notifications/from-teacher-to-class", middlewares.Auth, notificationTransport.CreateNotificationFromTeacherToClass)
		v1.GET("/notifications/class/:id/from-teacher-to-class", notificationTransport.GetNotificationFromTeacherToUser)
		v1.GET("/notifications", middlewares.Auth, notificationTransport.GetMyNotification)
		v1.PUT("/notifications/seen", middlewares.Auth, notificationTransport.SetSeen)

		v1.POST("/lessons", middlewares.Auth, lessonTransport.Create)
		v1.PUT("/lessons/:id", middlewares.Auth, lessonTransport.Edit)
		v1.GET("/lessons", lessonTransport.GetByFolderId)
		v1.GET("/lessons/:id", lessonTransport.GetDetail)
	}
}
