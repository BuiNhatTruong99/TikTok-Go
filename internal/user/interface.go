package user

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"
)

type Repository interface {
	GetUserByID(ctx context.Context, userID int64) (*entity.User, error)
	UpdateProfileUser(ctx context.Context, userID int64, request *entity.ProfileRequest) error
	UpdateAvatarUser(ctx context.Context, userID int64, request *entity.AvatarRequest) error
}

type UseCase interface {
	ChangeAvatar(ctx context.Context, userID int64, request *entity.AvatarRequest) error
	UpdateProfile(ctx context.Context, userID int64, request *entity.ProfileRequest) error
}
