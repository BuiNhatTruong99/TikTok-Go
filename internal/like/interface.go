package like

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/like/entity"
)

type Repository interface {
	CreateLike(ctx context.Context, likeRequest *entity.LikeRequest) (*entity.LikeResponse, error)
	DeleteLike(ctx context.Context, likeID int64) error
	GetLikeByUserAndPost(ctx context.Context, like *entity.LikeRequest) (*entity.Like, error)
	GetLikeByID(ctx context.Context, likeID int64) (*entity.Like, error)
}

type Usecase interface {
	LikePost(ctx context.Context, likeRequest *entity.LikeRequest) (*entity.LikeResponse, error)
	UndoLikePost(ctx context.Context, likeRequest *entity.LikeDeleteRequest) error
}
