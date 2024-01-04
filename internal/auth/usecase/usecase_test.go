package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/config"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/entity"
	mock_auth "github.com/BuiNhatTruong99/TikTok-Go/internal/auth/mock"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuthUC_Register(t *testing.T) {

	testCase := []struct {
		name       string
		user       *entity.UserRequest
		buildStubs func(repo *mock_auth.MockRepository)
		err        error
	}{
		{
			name: "OK",
			user: &entity.UserRequest{
				Username:     utils.RandomString(6),
				Email:        utils.RandomEmail(),
				HashPassword: utils.RandomString(6),
				AvatarUrl:    "default_url",
				Bio:          "this is my bio",
			},
			buildStubs: func(repo *mock_auth.MockRepository) {
				repo.EXPECT().Register(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name: "InternalError",
			user: &entity.UserRequest{
				Username:     utils.RandomString(6),
				Email:        utils.RandomEmail(),
				HashPassword: utils.RandomString(6),
				AvatarUrl:    "default_url",
				Bio:          "this is my bio",
			},
			buildStubs: func(repo *mock_auth.MockRepository) {
				repo.EXPECT().Register(gomock.Any(), gomock.Any()).Times(1).Return(sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			cfg := &config.Config{
				Server: config.ServerConfig{
					JwtSecretKey:         "secret",
					AccessTokenDuration:  15,
					RefreshTokenDuration: 24,
				},
			}

			mockAuthRepo := mock_auth.NewMockRepository(ctrl)
			authUseCase := NewAuthUsecase(mockAuthRepo, cfg)
			tc.buildStubs(mockAuthRepo)
			err := authUseCase.Register(context.TODO(), tc.user)
			require.Equal(t, tc.err, err)
		})
	}

}

func TestAuthUC_Login(t *testing.T) {
	passWord := utils.RandomString(6)
	hashPassword, err := entity.HashPassword(passWord)
	require.NoError(t, err)
	user := &entity.User{
		ID:           utils.RandomInt(1, 5),
		Username:     utils.RandomString(6),
		Email:        utils.RandomEmail(),
		HashPassword: hashPassword,
		AvatarUrl:    utils.RandomString(6),
		Bio:          utils.RandomString(6),
		CreatedAt:    nil,
	}

	testCase := []struct {
		name       string
		user       *entity.UserLogin
		buildStubs func(repo *mock_auth.MockRepository)
		err        error
	}{
		{
			name: "OK",
			user: &entity.UserLogin{
				Email:        user.Email,
				HashPassword: passWord,
			},
			buildStubs: func(repo *mock_auth.MockRepository) {
				repo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).Return(user, nil)
			},
			err: nil,
		},
		{
			name: "InternalError",
			user: &entity.UserLogin{
				Email:        user.Email,
				HashPassword: passWord,
			},
			buildStubs: func(repo *mock_auth.MockRepository) {
				repo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil,
					sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
		{
			name: "UserNotFound",
			user: &entity.UserLogin{
				Email:        "NotFound@gmail.com",
				HashPassword: passWord,
			},
			buildStubs: func(repo *mock_auth.MockRepository) {
				repo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil,
					errors.New("email or password incorrect"))
			},
			err: errors.New("email or password incorrect"),
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			cfg := &config.Config{
				Server: config.ServerConfig{
					JwtSecretKey:         utils.RandomString(32),
					AccessTokenDuration:  15,
					RefreshTokenDuration: 24,
				},
			}

			mockAuthRepo := mock_auth.NewMockRepository(ctrl)
			authUseCase := NewAuthUsecase(mockAuthRepo, cfg)
			tc.buildStubs(mockAuthRepo)
			response, err := authUseCase.Login(context.TODO(), tc.user)
			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.NotNil(t, response)
			}
		})
	}
}
