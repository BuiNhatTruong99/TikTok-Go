package postgresql

import (
	"fmt"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDB(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlPassword,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	return db, nil
}
