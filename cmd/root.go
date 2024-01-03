package cmd

import (
	"github.com/BuiNhatTruong99/TikTok-Go/composer"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
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

	router.Run(cfg.Server.Port)
}

func SetupRoutes(router *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	authAPIService := composer.ComposeAuthAPIService(db, cfg)

	auth := router.Group("/auth")
	{
		auth.POST("/register", authAPIService.Register())
		auth.POST("/login", authAPIService.Login())

	}
}
