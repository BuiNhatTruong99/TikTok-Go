package entity

import (
	"strings"
	"time"
)

type User struct {
	ID           int64      `gorm:"column:id;primaryKey" json:"id"`
	Username     string     `gorm:"column:username" json:"username"`
	Email        string     `gorm:"column:email" json:"email"`
	HashPassword string     `gorm:"column:hash_password" json:"hash_password"`
	AvatarUrl    string     `gorm:"column:avatar_url" json:"avatar_url"`
	Bio          string     `gorm:"column:bio" json:"bio"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
}

type ProfileRequest struct {
	Username string `gorm:"column:username" json:"username"`
	Bio      string `gorm:"column:bio" json:"bio"`
}

type AvatarRequest struct {
	AvatarUrl string `gorm:"column:avatar_url" json:"avatar_url"`
}

func TableName() string {
	return "users"
}

func (up *ProfileRequest) Validate() error {
	up.Username = strings.TrimSpace(up.Username)
	if err := validateName(up.Username); err != nil {
		return err
	}
	return nil
}
