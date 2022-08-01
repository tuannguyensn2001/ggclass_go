package class

import (
	"ggclass_go/src/config"
	"ggclass_go/src/services/user"
)

func BuildService() *service {
	userService := user.BuildService()
	repository := NewRepository(config.Cfg.GetDB())
	service := NewService(repository, config.Cfg.GetRedis())
	service.SetUserService(userService)
	return service
}
