package usecase

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user"
)

type commentUC struct {
	commentRepo comment.Repository
	userRepo    user.Repository
}

func NewCommentUsecase(commentRepo comment.Repository, userRepo user.Repository) comment.Usecase {
	return &commentUC{commentRepo: commentRepo, userRepo: userRepo}
}

func (c *commentUC) CommentPost(ctx context.Context, commentReq *entity.CommentReqest) error {
	return c.commentRepo.CreateComment(ctx, commentReq)
}

func (c *commentUC) DeleteComment(ctx context.Context, commentDeleteReq *entity.CommentDeleteReqest) error {
	getUser, err := c.userRepo.GetUserByID(ctx, commentDeleteReq.UserID)
	if err != nil {
		return errors.New("user_id not found")
	}

	getComment, err := c.commentRepo.GetCommentByID(ctx, commentDeleteReq.CommentID)
	if err != nil {
		return errors.New("comment_id not found")
	}

	if getUser.ID != getComment.UserID {
		return errors.New("no permission, only owner of this like can do this")
	}

	return c.commentRepo.DeleteComment(ctx, commentDeleteReq.CommentID)
}
