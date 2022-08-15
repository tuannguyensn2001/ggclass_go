package lesson

import (
	"ggclass_go/src/config"
	"ggclass_go/src/services/folder"
)

func BuildService() *service {
	repository := NewRepository(config.Cfg.GetDB())
	service := NewService(repository)
	service.folderService = folder.BuildService()
	return service
}
