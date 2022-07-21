package post

import (
	"ggclass_go/src/config"
	"ggclass_go/src/services/class"
)

func BuildService() *service {
	repository := NewRepository(config.Cfg.GetDB())
	classService := class.BuildService()
	service := NewService(repository, classService, config.Cfg.GetPusher(), config.Cfg.GetRabbitMQ())
	return service
}
