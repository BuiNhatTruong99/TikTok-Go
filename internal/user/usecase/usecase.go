package usecase

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"
)

type userUC struct {
	userRepo user.Repository
}

func NewUserUseCase(userRepo user.Repository) user.UseCase {
	return &userUC{userRepo: userRepo}
}

func (u *userUC) ChangeAvatar(ctx context.Context, userID int64, request *entity.AvatarRequest) error {
	_, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	return u.userRepo.UpdateAvatarUser(ctx, userID, request)
}

func (u *userUC) UpdateProfile(ctx context.Context, userID int64, request *entity.ProfileRequest) error {
	if err := request.Validate(); err != nil {
		return err
	}
	_, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	return u.userRepo.UpdateProfileUser(ctx, userID, request)
}
