package folder

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"ggclass_go/src/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	Create(ctx context.Context, input CreateFolderInput, userId int) (*models.Folder, error)
	GetFolders(ctx context.Context, query GetFoldersQuery) ([]models.Folder, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) Create(ctx *gin.Context) {
	var input CreateFolderInput
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

func (t *httpTransport) GetFolders(ctx *gin.Context) {
	query := GetFoldersQuery{}

	classIdStr := ctx.Query("classId")

	if len(classIdStr) > 0 {
		classId, err := strconv.Atoi(classIdStr)
		if err != nil {
			panic(app.BadRequestHttpError("data not valid", err))
		}
		query.classId = classId
	}

	result, err := t.service.GetFolders(ctx.Request.Context(), query)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}
