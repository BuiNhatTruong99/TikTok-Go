package repository

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user.Repository {
	return &userRepo{db: db}
}

func (u *userRepo) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	var user entity.User
	if err := u.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) UpdateProfileUser(ctx context.Context, userID int64,
	request *entity.ProfileRequest) error {
	if err := u.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?", userID).Updates(request).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepo) UpdateAvatarUser(ctx context.Context, userID int64, request *entity.AvatarRequest) error {
	if err := u.db.Table(entity.TableName()).WithContext(ctx).Where("id = ?",
		userID).Updates(request).Error; err != nil {
		return err
	}
	return nil
}
