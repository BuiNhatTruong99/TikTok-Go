package jwt

import (
	"errors"
	"fmt"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(email string, duration time.Duration, config *config.Config) (string, error) {
	claims := jwt.MapClaims{
		"email":     email,
		"expiresAt": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Server.JwtSecretKey))
	if err != nil {
		return "", fmt.Errorf("error when Singed token: %v", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, config *config.Config) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.Server.JwtSecretKey, nil
	})
	if err != nil {
		return fmt.Errorf("error pasre token: %v", err)
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
