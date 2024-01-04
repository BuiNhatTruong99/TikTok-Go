package entity

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           uuid.UUID  `json:"id" gorm:"column:id;primaryKey"`
	UserID       int64      `json:"user_id" gorm:"column:user_id"`
	RefreshToken string     `json:"refresh_token" gorm:"column:refresh_token"`
	ClientIP     string     `json:"client_ip" gorm:"column:client_ip"`
	IsBlocked    bool       `json:"is_blocked" gorm:"column:is_blocked"`
	ExpiresAt    time.Time  `json:"expires_at" gorm:"column:expires_at"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at"`
}

type SessionRequest struct {
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ClientIP     string    `json:"client_ip"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type NewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type NewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiredAt time.Time `json:"access_token_expired_at"`
}

func TableName() string {
	return "sessions"
}
