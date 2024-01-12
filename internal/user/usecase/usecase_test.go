package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"
	mock_user "github.com/BuiNhatTruong99/TikTok-Go/internal/user/mock"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserUC_ChangeAvatar(t *testing.T) {
	user := &entity.User{
		ID:           utils.RandomInt(1, 5),
		Username:     utils.RandomString(6),
		Email:        utils.RandomEmail(),
		HashPassword: utils.RandomString(6),
		AvatarUrl:    utils.RandomString(6),
		Bio:          utils.RandomString(6),
		CreatedAt:    nil,
	}
	testCase := []struct {
		name       string
		param      *entity.AvatarRequest
		buildStubs func(repo *mock_user.MockRepository)
		err        error
	}{
		{
			name:  "OK",
			param: &entity.AvatarRequest{AvatarUrl: utils.RandomString(50)},
			buildStubs: func(repo *mock_user.MockRepository) {
				repo.EXPECT().GetUserByID(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)
				repo.EXPECT().UpdateAvatarUser(gomock.Any(), gomock.Eq(user.ID), gomock.Any()).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name:  "InternalError",
			param: &entity.AvatarRequest{AvatarUrl: utils.RandomString(50)},
			buildStubs: func(repo *mock_user.MockRepository) {
				repo.EXPECT().GetUserByID(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
		{
			name:  "NotFound",
			param: &entity.AvatarRequest{AvatarUrl: utils.RandomString(50)},
			buildStubs: func(repo *mock_user.MockRepository) {
				repo.EXPECT().GetUserByID(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(nil, sql.ErrNoRows)
			},
			err: errors.New("sql: no rows in result set"),
		},
	}
	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mock_user.NewMockRepository(ctrl)
			userUsecase := NewUserUseCase(mockUserRepo)
			tc.buildStubs(mockUserRepo)
			err := userUsecase.ChangeAvatar(context.Background(), user.ID, tc.param)
			require.Equal(t, err, tc.err)
		})
	}

}

func TestUserUC_UpdateProfile(t *testing.T) {
	user := &entity.User{
		ID:           utils.RandomInt(1, 5),
		Username:     utils.RandomString(6),
		Email:        utils.RandomEmail(),
		HashPassword: utils.RandomString(6),
		AvatarUrl:    utils.RandomString(6),
		Bio:          utils.RandomString(6),
		CreatedAt:    nil,
	}
	testCase := []struct {
		name       string
		param      *entity.ProfileRequest
		buildStubs func(repo *mock_user.MockRepository)
		err        error
	}{
		{
			name: "OK",
			param: &entity.ProfileRequest{
				Username: utils.RandomString(6),
				Bio:      utils.RandomString(6),
			},
			buildStubs: func(repo *mock_user.MockRepository) {
				repo.EXPECT().GetUserByID(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)
				repo.EXPECT().UpdateProfileUser(gomock.Any(), gomock.Eq(user.ID), gomock.Any()).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name: "InternalError",
			param: &entity.ProfileRequest{
				Username: utils.RandomString(6),
				Bio:      utils.RandomString(6),
			},
			buildStubs: func(repo *mock_user.MockRepository) {
				repo.EXPECT().GetUserByID(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
		{
			name: "NotFound",
			param: &entity.ProfileRequest{
				Username: utils.RandomString(6),
				Bio:      utils.RandomString(6),
			},
			buildStubs: func(repo *mock_user.MockRepository) {
				repo.EXPECT().GetUserByID(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(nil, sql.ErrNoRows)
			},
			err: errors.New("sql: no rows in result set"),
		},
	}
	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mock_user.NewMockRepository(ctrl)
			userUsecase := NewUserUseCase(mockUserRepo)
			tc.buildStubs(mockUserRepo)
			err := userUsecase.UpdateProfile(context.Background(), user.ID, tc.param)
			require.Equal(t, err, tc.err)
		})
	}

}
