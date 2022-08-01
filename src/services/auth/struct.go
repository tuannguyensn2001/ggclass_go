package auth

import (
	"ggclass_go/src/models"
	"ggclass_go/src/services/class"
)

type RegisterInput struct {
	Username string `form:"username" binding:"required" validate:"required"`
	Password string `form:"password" binding:"required" validate:"required"`
	Email    string `form:"email" binding:"required" validate:"required,email"`
}

type LoginInput struct {
	Password string `form:"password" binding:"required" validate:"required"`
	Email    string `form:"email" binding:"required" validate:"required,email"`
}

type LoginOutput struct {
	AccessToken string       `json:"accessToken"`
	User        *models.User `json:"user"`
}

type GetMeOutput struct {
	*models.User
	Classes []class.GetMyClassOutput `json:"classes"`
}
