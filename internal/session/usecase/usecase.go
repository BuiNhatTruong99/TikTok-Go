package usecase

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/session"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/session/entity"
	"github.com/google/uuid"
	"time"
)

type sessionUC struct {
	sessionRepo session.Repository
}

func NewSessionUsecase(repository session.Repository) session.Usecase {
	return &sessionUC{sessionRepo: repository}
}

func (s *sessionUC) CreateSession(ctx context.Context, sessionID uuid.UUID, userID int64, refreshToken, clientIP string,
	expirestAt time.Time) (*entity.Session, error) {
	newSession := &entity.Session{
		ID:           sessionID,
		UserID:       userID,
		RefreshToken: refreshToken,
		ClientIP:     clientIP,
		IsBlocked:    false,
		ExpiresAt:    expirestAt,
	}
	return s.sessionRepo.CreateSession(ctx, newSession)
}

func (s *sessionUC) GetSession(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
	return s.sessionRepo.GetSession(ctx, id)
}
