package composer

import (
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/controller"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/repository"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthService interface {
	Register() func(ctx *gin.Context)
	Login() func(ctx *gin.Context)
}

func ComposeAuthAPIService(db *gorm.DB, cfg *config.Config) AuthService {
	authRepo := repository.NewAuthRepository(db)
	authUC := usecase.NewAuthUsecase(authRepo, cfg)
	authController := controller.NewAuthController(authUC, cfg)

	return authController
}
