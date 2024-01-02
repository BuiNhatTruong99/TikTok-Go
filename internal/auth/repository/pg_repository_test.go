package repository

import (
	"context"
	"errors"
	"github.com/BuiNhatTruong99/TikTok-Go/internal/auth/entity"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func setUpTest(t *testing.T) (sqlmock.Sqlmock, *authRepo) {
	sqlMock, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlMock,
	}), &gorm.Config{})
	require.NoError(t, err)
	authRepo := &authRepo{db: gormDB}
	return mock, authRepo
}

func TestAuthRepo_Register(t *testing.T) {
	mock, authRepository := setUpTest(t)

	userRequest := &entity.UserRequest{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "123123",
		AvatarUrl:    "avatar_url",
		Bio:          "this is my tiktok",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(userRequest.Username, userRequest.Email, userRequest.HashPassword, userRequest.AvatarUrl,
			userRequest.Bio).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := authRepository.Register(context.Background(), userRequest)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthRepo_Register_Error(t *testing.T) {
	mock, authRepo := setUpTest(t)

	userRequest := &entity.UserRequest{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "123123",
		AvatarUrl:    "avatar_url",
		Bio:          "this is my tiktok",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(userRequest.Username, userRequest.Email, userRequest.HashPassword, userRequest.AvatarUrl, userRequest.Bio).
		WillReturnError(errors.New("database error"))
	mock.ExpectRollback()

	err := authRepo.Register(context.Background(), userRequest)
	require.Error(t, err)
	require.EqualError(t, err, "database error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthRepo_GetUserByEmail(t *testing.T) {
	mock, authRepository := setUpTest(t)

	testdUser := &entity.User{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatart_url",
		Bio:          "this is my tiktok",
	}
	expectedUser := &entity.User{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatar_url",
		Bio:          "this is my tiktok",
	}

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1`).
		WithArgs(testdUser.Email).
		WillReturnRows(sqlmock.NewRows([]string{"username", "email", "avatar_url", "bio"}).
			AddRow(expectedUser.Username, expectedUser.Email, expectedUser.AvatarUrl,
				expectedUser.Bio))

	user, err := authRepository.GetUserByEmail(context.Background(), testdUser.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Email, testdUser.Email)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthRepo_GetUserByEmail_Error(t *testing.T) {
	mock, authRepository := setUpTest(t)

	testdUser := &entity.User{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatart_url",
		Bio:          "this is my tiktok",
	}

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1`).
		WithArgs(testdUser.Email).
		WillReturnError(errors.New("database error"))

	user, err := authRepository.GetUserByEmail(context.Background(), testdUser.Email)
	require.Error(t, err)
	require.Nil(t, user)
	require.EqualError(t, err, "database error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthRepo_GetUserByEmail_NotFound(t *testing.T) {
	mock, authRepo := setUpTest(t)

	testUser := &entity.User{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatart_url",
		Bio:          "this is my tiktok",
	}

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1`).
		WithArgs(testUser.Email).
		WillReturnError(gorm.ErrRecordNotFound)

	user, err := authRepo.GetUserByEmail(context.Background(), testUser.Email)
	require.Error(t, err)
	require.Nil(t, user)
	require.EqualError(t, err, "email is not exists")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthRepo_GetUserByName(t *testing.T) {
	mock, authRepository := setUpTest(t)

	testdUser := &entity.User{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatar_url",
		Bio:          "this is my tiktok",
	}
	expectedUser := &entity.User{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatar_url",
		Bio:          "this is my tiktok",
	}

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE username = \$1`).
		WithArgs(testdUser.Username).
		WillReturnRows(sqlmock.NewRows([]string{"username", "email", "avatar_url", "bio"}).
			AddRow(expectedUser.Username, expectedUser.Email, expectedUser.AvatarUrl,
				expectedUser.Bio))

	user, err := authRepository.GetUserByName(context.Background(), testdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, testdUser.Username)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthRepo_GetUserByName_Error(t *testing.T) {
	mock, authRepository := setUpTest(t)

	testUser := &entity.User{
		Username:     "truongbui",
		Email:        "bntruong@gmail.com",
		HashPassword: "abc",
		AvatarUrl:    "avatar_url",
		Bio:          "this is my tiktok",
	}
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE username = \$1`).
		WithArgs(testUser.Username).
		WillReturnError(errors.New("database error"))

	user, err := authRepository.GetUserByName(context.Background(), testUser.Username)
	require.Error(t, err)
	require.Nil(t, user)
	require.EqualError(t, err, "database error")
	require.NoError(t, mock.ExpectationsWereMet())
}
