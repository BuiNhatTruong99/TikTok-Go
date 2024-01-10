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

func TestPostUC_GetPostByID(t *testing.T) {
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
				repo.EXPECT().GetPostByID(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(post, nil)
			},
			err: nil,
		},
		{
			name:   "InternalError",
			postID: post.ID,
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().GetPostByID(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
		{
			name:   "PostNotFound",
			postID: post.ID,
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().GetPostByID(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(nil, sql.ErrNoRows)
			},
			err: errors.New("sql: no rows in result set"),
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
			postResult, err := postUC.GetPostByID(context.Background(), tc.postID)
			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.NotNil(t, postResult)
				require.Equal(t, postResult.UserID, post.UserID)
			}
		})
	}
}

func TestPostUC_GetPostsByUserID(t *testing.T) {
	userID := utils.RandomInt(1, 5)
	postsList := []entity.Post{}
	for i := 0; i < 5; i++ {
		post := entity.Post{
			ID:        utils.RandomInt(1, 5),
			UserID:    userID,
			VideoUrl:  utils.RandomString(6),
			Caption:   utils.RandomString(20),
			CreatedAt: nil,
		}
		postsList = append(postsList, post)
	}
	testCase := []struct {
		name       string
		userID     int64
		buildStubs func(repo *mock_post.MockRepository)
		err        error
	}{
		{
			name:   "OK",
			userID: userID,
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().GetPostByUserID(gomock.Any(), gomock.Eq(userID)).Times(1).Return(postsList, nil)
			},
			err: nil,
		},
		{
			name:   "InternalError",
			userID: userID,
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().GetPostByUserID(gomock.Any(), gomock.Eq(userID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
		{
			name:   "PostNotFound",
			userID: userID,
			buildStubs: func(repo *mock_post.MockRepository) {
				repo.EXPECT().GetPostByUserID(gomock.Any(), gomock.Eq(userID)).Times(1).Return(nil, sql.ErrNoRows)
			},
			err: errors.New("sql: no rows in result set"),
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
			postResult, err := postUC.GetPostsByUserID(context.Background(), tc.userID)
			require.Equal(t, tc.err, err)

			if tc.err == nil {
				require.NotNil(t, postResult)
				for i := 0; i < len(postResult); i++ {
					require.Equal(t, postResult[i].UserID, userID)
				}
			}
		})
	}
}
