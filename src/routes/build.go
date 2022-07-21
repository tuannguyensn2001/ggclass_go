package routes

import (
	"ggclass_go/src/services/auth"
	"ggclass_go/src/services/class"
	"ggclass_go/src/services/comment"
	"ggclass_go/src/services/exercise"
	"ggclass_go/src/services/folder"
	"ggclass_go/src/services/post"
	"ggclass_go/src/services/profile"
	"ggclass_go/src/services/user"
)

func buildAuthTransport() AuthHttpTransport {
	service := auth.BuildService()
	service.SetProfileService(profile.BuildService())
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
