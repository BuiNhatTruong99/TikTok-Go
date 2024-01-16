package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment/entity"
	mock_comment "github.com/BuiNhatTruong99/TikTok-Go/internal/comment/mock"
	userEntity "github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"
	mock_user "github.com/BuiNhatTruong99/TikTok-Go/internal/user/mock"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommentUC_CommentPost(t *testing.T) {
	userID := utils.RandomInt(1, 5)
	postID := utils.RandomInt(1, 5)
	text := utils.RandomString(15)

	testCase := []struct {
		name       string
		comment    *entity.CommentReqest
		buildStubs func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository)
		err        error
	}{
		{
			name: "OK",
			comment: &entity.CommentReqest{
				UserID: userID,
				PostID: postID,
				Text:   text,
			},
			buildStubs: func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository) {
				repo.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name: "InternalError",
			comment: &entity.CommentReqest{
				UserID: userID,
				PostID: postID,
				Text:   text,
			},
			buildStubs: func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository) {
				repo.EXPECT().CreateComment(gomock.Any(), gomock.Any()).Times(1).Return(sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCommentRepo := mock_comment.NewMockRepository(ctrl)
			mockUserRepo := mock_user.NewMockRepository(ctrl)

			commentUsecase := NewCommentUsecase(mockCommentRepo, mockUserRepo)
			tc.buildStubs(mockCommentRepo, mockUserRepo)

			err := commentUsecase.CommentPost(context.Background(), tc.comment)
			require.Equal(t, err, tc.err)
		})
	}
}

func TestCommentUC_DeleteComment(t *testing.T) {
	userID := utils.RandomInt(1, 5)
	commentID := utils.RandomInt(1, 5)

	testCase := []struct {
		name       string
		comment    *entity.CommentDeleteReqest
		buildStubs func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository)
		err        error
	}{
		{
			name: "OK",
			comment: &entity.CommentDeleteReqest{
				CommentID: commentID,
				UserID:    userID,
			},
			buildStubs: func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository) {
				userRepo.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Times(1).Return(&userEntity.User{}, nil)
				repo.EXPECT().GetCommentByID(gomock.Any(), gomock.Any()).Times(1).Return(&entity.Comment{}, nil)
				repo.EXPECT().DeleteComment(gomock.Any(), gomock.Eq(commentID)).Times(1).Return(nil)
			},
			err: nil,
		},
		{
			name: "InternalError",
			comment: &entity.CommentDeleteReqest{
				CommentID: commentID,
				UserID:    userID,
			},
			buildStubs: func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository) {
				userRepo.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Times(1).Return(&userEntity.User{}, nil)
				repo.EXPECT().GetCommentByID(gomock.Any(), gomock.Any()).Times(1).Return(&entity.Comment{}, nil)
				repo.EXPECT().DeleteComment(gomock.Any(), gomock.Eq(commentID)).Times(1).Return(sql.ErrConnDone)
			},
			err: errors.New("sql: connection is already closed"),
		},
		{
			name: "NotOwner",
			comment: &entity.CommentDeleteReqest{
				CommentID: commentID,
				UserID:    userID,
			},
			buildStubs: func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository) {
				userRepo.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Times(1).Return(&userEntity.User{}, nil)
				repo.EXPECT().GetCommentByID(gomock.Any(), gomock.Any()).Times(1).Return(&entity.Comment{}, nil)
				repo.EXPECT().DeleteComment(gomock.Any(), gomock.Eq(commentID)).Times(1).Return(errors.New("no permission, only owner of this like can do this"))
			},
			err: errors.New("no permission, only owner of this like can do this"),
		},
		{
			name: "CommentIDNotFound",
			comment: &entity.CommentDeleteReqest{
				CommentID: commentID,
				UserID:    userID,
			},
			buildStubs: func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository) {
				userRepo.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Times(1).Return(&userEntity.User{}, nil)
				repo.EXPECT().GetCommentByID(gomock.Any(), gomock.Any()).Times(1).Return(nil, errors.New("comment_id not found"))
			},
			err: errors.New("comment_id not found"),
		},
		{
			name: "UserIDNotFound",
			comment: &entity.CommentDeleteReqest{
				CommentID: commentID,
				UserID:    userID,
			},
			buildStubs: func(repo *mock_comment.MockRepository, userRepo *mock_user.MockRepository) {
				userRepo.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Times(1).Return(nil, errors.New("user_id not found"))
			},
			err: errors.New("user_id not found"),
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCommentRepo := mock_comment.NewMockRepository(ctrl)
			mockUserRepo := mock_user.NewMockRepository(ctrl)

			commentUsecase := NewCommentUsecase(mockCommentRepo, mockUserRepo)
			tc.buildStubs(mockCommentRepo, mockUserRepo)

			err := commentUsecase.DeleteComment(context.Background(), tc.comment)
			require.Equal(t, err, tc.err)
		})
	}

}
