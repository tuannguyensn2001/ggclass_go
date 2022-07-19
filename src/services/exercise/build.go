package exercise

import "ggclass_go/src/config"

func BuildService() *service {
	repository := NewRepository(config.Cfg.GetDB())
	return NewService(repository)
}
