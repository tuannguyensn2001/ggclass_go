package exercise

import (
	"ggclass_go/src/config"
	"ggclass_go/src/services/class"
	"ggclass_go/src/services/exercise_clone"
	"ggclass_go/src/services/exercise_multiple_choice"
)

func BuildService() *service {
	repository := NewRepository(config.Cfg.GetDB())
	service := NewService(repository)
	exerciseCloneService := exercise_clone.BuildService()
	exerciseCloneService.SetExerciseService(service)
	exerciseCloneService.SetExerciseMultipleChoiceService(exercise_multiple_choice.BuildService())

	service.SetExerciseCloneService(exerciseCloneService)
	service.SetClassService(class.BuildService())
	return service
}
