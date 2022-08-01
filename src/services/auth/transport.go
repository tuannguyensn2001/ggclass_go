package auth

import (
	"context"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IService interface {
	Register(ctx context.Context, input RegisterInput) (*models.User, error)
	Login(ctx context.Context, input LoginInput) (*LoginOutput, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{
		service: service,
	}
}

func (t *httpTransport) Register(ctx *gin.Context) {
	var input RegisterInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	user, err := t.service.Register(ctx, input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "register successfully",
		"data":    user,
	})

}

func (t *httpTransport) Login(ctx *gin.Context) {

	var input LoginInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.Login(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully",
		"data":    result,
	})

}

func (t *httpTransport) GetMe(ctx *gin.Context) {

	//userId, ok := ctx.Get("userId")
	//if !ok {
	//	panic(app.ForbiddenHttpError("forbidden", errors.New("forbidden")))
	//}
	//
	//_, span := trace.StartSpan(ctx.Request.Context(), "start get me")
	//defer span.End()
	//
	//user, err := t.service.GetUserById(ctx.Request.Context(), userId.(int))
	//if err != nil {
	//	panic(err)
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"message": "done",
	//	"data":    user,
	//})

}
