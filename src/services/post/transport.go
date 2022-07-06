package post

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IService interface {
	Create(ctx context.Context, userId int, input CreatePostInput) error
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) Create(ctx *gin.Context) {
	var input CreatePostInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		panic(app.ForbiddenHttpError("forbidden", errors.New("forbidden")))
	}

	err = t.service.Create(ctx.Request.Context(), userId, input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})

}
