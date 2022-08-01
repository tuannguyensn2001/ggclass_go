package post

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"ggclass_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	Create(ctx context.Context, userId int, input CreatePostInput) (*models.Post, error)
	Delete(ctx context.Context, id int, userId int) error
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

	userId, err := util.GetUserIdFromContext(ctx)

	if err != nil {
		panic(app.ForbiddenHttpError("forbidden", errors.New("forbidden")))
	}

	result, err := t.service.Create(ctx.Request.Context(), userId, input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}

func (t *httpTransport) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("id not valid", errors.New("id not valid")))
	}

	userId, err := util.GetUserIdFromContext(ctx)

	if err != nil {
		panic(app.ForbiddenHttpError("forbidden", errors.New("forbidden")))
	}

	err = t.service.Delete(ctx.Request.Context(), postId, userId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})

}
