package composer

import (
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	authController "github.com/BuiNhatTruong99/TikTok-Go/internal/auth/controller"
	authPGRepository "github.com/BuiNhatTruong99/TikTok-Go/internal/auth/repository"
	authPGUsecase "github.com/BuiNhatTruong99/TikTok-Go/internal/auth/usecase"
	sessionController "github.com/BuiNhatTruong99/TikTok-Go/internal/session/controller"
	sessionPGRepository "github.com/BuiNhatTruong99/TikTok-Go/internal/session/repository"
	sessionPGUsecase "github.com/BuiNhatTruong99/TikTok-Go/internal/session/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthService interface {
	Register() func(ctx *gin.Context)
	Login() func(ctx *gin.Context)
}

type SessionService interface {
	ReGenerateAccessToKen() func(ctx *gin.Context)
}

func ComposeAuthAPIService(db *gorm.DB, cfg *config.Config) AuthService {
	authRepo := authPGRepository.NewAuthRepository(db)
	authUC := authPGUsecase.NewAuthUsecase(authRepo, cfg)
	sessionRepo := sessionPGRepository.NewSessionRepository(db)
	sessionUC := sessionPGUsecase.NewSessionUsecase(sessionRepo)
	authController := authController.NewAuthController(authUC, sessionUC, cfg)

	return authController
}

func ComposeSessionAPIService(db *gorm.DB, cfg *config.Config) SessionService {
	sessionRepo := sessionPGRepository.NewSessionRepository(db)
	sessionUC := sessionPGUsecase.NewSessionUsecase(sessionRepo)
	sessionController := sessionController.NewSessionController(sessionUC, cfg)

	return sessionController
}
