package auth

import "ggclass_go/src/config"

func BuildService() *service {
	repository := NewRepository(config.Cfg.GetDB())
	service := NewService(repository, config.Cfg.SecretKey())
	return service
}
