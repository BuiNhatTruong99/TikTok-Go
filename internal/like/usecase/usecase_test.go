package usecase

import (
	"context"
	"database/sql"
	"errors"
	userEntity "github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"

	"github.com/BuiNhatTruong99/TikTok-Go/internal/like/entity"
	mock_like "github.com/BuiNhatTruong99/TikTok-Go/internal/like/mock"
	mock_user "github.com/BuiNhatTruong99/TikTok-Go/internal/user/mock"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLikeUC_LikePost(t *testing.T) {
	userID := utils.RandomInt(1, 5)
	postID := utils.RandomInt(1, 5)

	testCase := []struct {
		name       string
		like       *entity.LikeRequest
		buildStubs func(repo *mock_like.MockRepository)
		err        error
	}{
		{
			name: "OK",
			like: &entity.LikeRequest{
				UserID: userID,
				PostID: postID,
			},
			buildStubs: func(repo *mock_like.MockRepository) {
				repo.EXPECT().GetLikeByUserAndPost(gomock.Any(), gomock.Any()).Times(1).Return(nil, sql.ErrNoRows)
				repo.EXPECT().CreateLike(gomock.Any(), gomock.Any()).Times(1).Return(&entity.LikeResponse{
					UserID: userID,
					PostID: postID,
				}, nil)
			},
			err: nil,
		},
		{
			name: "InternalError",
			like: &entity.LikeRequest{
				UserID: userID,
				PostID: postID,
			},
			buildStubs: func(repo *mock_like.MockRepository) {
				repo.EXPECT().GetLikeByUserAndPost(gomock.Any(), gomock.Any()).Times(1).Return(nil, sql.ErrNoRows)
				repo.EXPECT().CreateLike(gomock.Any(), gomock.Any()).Times(1).Return(&entity.LikeResponse{}, sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
		{
			name: "AlreadyLike",
			like: &entity.LikeRequest{
				UserID: userID,
				PostID: postID,
			},
			buildStubs: func(repo *mock_like.MockRepository) {
				repo.EXPECT().GetLikeByUserAndPost(gomock.Any(), gomock.Any()).Times(1).Return(&entity.Like{}, nil)
			},
			err: errors.New("this user already like this post"),
		},
		{
			name: "AlreadyLike",
			like: &entity.LikeRequest{
				UserID: userID,
				PostID: postID,
			},
			buildStubs: func(repo *mock_like.MockRepository) {
				repo.EXPECT().GetLikeByUserAndPost(gomock.Any(), gomock.Any()).Times(1).Return(&entity.Like{}, nil)
			},
			err: errors.New("this user already like this post"),
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLikeRepo := mock_like.NewMockRepository(ctrl)
			mockUserRepo := mock_user.NewMockRepository(ctrl)
			likeUsecase := NewLikeUsecase(mockLikeRepo, mockUserRepo)
			tc.buildStubs(mockLikeRepo)
			like, err := likeUsecase.LikePost(context.Background(), tc.like)
			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.NotNil(t, like)
			}
		})
	}
}

func TestLikeUC_UndoLikePost(t *testing.T) {
	userID := utils.RandomInt(1, 5)
	LikeID := utils.RandomInt(1, 5)

	testCase := []struct {
		name       string
		like       *entity.LikeDeleteRequest
		buildStubs func(repo *mock_like.MockRepository, userRepo *mock_user.MockRepository)
		err        error
	}{
		{
			name: "OK",
			like: &entity.LikeDeleteRequest{
				UserID: userID,
				LikeID: LikeID,
			},
			buildStubs: func(repo *mock_like.MockRepository, userRepo *mock_user.MockRepository) {
				repo.EXPECT().GetLikeByID(gomock.Any(), gomock.Any()).Times(1).Return(&entity.Like{}, nil)
				userRepo.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Times(1).Return(&userEntity.User{}, nil)
				repo.EXPECT().DeleteLike(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name: "InternalError",
			like: &entity.LikeDeleteRequest{
				UserID: userID,
				LikeID: LikeID,
			},
			buildStubs: func(repo *mock_like.MockRepository, userRepo *mock_user.MockRepository) {
				repo.EXPECT().GetLikeByID(gomock.Any(), gomock.Any()).Times(1).Return(nil, sql.ErrNoRows)
			},
			err: errors.New("sql: no rows in result set"),
		},
		{
			name: "NotOwner",
			like: &entity.LikeDeleteRequest{
				UserID: userID,
				LikeID: LikeID,
			},
			buildStubs: func(repo *mock_like.MockRepository, userRepo *mock_user.MockRepository) {
				repo.EXPECT().GetLikeByID(gomock.Any(), gomock.Any()).Times(1).Return(&entity.Like{}, nil)
				userRepo.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Times(1).Return(&userEntity.User{ID: 999},
					nil)
			},
			err: errors.New("no permission, only owner of this like can do this"),
		},
	}
	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLikeRepo := mock_like.NewMockRepository(ctrl)
			mockUserRepo := mock_user.NewMockRepository(ctrl)
			likeUsecase := NewLikeUsecase(mockLikeRepo, mockUserRepo)
			tc.buildStubs(mockLikeRepo, mockUserRepo)
			err := likeUsecase.UndoLikePost(context.Background(), tc.like)
			require.Equal(t, tc.err, err)

		})
	}
}
