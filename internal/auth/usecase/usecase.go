package usecase

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/jwt"
)

type authUC struct {
	auRepository auth.Repository
	config       *config.Config
}

func NewAuthUsecase(auRepository auth.Repository, config *config.Config) auth.Usecase {
	return &authUC{auRepository: auRepository, config: config}
}

func (u *authUC) checkExistingUser(ctx context.Context, email, username string) error {
	if _, err := u.auRepository.GetUserByEmail(ctx, email); err == nil {
		return errors.New("email already exists")
	}

	if _, err := u.auRepository.GetUserByName(ctx, username); err == nil {
		return errors.New("username already exists")
	}

	return nil
}

func (u *authUC) Register(ctx context.Context, user *entity.UserRequest) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := u.checkExistingUser(ctx, user.Email, user.Username); err != nil {
		return err
	}

	hashPassword, err := entity.HashPassword(user.HashPassword)
	if err != nil {
		return err
	}
	user.HashPassword = hashPassword

	return u.auRepository.Register(ctx, user)
}

func (u *authUC) Login(ctx context.Context, data *entity.UserLogin) (*entity.LoginResponse, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	findUser, err := u.auRepository.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if err := entity.ComparePassword(data.HashPassword, findUser.HashPassword); err != nil {
		return nil, errors.New("email or Password incorrect")
	}

	accessToken, err := jwt.GenerateToken(findUser.Email, u.config.Server.AccessTokenDuration, u.config)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateToken(findUser.Email, u.config.Server.RefreshTokenDuration, u.config)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		User:         findUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
