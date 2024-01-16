package repository

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/comment/entity"
	"github.com/BuiNhatTruong99/TikTok-Go/pkg/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func setUpTest(t *testing.T) (sqlmock.Sqlmock, *commentRepo) {
	sqlMock, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlMock,
	}), &gorm.Config{})
	require.NoError(t, err)

	commentRepository := &commentRepo{db: gormDB}
	return mock, commentRepository
}

func TestCommentRepo_CreateComment(t *testing.T) {
	mock, commentRepository := setUpTest(t)

	commentReq := &entity.CommentReqest{
		UserID: utils.RandomInt(1, 5),
		PostID: utils.RandomInt(1, 5),
		Text:   utils.RandomString(15),
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "comments"`).WithArgs(commentReq.UserID, commentReq.PostID, commentReq.Text).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := commentRepository.CreateComment(context.Background(), commentReq)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCommentRepo_DeleteComment(t *testing.T) {
	mock, commentRepository := setUpTest(t)

	commetID := utils.RandomInt(1, 5)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "comments" WHERE id = \$1`).WithArgs(commetID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := commentRepository.DeleteComment(context.Background(), commetID)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCommentRepo_GetCommentByID(t *testing.T) {
	mock, commentRepository := setUpTest(t)

	commentID := utils.RandomInt(1, 5)
	commentExpect := &entity.Comment{
		ID:        commentID,
		UserID:    utils.RandomInt(1, 5),
		PostID:    utils.RandomInt(1, 5),
		Text:      utils.RandomString(15),
		CreatedAt: nil,
	}

	mock.ExpectQuery(`SELECT \* FROM "comments" WHERE id = \$1`).WithArgs(commentID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "post_id", "text", "created_at"}).
			AddRow(commentExpect.ID, commentExpect.UserID, commentExpect.PostID, commentExpect.Text, commentExpect.CreatedAt))

	getComment, err := commentRepository.GetCommentByID(context.Background(), commentID)
	require.NoError(t, err)
	require.NotNil(t, getComment)
	require.NoError(t, mock.ExpectationsWereMet())
}
