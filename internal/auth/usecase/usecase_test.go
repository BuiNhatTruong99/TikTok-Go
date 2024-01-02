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
				repo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
				repo.EXPECT().GetUserByName(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
				repo.EXPECT().Register(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name: "ExistingEmail",
			user: &entity.UserRequest{
				Username:     utils.RandomString(6),
				Email:        utils.RandomEmail(),
				HashPassword: utils.RandomString(6),
				AvatarUrl:    "default_url",
				Bio:          "this is my bio",
			},
			buildStubs: func(repo *mock_auth.MockRepository) {
				repo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(&entity.User{}, nil)
			},
			err: errors.New("email already exists"),
		},
		{
			name: "ExistingUsername",
			user: &entity.UserRequest{
				Username:     utils.RandomString(6),
				Email:        utils.RandomEmail(),
				HashPassword: utils.RandomString(6),
				AvatarUrl:    "default_url",
				Bio:          "this is my bio",
			},
			buildStubs: func(repo *mock_auth.MockRepository) {
				repo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
				repo.EXPECT().GetUserByName(gomock.Any(), gomock.Any()).Return(&entity.User{}, nil)
			},
			err: errors.New("username already exists"),
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
				repo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
				repo.EXPECT().GetUserByName(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows)
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
