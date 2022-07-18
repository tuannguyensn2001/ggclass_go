package comment

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"ggclass_go/src/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IService interface {
	Create(ctx context.Context, input CreateCommentInput, userId int) (*models.Comment, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) Create(ctx *gin.Context) {
	var input CreateCommentInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := auth.GetUserIdFromContext(ctx)

	if err != nil {
		panic(app.ForbiddenHttpError("forbidden", errors.New("forbidden")))
	}

	result, err := t.service.Create(ctx.Request.Context(), input, userId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}
