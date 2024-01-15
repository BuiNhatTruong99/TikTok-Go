package cmd

import (
	"github.com/BuiNhatTruong99/TikTok-Go/composer"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/middleware"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/db/postgresql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func Execute() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	psqlDB, err := postgresql.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	router := gin.Default()

	v1 := router.Group("/v1")
	SetupRoutes(v1, psqlDB, cfg)

	err = router.Run(cfg.Server.Port)
	if err != nil {
		return
	}
}

func SetupRoutes(router *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	authAPIService := composer.ComposeAuthAPIService(db, cfg)
	sessionAPIService := composer.ComposeSessionAPIService(db, cfg)
	postAPIService := composer.ComposePostAPIService(db, cfg)
	userAPIService := composer.ComposeUserAPIService(db, cfg)
	likeAPIService := composer.ComposeLikeAPIService(db)

	auth := router.Group("/auth")
	{
		auth.POST("/register", authAPIService.Register())
		auth.POST("/login", authAPIService.Login())

	}

	router.POST("/tokens/new-access-token", sessionAPIService.ReGenerateAccessToKen())

	router.GET("/post", postAPIService.GetAllPosts())
	post := router.Group("/post").Use(middleware.RequireAuth(cfg))
	{
		post.GET("/:user-id", postAPIService.GetPostsByUserID())
		post.POST("/create", middleware.FileUploadMiddleware(), postAPIService.CreatePost())
		post.DELETE("/:post-id", postAPIService.DeletePost())
	}

	user := router.Group("/user").Use(middleware.RequireAuth(cfg))
	{
		user.PUT("/:user-id/change-avatar", middleware.FileUploadMiddleware(), userAPIService.ChangeAvatar())
		user.PUT("/:user-id/update-profile", userAPIService.UpdateProfile())
	}

	like := router.Group("/like").Use(middleware.RequireAuth(cfg))
	{
		like.POST("/create", likeAPIService.LikePost())
		like.DELETE("/delete", likeAPIService.UndoLikePost())
	}
}
