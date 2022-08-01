package util

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

func GetUserIdFromContextWithError(ctx *gin.Context) (int, error) {
	result, err := GetUserIdFromContext(ctx)
	if err != nil {
		return -1, app.ForbiddenHttpError("forbidden", errors.New("forbidden"))
	}
	return result, nil
}
