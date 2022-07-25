package exercise_multiple_choice

import "ggclass_go/src/config"

func BuildService() *service {
	repository := NewRepository(config.Cfg.GetDB())
	service := NewService(repository)
	return service
}
