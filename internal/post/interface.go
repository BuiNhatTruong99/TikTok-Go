package post

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post/entity"
)

type Repository interface {
	CreatePost(ctx context.Context, request *entity.PostRequest) error
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	GetPostByID(ctx context.Context, postID int64) (*entity.Post, error)
	GetPostByUserID(ctx context.Context, userID int64) ([]entity.Post, error)
	DeletePostByID(ctx context.Context, postID int64) error
}

type UseCase interface {
	CreatePost(ctx context.Context, request *entity.PostRequest) error
	DeletePost(ctx context.Context, postID int64) error
	GetPostByID(ctx context.Context, postID int64) (*entity.Post, error)
	GetPostsByUserID(ctx context.Context, userID int64) ([]entity.Post, error)
	GetAllPost(ctx context.Context) ([]entity.Post, error)
}
