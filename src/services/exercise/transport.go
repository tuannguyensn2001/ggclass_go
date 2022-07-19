package exercise

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
	CreateMultipleChoice(ctx context.Context, input CreateExerciseMultipleChoiceInput, userId int) (*models.Exercise, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) CreateMultipleChoice(ctx *gin.Context) {
	var input CreateExerciseMultipleChoiceInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := auth.GetUserIdFromContext(ctx)
	if err != nil {
		panic(app.ForbiddenHttpError("forbbiden", errors.New("forbidden")))
	}

	result, err := t.service.CreateMultipleChoice(ctx.Request.Context(), input, userId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}
