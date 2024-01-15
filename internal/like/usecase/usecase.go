package usecase

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/like"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/like/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user"
)

type likeUC struct {
	likeRepo like.Repository
	userRepo user.Repository
}

func NewLikeUsecase(likeRepo like.Repository, userRepo user.Repository) like.Usecase {
	return &likeUC{likeRepo: likeRepo, userRepo: userRepo}
}

func (l *likeUC) LikePost(ctx context.Context, likeRequest *entity.LikeRequest) (*entity.LikeResponse, error) {
	getLike, err := l.likeRepo.GetLikeByUserAndPost(ctx, likeRequest)
	if getLike != nil && err == nil {
		return nil, errors.New("this user already like this post")
	}
	return l.likeRepo.CreateLike(ctx, likeRequest)
}

func (l *likeUC) UndoLikePost(ctx context.Context, likeRequest *entity.LikeDeleteRequest) error {
	getLike, err := l.likeRepo.GetLikeByID(ctx, likeRequest.LikeID)
	if err != nil {
		return err
	}
	getUser, err := l.userRepo.GetUserByID(ctx, likeRequest.UserID)
	if err != nil {
		return err
	}
	if getLike.UserID != getUser.ID {
		return errors.New("no permission, only owner of this like can do this")
	}

	return l.likeRepo.DeleteLike(ctx, getLike.ID)
}
