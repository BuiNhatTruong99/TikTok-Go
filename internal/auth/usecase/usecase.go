package usecase

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/jwt"
	"gorm.io/gorm"
)

type authUC struct {
	auRepository auth.Repository
	config       *config.Config
}

func NewAuthUsecase(auRepository auth.Repository, config *config.Config) auth.Usecase {
	return &authUC{auRepository: auRepository, config: config}
}

func (u *authUC) Register(ctx context.Context, user *entity.UserRequest) error {
	if err := user.Validate(); err != nil {
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email or password incorrect")
		}
		return nil, err
	}

	if err := entity.ComparePassword(data.HashPassword, findUser.HashPassword); err != nil {
		return nil, errors.New("email or password incorrect")
	}

	accessToken, accessPayload, err := jwt.GenerateToken(findUser.ID, u.config.Server.AccessTokenDuration, u.config)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := jwt.GenerateToken(findUser.ID, u.config.Server.RefreshTokenDuration, u.config)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		SessionID:             refreshPayload.ID,
		User:                  findUser,
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  accessPayload.ExpiresAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: refreshPayload.ExpiresAt,
	}, nil
}
