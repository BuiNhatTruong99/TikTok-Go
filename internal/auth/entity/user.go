package entity

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

type UserRequest struct {
	Username     string `gorm:"column:username" json:"username"`
	Email        string `gorm:"column:email" json:"email"`
	HashPassword string `gorm:"column:hash_password" json:"hash_password"`
	AvatarUrl    string `gorm:"column:avatar_url" json:"avatar_url"`
	Bio          string `gorm:"column:bio" json:"bio"`
}

type UserLogin struct {
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
}

type LoginResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func TableName() string {
	return "users"
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
