package middleware

import (
	"errors"
	"fmt"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationPayloadKey = "authorization_payload"
)

func RequireAuth(cfg *config.Config) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		payload, err := jwt.VerifyToken(token, cfg)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func extractTokenFromHeaderString(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("authorization header is not provided")
	}

	field := strings.Fields(s)
	if len(field) < 2 {
		return "", errors.New("invalid authorization header format")
	}

	if field[0] != "Bearer" {
		return "", fmt.Errorf("unsupported authorization type %s", field[0])

	}

	return field[1], nil
}
