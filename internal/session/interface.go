package session

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/session/entity"
	"github.com/google/uuid"
	"time"
)

type Repository interface {
	CreateSession(ctx context.Context, session *entity.Session) (*entity.Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (*entity.Session, error)
}

type Usecase interface {
	CreateSession(ctx context.Context, sessionID uuid.UUID, userID int64, refreshToken, clientIP string,
		expirestAt time.Time) (*entity.Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (*entity.Session, error)
}
