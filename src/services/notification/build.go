package notification

import (
	"ggclass_go/src/config"
	"ggclass_go/src/services/assignment"
	"ggclass_go/src/services/class"
	"ggclass_go/src/services/members"
	"ggclass_go/src/services/user"
)

func BuildService() *service {
	repository := NewRepository(config.Cfg.GetDB())
	service := NewService(repository)
	service.SetMemberService(members.BuildService())
	service.SetUserService(user.BuildService())
	service.SetAssignmentService(assignment.BuildService())
	service.SetClassService(class.BuildService())
	return service
}
