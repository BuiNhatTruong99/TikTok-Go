package repository

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/post/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func setUpTest(t *testing.T) (sqlmock.Sqlmock, *postRepo) {
	sqlMock, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlMock,
	}), &gorm.Config{})
	require.NoError(t, err)
	postRepo := &postRepo{db: gormDB}
	return mock, postRepo
}

func TestPostRepo_CreatePost(t *testing.T) {
	mock, postRepository := setUpTest(t)

	postRequest := &entity.PostRequest{
		UserID:   1,
		VideoUrl: "default_url",
		Caption:  "This is post 1",
	}
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "posts"`).
		WithArgs(postRequest.UserID, postRequest.VideoUrl, postRequest.Caption).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := postRepository.CreatePost(context.Background(), postRequest)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostRepo_GetPostByID(t *testing.T) {
	mock, postRepository := setUpTest(t)

	testPost := &entity.Post{
		ID:        1,
		UserID:    1,
		VideoUrl:  "default_url",
		Caption:   "This is post 1",
		CreatedAt: nil,
	}

	expectedPost := &entity.Post{
		ID:        1,
		UserID:    1,
		VideoUrl:  "default_url",
		Caption:   "This is post 1",
		CreatedAt: nil,
	}

	mock.ExpectQuery(`SELECT \* FROM "posts" WHERE id = \$1`).
		WithArgs(testPost.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "video_url", "caption", "created_at"}).
			AddRow(expectedPost.ID, expectedPost.UserID, expectedPost.VideoUrl, expectedPost.Caption,
				expectedPost.CreatedAt))
	post, err := postRepository.GetPostByID(context.Background(), testPost.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, post.ID, expectedPost.ID)
	require.Equal(t, post.UserID, expectedPost.UserID)
	require.Equal(t, post.VideoUrl, expectedPost.VideoUrl)
	require.Equal(t, post.Caption, expectedPost.Caption)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestPostRepo_GetPostByUserID(t *testing.T) {
	mock, postRepository := setUpTest(t)

	userID := int64(123)
	expectedPosts := []entity.Post{
		{ID: 1, UserID: userID, VideoUrl: "default_url", Caption: "Post 1", CreatedAt: nil},
		{ID: 2, UserID: userID, VideoUrl: "default_url2", Caption: "Post 2", CreatedAt: nil},
	}

	mock.ExpectQuery(`SELECT \* FROM "posts" WHERE user_id = \$1`).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "video_url", "caption", "created_at"}).
			AddRow(expectedPosts[0].ID, expectedPosts[0].UserID, expectedPosts[0].VideoUrl, expectedPosts[0].Caption, expectedPosts[0].CreatedAt).
			AddRow(expectedPosts[1].ID, expectedPosts[1].UserID, expectedPosts[1].VideoUrl, expectedPosts[1].Caption, expectedPosts[1].CreatedAt))

	posts, err := postRepository.GetPostByUserID(context.Background(), userID)
	require.NoError(t, err)
	require.NotNil(t, posts)
	require.Equal(t, expectedPosts, posts)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestPostRepo_DeletePostByID(t *testing.T) {
	mock, postRepository := setUpTest(t)

	postID := int64(4)
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "posts" WHERE id = \$1`).
		WithArgs(postID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := postRepository.DeletePostByID(context.Background(), postID)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
