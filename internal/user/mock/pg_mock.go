// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user/interface.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"

	entity "github.com/BuiNhatTruong99/TikTok-Go/internal/user/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetUserByID mocks base method.
func (m *MockRepository) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, userID)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockRepositoryMockRecorder) GetUserByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockRepository)(nil).GetUserByID), ctx, userID)
}

// UpdateAvatarUser mocks base method.
func (m *MockRepository) UpdateAvatarUser(ctx context.Context, userID int64, request *entity.AvatarRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatarUser", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatarUser indicates an expected call of UpdateAvatarUser.
func (mr *MockRepositoryMockRecorder) UpdateAvatarUser(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatarUser", reflect.TypeOf((*MockRepository)(nil).UpdateAvatarUser), ctx, userID, request)
}

// UpdateProfileUser mocks base method.
func (m *MockRepository) UpdateProfileUser(ctx context.Context, userID int64, request *entity.ProfileRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfileUser", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfileUser indicates an expected call of UpdateProfileUser.
func (mr *MockRepositoryMockRecorder) UpdateProfileUser(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfileUser", reflect.TypeOf((*MockRepository)(nil).UpdateProfileUser), ctx, userID, request)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// ChangeAvatar mocks base method.
func (m *MockUseCase) ChangeAvatar(ctx context.Context, userID int64, request *entity.AvatarRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeAvatar", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeAvatar indicates an expected call of ChangeAvatar.
func (mr *MockUseCaseMockRecorder) ChangeAvatar(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAvatar", reflect.TypeOf((*MockUseCase)(nil).ChangeAvatar), ctx, userID, request)
}

// UpdateProfile mocks base method.
func (m *MockUseCase) UpdateProfile(ctx context.Context, userID int64, request *entity.ProfileRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUseCaseMockRecorder) UpdateProfile(ctx, userID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUseCase)(nil).UpdateProfile), ctx, userID, request)
}
