package members

import (
	"context"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"ggclass_go/src/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	JoinClass(ctx context.Context, input JoinClassInput) error
	AcceptInvite(ctx context.Context, input AcceptInviteInput, userId int) error
	GetStudentsPendingByClass(ctx context.Context, classId int) ([]models.User, error)
	AcceptAll(ctx context.Context, classId int) error
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{
		service: service,
	}
}

func (t *httpTransport) JoinClass(ctx *gin.Context) {
	var input JoinClassInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := auth.GetUserIdFromContextWithError(ctx)
	if err != nil {
		panic(err)
	}

	input.UserId = userId
	err = t.service.JoinClass(ctx, input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})

}

func (t *httpTransport) AcceptInvite(ctx *gin.Context) {
	var input AcceptInviteInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := auth.GetUserIdFromContextWithError(ctx)
	if err != nil {
		panic(err)
	}
	err = t.service.AcceptInvite(ctx.Request.Context(), input, userId)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})

}

func (t *httpTransport) GetStudentsPendingByClass(ctx *gin.Context) {
	classIdStr := ctx.Param("id")
	classId, err := strconv.Atoi(classIdStr)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}
	result, err := t.service.GetStudentsPendingByClass(ctx.Request.Context(), classId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}

func (t *httpTransport) AcceptAll(ctx *gin.Context) {
	classIdStr := ctx.Param("id")
	classId, err := strconv.Atoi(classIdStr)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}
	err = t.service.AcceptAll(ctx.Request.Context(), classId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})
}
