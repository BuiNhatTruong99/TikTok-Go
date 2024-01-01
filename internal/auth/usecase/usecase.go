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
}

func NewAuthUsecase(auRepository auth.Repository) auth.Usecase {
	return &authUC{auRepository: auRepository}
}

func (u *authUC) Register(ctx context.Context, user *entity.UserRequest) error {
	existsUserWithEmail, err := u.auRepository.GetUserByEmail(ctx, user.Email)
	if existsUserWithEmail != nil || err == nil {
		return errors.New("email already exists")
	}

	existsUserWithUsername, err := u.auRepository.GetUserByEmail(ctx, user.Username)
	if existsUserWithUsername != nil || err == nil {
		return errors.New("username already exists")
	}

	hashPassword, err := entity.HashPassword(user.HashPassword)
	if err != nil {
		return err
	}
	user.HashPassword = hashPassword

	return u.auRepository.Register(ctx, user)
}

func (u *authUC) Login(ctx context.Context, data *entity.UserLogin, config *config.Config) (*entity.LoginResponse, error) {
	findUser, err := u.auRepository.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if err := entity.ComparePassword(data.HashPassword, findUser.HashPassword); err != nil {
		return nil, errors.New("email or Password incorrect")
	}

	accessToken, err := jwt.GenerateToken(findUser.Email, config.Server.AccessTokenDuration, config)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateToken(findUser.Email, config.Server.RefreshTokenDuration, config)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		User:         findUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
