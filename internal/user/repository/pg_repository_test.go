package repository

import (
	"context"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func setUpTest(t *testing.T) (sqlmock.Sqlmock, *userRepo) {
	sqlMock, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlMock,
	}), &gorm.Config{})
	require.NoError(t, err)
	userRepo := &userRepo{db: gormDB}
	return mock, userRepo
}

func TestUserRepo_GetUserByID(t *testing.T) {
	mock, userRepository := setUpTest(t)

	testdUser := &entity.User{
		ID:           1,
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatart_url",
		Bio:          "this is my tiktok",
		CreatedAt:    nil,
	}
	expectedUser := &entity.User{
		ID:           1,
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatar_url",
		Bio:          "this is my tiktok",
		CreatedAt:    nil,
	}

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1`).
		WithArgs(testdUser.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "email", "avatar_url", "bio", "created_at"}).
			AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Email, expectedUser.AvatarUrl,
				expectedUser.Bio, expectedUser.CreatedAt))

	user, err := userRepository.GetUserByID(context.Background(), testdUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Email, testdUser.Email)
	require.Equal(t, user.ID, testdUser.ID)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepo_UpdateProfileUser(t *testing.T) {
	mock, userRepository := setUpTest(t)

	userID := int64(2)
	profileRequest := &entity.ProfileRequest{
		Username: "Truongbui",
		Bio:      "My tiktok account",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "username"=\$1,"bio"=\$2 WHERE id = \$3`).
		WithArgs(profileRequest.Username, profileRequest.Bio, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userRepository.UpdateProfileUser(context.Background(), userID, profileRequest)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepo_UpdateAvatarUser(t *testing.T) {
	mock, userRepository := setUpTest(t)

	userID := int64(3)
	avatarRequest := &entity.AvatarRequest{
		AvatarUrl: "https://res.cloudinary.com/ddncxed3m/image/upload/v1704641039/samples/upscale-face-1.jpg",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "avatar_url"=\$1 WHERE id = \$2`).
		WithArgs(avatarRequest.AvatarUrl, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userRepository.UpdateAvatarUser(context.Background(), userID, avatarRequest)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
