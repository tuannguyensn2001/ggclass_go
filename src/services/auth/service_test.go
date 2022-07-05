package auth

import (
	"context"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestServiceRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := NewMockIRepository(ctrl)

	service := NewService(repository, "abc")

	t.Run("user existed", func(t *testing.T) {
		repository.EXPECT().FindByEmail(context.TODO(), "tuannguyensn2001a@gmail.com").Return(&models.User{}, nil)

		result, err := service.Register(context.Background(), RegisterInput{
			Email:    "tuannguyensn2001a@gmail.com",
			Password: "java2001",
			Username: "tuannguyensn2001",
		})

		assert.Nil(t, result)

		httpError, ok := err.(*app.HttpError)

		assert.True(t, ok)
		assert.Equal(t, http.StatusConflict, httpError.StatusCode)

	})

}
