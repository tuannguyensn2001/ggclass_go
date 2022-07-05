package routes

import (
	"ggclass_go/src/config"
	"ggclass_go/src/middlewares"
	"ggclass_go/src/services/auth"
	"github.com/gin-gonic/gin"
)

type AuthHttpTransport interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetMe(ctx *gin.Context)
}

func MatchRoutes(r *gin.Engine) {

	authTransport := buildAuthTransport()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authTransport.Register)
		v1.POST("/auth/login", authTransport.Login)
		v1.GET("/auth/me", middlewares.Auth, authTransport.GetMe)
	}
}

func buildAuthTransport() AuthHttpTransport {
	repository := auth.NewRepository(config.Cfg.GetDB())
	service := auth.NewService(repository, config.Cfg.SecretKey())
	transport := auth.NewHttpTransport(service)

	return transport
}
