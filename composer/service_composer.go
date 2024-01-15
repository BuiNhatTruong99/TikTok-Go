package composer

import (
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	authController "github.com/BuiNhatTruong99/TikTok-Go/internal/auth/controller"
	authPGRepository "github.com/BuiNhatTruong99/TikTok-Go/internal/auth/repository"
	authPGUsecase "github.com/BuiNhatTruong99/TikTok-Go/internal/auth/usecase"
	likeController "github.com/BuiNhatTruong99/TikTok-Go/internal/like/controller"
	likeRepository "github.com/BuiNhatTruong99/TikTok-Go/internal/like/repository"
	likeUsecase "github.com/BuiNhatTruong99/TikTok-Go/internal/like/usecase"
	postController "github.com/BuiNhatTruong99/TikTok-Go/internal/post/controller"
	postPGRepository "github.com/BuiNhatTruong99/TikTok-Go/internal/post/repository"
	postPGUsecase "github.com/BuiNhatTruong99/TikTok-Go/internal/post/usecase"
	sessionController "github.com/BuiNhatTruong99/TikTok-Go/internal/session/controller"
	sessionPGRepository "github.com/BuiNhatTruong99/TikTok-Go/internal/session/repository"
	sessionPGUsecase "github.com/BuiNhatTruong99/TikTok-Go/internal/session/usecase"
	userController "github.com/BuiNhatTruong99/TikTok-Go/internal/user/controller"
	userRepository "github.com/BuiNhatTruong99/TikTok-Go/internal/user/repository"
	userUsecase "github.com/BuiNhatTruong99/TikTok-Go/internal/user/usecase"
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

type PostService interface {
	CreatePost() func(ctx *gin.Context)
	GetPostsByUserID() func(ctx *gin.Context)
	GetAllPosts() func(ctx *gin.Context)
	DeletePost() func(ctx *gin.Context)
}

type UserService interface {
	ChangeAvatar() func(ctx *gin.Context)
	UpdateProfile() func(ctx *gin.Context)
}

type LikeService interface {
	LikePost() func(ctx *gin.Context)
	UndoLikePost() func(ctx *gin.Context)
}

func ComposeAuthAPIService(db *gorm.DB, cfg *config.Config) AuthService {
	authRepo := authPGRepository.NewAuthRepository(db)
	authUC := authPGUsecase.NewAuthUsecase(authRepo, cfg)
	sessionRepo := sessionPGRepository.NewSessionRepository(db)
	sessionUC := sessionPGUsecase.NewSessionUsecase(sessionRepo)
	authAPIController := authController.NewAuthController(authUC, sessionUC, cfg)

	return authAPIController
}

func ComposeSessionAPIService(db *gorm.DB, cfg *config.Config) SessionService {
	sessionRepo := sessionPGRepository.NewSessionRepository(db)
	sessionUC := sessionPGUsecase.NewSessionUsecase(sessionRepo)
	sessionAPIController := sessionController.NewSessionController(sessionUC, cfg)

	return sessionAPIController
}

func ComposePostAPIService(db *gorm.DB, cfg *config.Config) PostService {
	postRepo := postPGRepository.NewPostRepository(db)
	postUC := postPGUsecase.NewPostUsecase(postRepo)
	postAPIController := postController.NewPostController(postUC, cfg)

	return postAPIController
}

func ComposeUserAPIService(db *gorm.DB, cfg *config.Config) UserService {
	userRepo := userRepository.NewUserRepo(db)
	userUC := userUsecase.NewUserUseCase(userRepo)
	userAPIController := userController.NewUserController(userUC, cfg)

	return userAPIController
}

func ComposeLikeAPIService(db *gorm.DB) LikeService {
	likeRepo := likeRepository.NewLikeRepository(db)
	userRepo := userRepository.NewUserRepo(db)
	likeUC := likeUsecase.NewLikeUsecase(likeRepo, userRepo)
	likeAPIService := likeController.NewLikeController(likeUC)
	return likeAPIService
}
