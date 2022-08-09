package lesson

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
	Create(ctx context.Context, input CreateLessonInput, userId int) (*models.Lesson, error)
	Edit(ctx context.Context, id int, input EditLessonInput) (*models.Lesson, error)
	GetByFolderId(ctx context.Context, folderId int) ([]models.Lesson, error)
	GetDetail(ctx context.Context, id int) (*models.Lesson, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) Create(ctx *gin.Context) {
	var input CreateLessonInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		panic(err)
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

func (t *httpTransport) Edit(ctx *gin.Context) {
	var input EditLessonInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	id := ctx.Param("id")
	lessonId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.Edit(ctx.Request.Context(), lessonId, input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}

func (t *httpTransport) GetByFolderId(ctx *gin.Context) {
	folderIdStr := ctx.Query("folderId")

	folderId, err := strconv.Atoi(folderIdStr)
	if err != nil {
		panic(app.BadRequestHttpError("folder not valid", err))
	}

	result, err := t.service.GetByFolderId(ctx.Request.Context(), folderId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}

func (t *httpTransport) GetDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	lessonId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.GetDetail(ctx, lessonId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}
