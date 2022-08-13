package routes

import (
	"ggclass_go/src/config"
	"ggclass_go/src/services/assignment"
	"ggclass_go/src/services/auth"
	"ggclass_go/src/services/class"
	"ggclass_go/src/services/comment"
	"ggclass_go/src/services/exercise"
	"ggclass_go/src/services/exercise_clone"
	"ggclass_go/src/services/exercise_multiple_choice"
	"ggclass_go/src/services/folder"
	"ggclass_go/src/services/lesson"
	"ggclass_go/src/services/members"
	"ggclass_go/src/services/notification"
	"ggclass_go/src/services/post"
	"ggclass_go/src/services/profile"
	score_repository "ggclass_go/src/services/score/repository"
	score_service "ggclass_go/src/services/score/service"
	score_transport "ggclass_go/src/services/score/transport"
	"ggclass_go/src/services/user"
)

func buildAuthTransport() AuthHttpTransport {
	service := auth.BuildService()
	service.SetProfileService(profile.BuildService())
	//service.SetClassService(class.BuildService())
	transport := auth.NewHttpTransport(service)

	return transport
}

func buildClassTransport() ClassHttpTransport {

	service := class.BuildService()
	service.SetPostService(post.BuildService())
	transport := class.NewHttpTransport(service)
	return transport
}

func buildPostTransport() PostHttpTransport {

	service := post.BuildService()
	service.SetUserService(user.BuildService())
	transport := post.NewHttpTransport(service)

	return transport
}

func buildExerciseTransport() ExerciseHttpTransport {
	service := exercise.BuildService()
	exerciseCloneService := exercise_clone.BuildService()
	exerciseCloneService.SetExerciseService(service)
	exerciseCloneService.SetExerciseMultipleChoiceService(exercise_multiple_choice.BuildService())

	service.SetExerciseCloneService(exerciseCloneService)
	transport := exercise.NewHttpTransport(service)
	return transport
}

func buildCommentTransport() CommentHttpTransport {
	service := comment.BuildService()
	service.SetPostService(post.BuildService())
	service.SetUserService(user.BuildService())
	transport := comment.NewHttpTransport(service)
	return transport
}

func buildFolderTransport() FolderHttpTransport {
	service := folder.BuildService()
	service.SetClassService(class.BuildService())
	transport := folder.NewHttpTransport(service)
	return transport
}

func buildMemberTransport() MemberHttpTransport {
	service := members.BuildService()
	service.SetClassService(class.BuildService())
	service.SetUserService(user.BuildService())
	transport := members.NewHttpTransport(service)
	return transport
}

func buildAssignmentTransport() AssignmentHttpTransport {
	service := assignment.BuildService()
	exerciseCloneService := exercise_clone.BuildService()
	exerciseCloneService.SetExerciseService(exercise.BuildService())
	exerciseCloneService.SetExerciseMultipleChoiceService(exercise_multiple_choice.BuildService())
	service.SetExerciseCloneService(exerciseCloneService)
	service.SetExerciseMultipleChoiceService(exercise_multiple_choice.BuildService())

	return assignment.NewHttpTransport(service)
}

func buildNotificationTransport() NotificationHttpTransport {
	service := notification.BuildService()
	return notification.NewHttpTransport(service)
}

func buildExerciseCloneTransport() ExerciseCloneHttpTransport {
	service := exercise_clone.BuildService()
	return exercise_clone.NewHttpTransport(service)
}

func buildLessonTransport() LessonHttpTransport {
	service := lesson.BuildService()
	return lesson.NewHttpTransport(service)
}

func buildScoreTransport() ScoreHttpTransport {
	service := score_service.NewService(score_repository.NewRepository(config.Cfg.GetDB()))
	service.SetExerciseService(exercise.NewService(exercise.NewRepository(config.Cfg.GetDB())))
	service.SetUserService(user.BuildService())
	transport := score_transport.NewHttpTransport(service)

	return transport
}
