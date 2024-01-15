package repository

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/like/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func setUpTest(t *testing.T) (sqlmock.Sqlmock, *likeRepo) {
	sqlMock, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlMock,
	}), &gorm.Config{})
	require.NoError(t, err)
	likeRepo := &likeRepo{db: gormDB}
	return mock, likeRepo
}

func TestLikeRepo_CreateLike(t *testing.T) {
	mock, likeRepository := setUpTest(t)

	likeReq := &entity.LikeRequest{
		UserID: 1,
		PostID: 1,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "likes"`).
		WithArgs(likeReq.UserID, likeReq.PostID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	like, err := likeRepository.CreateLike(context.Background(), likeReq)
	require.NoError(t, err)
	require.NotNil(t, like)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestLikeRepo_GetLikeByUserAndPost(t *testing.T) {
	mock, likeRepository := setUpTest(t)
	userID := utils.RandomInt(1, 5)
	postID := utils.RandomInt(1, 5)
	likeReq := &entity.LikeRequest{
		UserID: userID,
		PostID: postID,
	}

	likeExpect := &entity.Like{
		ID:        utils.RandomInt(1, 5),
		UserID:    userID,
		PostID:    postID,
		CreatedAt: nil,
	}

	mock.ExpectQuery(`SELECT \* FROM "likes" WHERE user_id = \$1 AND post_id = \$2`).
		WithArgs(likeReq.UserID, likeReq.PostID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "post_id", "created_at"}).
			AddRow(likeExpect.ID, likeExpect.UserID, likeExpect.PostID, likeExpect.CreatedAt))

	getLike, err := likeRepository.GetLikeByUserAndPost(context.Background(), likeReq)
	require.NoError(t, err)
	require.NotNil(t, getLike)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestLikeRepo_GetLikeByID(t *testing.T) {
	mock, likeRepository := setUpTest(t)
	userID := utils.RandomInt(1, 5)
	postID := utils.RandomInt(1, 5)
	likeExpect := &entity.Like{
		ID:        utils.RandomInt(1, 5),
		UserID:    userID,
		PostID:    postID,
		CreatedAt: nil,
	}

	mock.ExpectQuery(`SELECT \* FROM "likes" WHERE id = \$1`).
		WithArgs(likeExpect.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "post_id", "created_at"}).
			AddRow(likeExpect.ID, likeExpect.UserID, likeExpect.PostID, likeExpect.CreatedAt))

	getLike, err := likeRepository.GetLikeByID(context.Background(), likeExpect.ID)
	require.NoError(t, err)
	require.NotNil(t, getLike)
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLikeRepo_DeleteLike(t *testing.T) {
	mock, likeRepository := setUpTest(t)
	userID := utils.RandomInt(1, 5)
	postID := utils.RandomInt(1, 5)
	likeDelete := &entity.Like{
		ID:        utils.RandomInt(1, 5),
		UserID:    userID,
		PostID:    postID,
		CreatedAt: nil,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "likes" WHERE id = \$1`).
		WithArgs(likeDelete.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := likeRepository.DeleteLike(context.Background(), likeDelete.ID)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
