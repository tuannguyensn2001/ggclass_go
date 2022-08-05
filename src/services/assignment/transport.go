package assignment

import (
	"context"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"ggclass_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	Start(ctx context.Context, input StartAssignmentInput) (*models.Assigment, error)
	CreateLog(ctx context.Context, input createLogInput) error
	GetLogs(ctx context.Context, assignmentId int) ([]models.LogAssignment, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) Start(ctx *gin.Context) {
	var input StartAssignmentInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := util.GetUserIdFromContextWithError(ctx)
	if err != nil {
		panic(err)
	}
	input.UserId = userId
	result, err := t.service.Start(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})
}

func (t *httpTransport) CreateLog(ctx *gin.Context) {
	var input createLogInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	err := t.service.CreateLog(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})

}

func (t *httpTransport) GetLogs(ctx *gin.Context) {
	id := ctx.Param("id")
	assignmentId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}
	result, err := t.service.GetLogs(ctx.Request.Context(), assignmentId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
