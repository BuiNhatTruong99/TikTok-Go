package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post/entity"
	mock_post "github.com/BuiNhatTruong99/TikTok-Go/internal/post/mock"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostUC_CreatePost(t *testing.T) {
	testCase := []struct {
		name       string
		post       *entity.PostRequest
		buildStubs func(repo *mock_post.MockRepository)
		err        error
	}{
		{
			name: "OK",
			post: &entity.PostRequest{
				UserID:   utils.RandomInt(1, 5),
				VideoUrl: utils.RandomString(6),
				Caption:  utils.RandomString(20),
			},
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name: "InternalError",
			post: &entity.PostRequest{
				UserID:   utils.RandomInt(1, 5),
				VideoUrl: utils.RandomString(6),
				Caption:  utils.RandomString(20),
			},
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Times(1).Return(sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPostRepo := mock_post.NewMockRepository(ctrl)
			postUC := NewPostUsecase(mockPostRepo)
			tc.buildStubs(mockPostRepo)
			err := postUC.CreatePost(context.Background(), tc.post)
			require.Equal(t, tc.err, err)
		})
	}
}

func TestPostUC_DeletePost(t *testing.T) {
	post := &entity.Post{
		ID:        utils.RandomInt(1, 5),
		UserID:    utils.RandomInt(1, 5),
		VideoUrl:  utils.RandomString(6),
		Caption:   utils.RandomString(20),
		CreatedAt: nil,
	}
	testCase := []struct {
		name       string
		postID     int64
		buildStubs func(repo *mock_post.MockRepository)
		err        error
	}{
		{
			name:   "OK",
			postID: post.ID,
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().DeletePostByID(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name:   "InternalError",
			postID: post.ID,
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().DeletePostByID(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
	}
	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPostRepo := mock_post.NewMockRepository(ctrl)
			postUC := NewPostUsecase(mockPostRepo)
			tc.buildStubs(mockPostRepo)
			err := postUC.DeletePost(context.Background(), tc.postID)
			require.Equal(t, tc.err, err)
		})
	}
}
