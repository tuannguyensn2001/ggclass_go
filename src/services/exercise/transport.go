package exercise

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"ggclass_go/src/packages/logger"
	"ggclass_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	CreateMultipleChoice(ctx context.Context, input CreateExerciseMultipleChoiceInput, userId int) (*models.Exercise, error)
	EditMultipleChoice(ctx context.Context, id int, input editExerciseMultipleChoiceInput) error
	GetByClassId(ctx context.Context, classId int) ([]models.Exercise, error)
	GetMultipleChoiceExercise(ctx context.Context, id int) (*getMultipleChoiceOutput, error)
	GetDetailMultipleChoice(ctx context.Context, id int) (*getMultipleChoiceDetailOutput, error)
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

	userId, err := util.GetUserIdFromContext(ctx)
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

func (t *httpTransport) EditMultipleChoice(ctx *gin.Context) {
	var input editExerciseMultipleChoiceInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	id := ctx.Param("id")
	exerciseId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	err = t.service.EditMultipleChoice(ctx, exerciseId, input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})

}

func (t *httpTransport) GetByClassId(ctx *gin.Context) {
	classIdStr := ctx.Query("classId")
	if len(classIdStr) > 0 {
		classId, err := strconv.Atoi(classIdStr)
		if err != nil {
			panic(app.BadRequestHttpError("data not valid", err))
		}
		result, err := t.service.GetByClassId(ctx, classId)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "done",
			"data":    result,
		})
	}
}

func (t *httpTransport) GetMultipleChoice(ctx *gin.Context) {
	id := ctx.Param("id")
	exerciseId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.GetMultipleChoiceExercise(ctx.Request.Context(), exerciseId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}

func (t *httpTransport) GetDetailMultipleChoice(ctx *gin.Context) {
	logger.Sugar().Info("request")
	id := ctx.Param("id")
	exerciseId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.GetDetailMultipleChoice(ctx, exerciseId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "done",
	})

}
