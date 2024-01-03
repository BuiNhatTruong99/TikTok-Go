package repository

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/entity"
	"gorm.io/gorm"
	"strings"
)

const (
	uniqueConstraintUsername = "users_username_key"
	uniqueConstraintEmail    = "users_email_key"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepo{db: db}
}

func handleDuplicateKeyViolationError(err error) error {
	if strings.Contains(err.Error(), uniqueConstraintUsername) {
		return errors.New("username already exists")
	}
	if strings.Contains(err.Error(), uniqueConstraintEmail) {
		return errors.New("email already exists")
	}
	return err
}

func (r *authRepo) Register(ctx context.Context, userRequest *entity.UserRequest) error {
	if err := r.db.Table(entity.TableName()).WithContext(ctx).Create(userRequest).
		Error; err != nil {
		return handleDuplicateKeyViolationError(err)
	}
	return nil
}

func (r *authRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Table(entity.TableName()).WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *authRepo) GetUserByName(ctx context.Context, name string) (*entity.User, error) {
	var user entity.User

	if err := r.db.Table(entity.TableName()).WithContext(ctx).Where("username = ?", name).First(&user).Error; err != nil {

		return nil, err
	}
	return &user, nil
}
