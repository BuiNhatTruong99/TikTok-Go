package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Server     ServerConfig
	Postgres   PosgresConfig
	Cloudinary Cloudinary
}

type ServerConfig struct {
	Port                 string
	JwtSecretKey         string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

type PosgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  string
	PgDriver           string
}

type Cloudinary struct {
	CloudName               string
	CloudAPIKey             string
	CloudAPISecretKey       string
	CloudUploadFolderVideo  string
	CloudUploadFolderAvatar string
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config-prod")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %s", err)
		return nil, err
	}

	return &config, nil
}
