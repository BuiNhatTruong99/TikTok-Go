package repository

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/like"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/like/entity"
	"gorm.io/gorm"
)

type likeRepo struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) like.Repository {
	return &likeRepo{db: db}
}

func (l *likeRepo) CreateLike(ctx context.Context, likeRequest *entity.LikeRequest) (*entity.LikeResponse, error) {
	if err := l.db.Table(entity.TableName()).WithContext(ctx).Create(&likeRequest).Error; err != nil {
		return nil, err
	}
	return &entity.LikeResponse{
		UserID: likeRequest.UserID,
		PostID: likeRequest.PostID,
	}, nil
}

func (l *likeRepo) DeleteLike(ctx context.Context, likeID int64) error {
	if err := l.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?", likeID).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (l *likeRepo) GetLikeByUserAndPost(ctx context.Context, likeRequest *entity.LikeRequest) (*entity.Like, error) {
	var getLike entity.Like
	if err := l.db.Table(entity.TableName()).WithContext(ctx).
		Where("user_id = ? AND post_id = ?", likeRequest.UserID, likeRequest.PostID).First(&getLike).Error; err != nil {
		return nil, err
	}
	return &getLike, nil
}

func (l *likeRepo) GetLikeByID(ctx context.Context, likeID int64) (*entity.Like, error) {
	var getLike entity.Like
	if err := l.db.Table(entity.TableName()).WithContext(ctx).
		Where("id = ?", likeID).First(&getLike).Error; err != nil {
		return nil, errors.New("like record not found")
	}
	return &getLike, nil
}
