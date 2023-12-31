package repository

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/entity"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(ctx context.Context, userRequest *entity.UserRequest) error {
	if err := r.db.Table(entity.TableName()).WithContext(ctx).Create(userRequest).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Table(entity.TableName()).WithContext(ctx).Where("email = ?", email).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrUserNotFoundWithEmail
		}
		return nil, err
	}
	return &user, nil
}

func (r *authRepo) GetUserByName(ctx context.Context, name string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Table(entity.TableName()).WithContext(ctx).Where("username = ?", name).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrUserNotFoundWithName
		}
		return nil, err
	}
	return &user, nil
}
