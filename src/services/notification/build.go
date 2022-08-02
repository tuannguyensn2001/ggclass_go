package notification

import (
	"ggclass_go/src/config"
	"ggclass_go/src/services/members"
	"ggclass_go/src/services/user"
)

func BuildService() *service {
	repository := NewRepository(config.Cfg.GetDB())
	service := NewService(repository)
	service.SetMemberService(members.BuildService())
	service.SetUserService(user.BuildService())
	return service
}
