package class

import (
	"testing"
)

func TestCreate(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//
	//repository := NewMockIRepository(ctrl)
	//
	//service := NewService(repository)
	//
	//t.Run("validate input", func(t *testing.T) {
	//
	//	tests := []struct {
	//		input CreateClassInput
	//	}{
	//		{
	//			input: CreateClassInput{},
	//		},
	//		{
	//			input: CreateClassInput{
	//				Description: "abc",
	//			},
	//		},
	//	}
	//
	//	for _, item := range tests {
	//		_, err := service.Create(context.TODO(), item.input)
	//
	//		_, ok := err.(validator.ValidationErrors)
	//
	//		assert.True(t, ok)
	//	}
	//})
	//
	//t.Run("create success", func(t *testing.T) {
	//	//repository.EXPECT().Create(context.TODO(), nil).Return(nil)
	//
	//	class := models.Class{
	//		Name: "lop hoc 1",
	//	}
	//
	//	repository.EXPECT().Create(context.TODO(), &class).Return(nil)
	//
	//	input := CreateClassInput{
	//		Name: "lop hoc 1",
	//	}
	//
	//	result, err := service.Create(context.TODO(), input)
	//
	//	assert.Nil(t, err)
	//	assert.Equal(t, input.Name, result.Name)
	//
	//})
}
