package controller

import (
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/session"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct {
	authUC    auth.Usecase
	sessionUC session.Usecase
	config    *config.Config
}

func NewAuthController(authUC auth.Usecase, sessionUC session.Usecase, config *config.Config) *authController {
	return &authController{
		authUC:    authUC,
		sessionUC: sessionUC,
		config:    config,
	}
}

func (c *authController) Register() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var userReq entity.UserRequest

		if err := ctx.ShouldBindJSON(&userReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if err := c.authUC.Register(ctx, &userReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, httpResponse.ResponseData(true))
	}
}

func (c *authController) Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var loginReq entity.UserLogin

		if err := ctx.ShouldBindJSON(&loginReq); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		userLoginResp, err := c.authUC.Login(ctx, &loginReq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		userSession, err := c.sessionUC.CreateSession(ctx, userLoginResp.SessionID, userLoginResp.User.ID,
			userLoginResp.RefreshToken,
			ctx.ClientIP(), userLoginResp.RefreshTokenExpiredAt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		userLoginResp.SessionID = userSession.ID
		ctx.JSON(http.StatusOK, userLoginResp)
	}
}
