package controller

import (
	"errors"
	"fmt"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/session"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/session/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/httpResponse"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type sessionController struct {
	sessionUC session.Usecase
	cfg       *config.Config
}

func NewSessionController(repository session.Usecase, cfg *config.Config) *sessionController {
	return &sessionController{sessionUC: repository, cfg: cfg}
}

func (s *sessionController) ReGenerateAccessToKen() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var req entity.NewAccessTokenRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		payload, err := jwt.VerifyToken(req.RefreshToken, s.cfg)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		session, err := s.sessionUC.GetSession(ctx, payload.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, httpResponse.ErrorResponse{Message: err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}
		if session.IsBlocked {
			err := fmt.Errorf("blocked session")
			ctx.JSON(http.StatusUnauthorized, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}
		if session.UserID != payload.UserID {
			err := fmt.Errorf("incorrect session user")
			ctx.JSON(http.StatusUnauthorized, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		if time.Now().After(session.ExpiresAt) {
			err := fmt.Errorf("expired session")
			ctx.JSON(http.StatusUnauthorized, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		accessToken, accessPayload, err := jwt.GenerateToken(payload.UserID, s.cfg.Server.AccessTokenDuration, s.cfg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{Message: err.Error()})
			return
		}

		resp := entity.NewAccessTokenResponse{
			AccessToken:          accessToken,
			AccessTokenExpiredAt: accessPayload.ExpiresAt,
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
