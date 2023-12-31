package cmd

import (
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/postgresql"
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

	log.Print(psqlDB)
}
