package notification

import (
	"context"
	"ggclass_go/src/app"
	"ggclass_go/src/enums"
	"ggclass_go/src/models"
	"ggclass_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	CreateNotificationFromTeacherToClass(ctx context.Context, input createNotificationFromTeacherToClassInput, userId int) error
	GetByClassIdAndType(ctx context.Context, classId int, typeNotification enums.NotificationType) ([]models.NotificationV2, error)
	GetByUserId(ctx context.Context, userId int) ([]models.NotificationV2, error)
	SetSeen(ctx context.Context, userId int, notificationId string) error
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

func (t *httpTransport) GetNotificationFromTeacherToUser(ctx *gin.Context) {
	id := ctx.Param("id")
	classId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.GetByClassIdAndType(ctx.Request.Context(), classId, enums.NotificationFromTeacherToClass)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (t *httpTransport) GetMyNotification(ctx *gin.Context) {
	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		panic(err)
	}

	result, err := t.service.GetByUserId(ctx, userId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (t *httpTransport) SetSeen(ctx *gin.Context) {
	var input setSeenInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		panic(err)
	}

	err = t.service.SetSeen(ctx, userId, input.NotificationId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "done"})

}
