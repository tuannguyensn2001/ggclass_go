package notification

import (
	"context"
	"ggclass_go/src/app"
	"ggclass_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IService interface {
	CreateNotificationFromTeacherToClass(ctx context.Context, input createNotificationFromTeacherToClassInput, userId int) error
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) CreateNotificationFromTeacherToClass(ctx *gin.Context) {
	var input createNotificationFromTeacherToClassInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}
	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		panic(err)
	}

	err = t.service.CreateNotificationFromTeacherToClass(ctx.Request.Context(), input, userId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})
}
