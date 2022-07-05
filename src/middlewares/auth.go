package middlewares

import (
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/config"
	"ggclass_go/src/packages/jwt"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
	"strings"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("Authorization not valid")
	}

	return parts[1], nil
}

func Auth(ctx *gin.Context) {

	_, span := trace.StartSpan(ctx.Request.Context(), "start authorization")

	token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))

	if err != nil {
		panic(app.ForbiddenHttpError("no permission", err))
		return
	}

	userId, err := jwt.ValidateToken(config.Cfg.SecretKey(), token)

	if err != nil {
		panic(app.ForbiddenHttpError("no permission", err))
		return
	}

	span.End()
	ctx.Set("userId", userId)
	ctx.Next()
}
