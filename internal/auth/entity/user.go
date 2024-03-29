package entity

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	SessionID             uuid.UUID `json:"session_id"`
	User                  *User     `json:"user"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiredAt  time.Time `json:"access_token_expired_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time `json:"refresh_token_expired_at"`
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

func (u *UserRequest) Validate() error {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	u.HashPassword = strings.TrimSpace(u.HashPassword)

	if err := validateName(u.Username); err != nil {
		return err
	}

	if err := validatePassword(u.HashPassword); err != nil {
		return err
	}

	if err := validateEmail(u.Email); err != nil {
		return err
	}
	return nil
}

func (u *UserLogin) Validate() error {
	u.Email = strings.TrimSpace(u.Email)
	u.HashPassword = strings.TrimSpace(u.HashPassword)

	if err := validatePassword(u.HashPassword); err != nil {
		return err
	}

	if err := validateEmail(u.Email); err != nil {
		return err
	}
	return nil
}
