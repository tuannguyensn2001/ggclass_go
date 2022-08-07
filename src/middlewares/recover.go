package middlewares

import (
	"ggclass_go/src/app"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Header("Content-Type", "application/json")

			log.Println(err)

			if httpError, ok := err.(*app.HttpError); ok {
				ctx.AbortWithStatusJSON(httpError.StatusCode, httpError)
				return
			}

			if validationError, ok := err.(validator.ValidationErrors); ok {
				ctx.AbortWithStatusJSON(422, app.HttpError{
					Message: validationError[0].Error(),
				})
				return
			}

			httpError := app.InternalHttpError("internal server", err.(error))
			ctx.AbortWithStatusJSON(httpError.StatusCode, httpError)
			return
		}
	}()

	ctx.Next()
}
