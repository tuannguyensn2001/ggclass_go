package auth

import (
	"errors"
	"ggclass_go/src/app"
	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get("userId")

	if !ok {
		return -1, app.ForbiddenHttpError("forbidden", errors.New("forbidden"))
	}

	return userId.(int), nil

}
