package score_transport

import (
	"context"
	"ggclass_go/src/app"
	score_struct "ggclass_go/src/services/score/struct"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	GetScore(ctx context.Context, classId int) ([]score_struct.GetScoreOutput, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) GetScore(ctx *gin.Context) {
	id := ctx.Param("id")
	classId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.GetScore(ctx.Request.Context(), classId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}
