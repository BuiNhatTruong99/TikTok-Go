package comment

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment/entity"
)

type Repository interface {
	CreateComment(ctx context.Context, commentReq *entity.CommentReqest) error
	DeleteComment(ctx context.Context, commentID int64) error
	GetCommentByID(ctx context.Context, commentID int64) (*entity.Comment, error)
}

type Usecase interface {
	CommentPost(ctx context.Context, commentReq *entity.CommentReqest) error
	DeleteComment(ctx context.Context, commentDeleteReq *entity.CommentDeleteReqest) error
}
