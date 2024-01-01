package auth

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/entity"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByName(ctx context.Context, name string) (*entity.User, error)
	Register(ctx context.Context, user *entity.UserRequest) error
}

type Usecase interface {
	Register(ctx context.Context, user *entity.UserRequest) error
	Login(ctx context.Context, data *entity.UserLogin, config *config.Config) (*entity.LoginResponse, error)
}
