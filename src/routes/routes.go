package routes

import (
	"ggclass_go/src/config"
	"ggclass_go/src/middlewares"
	"ggclass_go/src/services/auth"
	"ggclass_go/src/services/class"
	"ggclass_go/src/services/user"
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
}

func MatchRoutes(r *gin.Engine) {

	authTransport := buildAuthTransport()
	classTransport := buildClassTransport()

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
	}
}

func buildAuthTransport() AuthHttpTransport {
	repository := auth.NewRepository(config.Cfg.GetDB())
	service := auth.NewService(repository, config.Cfg.SecretKey())
	transport := auth.NewHttpTransport(service)

	return transport
}

func buildClassTransport() ClassHttpTransport {
	repository := class.NewRepository(config.Cfg.GetDB())

	userRepository := user.NewRepository(config.Cfg.GetDB())
	userService := user.NewService(userRepository)

	service := class.NewService(repository, userService)
	transport := class.NewHttpTransport(service)
	return transport
}
