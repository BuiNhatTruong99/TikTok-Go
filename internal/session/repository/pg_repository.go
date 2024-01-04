package repository

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/session"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/session/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) session.Repository {
	return &sessionRepo{db: db}
}

func (r *sessionRepo) CreateSession(ctx context.Context, session *entity.Session) (*entity.Session, error) {
	if err := r.db.Table(entity.TableName()).WithContext(ctx).Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (r *sessionRepo) GetSession(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
	var getSession entity.Session

	if err := r.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?", id).First(&getSession).Error; err != nil {

		return nil, err
	}
	return &getSession, nil
}
