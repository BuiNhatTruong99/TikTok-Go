package jwt

import (
	"errors"
	"fmt"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    int64     `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewPayload(userID int64, expiresAt time.Duration) (*Payload, error) {
	claimsID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:        claimsID,
		UserID:    userID,
		ExpiresAt: time.Now().Add(expiresAt),
	}, nil
}

func GenerateToken(userID int64, duration time.Duration, config *config.Config) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", payload, err
	}

	claims := jwt.MapClaims{
		"payload_id": payload.ID,
		"user_id":    payload.UserID,
		"expiresAt":  time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Server.JwtSecretKey))
	if err != nil {
		return "", payload, fmt.Errorf("error when Singed token: %v", err)
	}

	return tokenString, payload, nil
}

func VerifyToken(tokenString string, config *config.Config) (*Payload, error) {
	mapClaims, err := extractClaims(tokenString, config)
	if err != nil {
		return nil, err
	}

	exp, ok := mapClaims["expiresAt"].(float64)
	if !ok {
		return nil, errors.New("expiration claim not found in token")
	}

	expirationTime := time.Unix(int64(exp), 0)

	if time.Now().After(expirationTime) {
		return nil, errors.New("token has expired")
	}

	payloadID, err := uuid.Parse(mapClaims["payload_id"].(string))
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        payloadID,
		UserID:    int64(mapClaims["user_id"].(float64)),
		ExpiresAt: time.Unix(int64(exp), 0),
	}

	return payload, nil
}

func extractClaims(tokenString string, config *config.Config) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Server.JwtSecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error extracting claims from token")
	}

	return claims, nil
}
