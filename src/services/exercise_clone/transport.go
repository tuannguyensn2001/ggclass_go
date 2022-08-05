package exercise_clone

import (
	"context"
	"ggclass_go/src/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	GetMultipleChoiceExerciseClone(ctx context.Context, id int) (*getMultipleChoiceOutput, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) GetMultipleChoice(ctx *gin.Context) {
	id := ctx.Param("id")
	exerciseId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.GetMultipleChoiceExerciseClone(ctx.Request.Context(), exerciseId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}
