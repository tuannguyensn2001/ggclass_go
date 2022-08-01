//go:generate mockgen --source=transport.go --destination=service.mock.go --package=class
package class

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"ggclass_go/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IService interface {
	Create(ctx context.Context, input CreateClassInput, userId int) (*GetMyClassOutput, error)
	AddMember(ctx context.Context, input InviteMemberInput) (*models.User, error)
	DeleteMember(ctx context.Context, input DeleteMemberInput, userId int) error
	GetMembers(ctx context.Context, classId int) ([]GetMembersOutput, error)
	AcceptInvite(ctx context.Context, userId int, classId int) error
	GetMyClass(ctx context.Context, userId int) ([]GetMyClassOutput, error)
	GetPosts(ctx context.Context, classId int) ([]models.Post, error)
	GetById(ctx context.Context, id int) (*models.Class, error)
	GetRoles(ctx context.Context, classId int) (*GetRoleOutput, error)
}

type httpTransport struct {
	service IService
}

func NewHttpTransport(service IService) *httpTransport {
	return &httpTransport{service: service}
}

func (t *httpTransport) Create(ctx *gin.Context) {
	var input CreateClassInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, err := util.GetUserIdFromContext(ctx)

	if err != nil {
		panic(err)
	}

	class, err := t.service.Create(ctx.Request.Context(), input, userId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    class,
	})

}

func (t *httpTransport) InviteMember(ctx *gin.Context) {
	var input InviteMemberInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	result, err := t.service.AddMember(ctx.Request.Context(), input)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})
}

func (t *httpTransport) DeleteMember(ctx *gin.Context) {
	var input DeleteMemberInput
	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	userId, _ := util.GetUserIdFromContext(ctx)

	err := t.service.DeleteMember(ctx.Request.Context(), input, userId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})

}

func (t *httpTransport) GetMembers(ctx *gin.Context) {
	id := ctx.Param("id")

	classId, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	result, err := t.service.GetMembers(ctx, classId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})
}

func (t *httpTransport) AcceptInvite(ctx *gin.Context) {
	id := ctx.Param("id")

	classId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	userId, err := util.GetUserIdFromContext(ctx)

	if err != nil {
		panic(app.ForbiddenHttpError("forbidden", errors.New("forbidden")))
	}

	err = t.service.AcceptInvite(ctx, userId, classId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})
}

func (t *httpTransport) GetMyClass(ctx *gin.Context) {
	userId, err := util.GetUserIdFromContext(ctx)

	if err != nil {
		panic(app.ForbiddenHttpError("forbidden", errors.New("forbidden")))
	}

	result, err := t.service.GetMyClass(ctx.Request.Context(), userId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}

func (t *httpTransport) GetPosts(ctx *gin.Context) {
	id := ctx.Param("id")
	classId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	result, err := t.service.GetPosts(ctx.Request.Context(), classId)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})

}

func (t *httpTransport) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	classId, err := strconv.Atoi(id)
	if err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}
	result, err := t.service.GetById(ctx.Request.Context(), classId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})
}

func (t *httpTransport) GetRoles(ctx *gin.Context) {
	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		panic(err)
	}

	result, err := t.service.GetRoles(ctx.Request.Context(), userId)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
		"data":    result,
	})
}
