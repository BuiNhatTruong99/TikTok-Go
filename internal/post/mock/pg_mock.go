// Code generated by MockGen. DO NOT EDIT.
// Source: internal/post/interface.go

// Package mock_post is a generated GoMock package.
package mock_post

import (
	context "context"
	reflect "reflect"

	entity "github.com/BuiNhatTruong99/TikTok-Go/internal/post/entity"
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

// CreatePost mocks base method.
func (m *MockRepository) CreatePost(ctx context.Context, request *entity.PostRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockRepositoryMockRecorder) CreatePost(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockRepository)(nil).CreatePost), ctx, request)
}

// DeletePostByID mocks base method.
func (m *MockRepository) DeletePostByID(ctx context.Context, postID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostByID", ctx, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostByID indicates an expected call of DeletePostByID.
func (mr *MockRepositoryMockRecorder) DeletePostByID(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostByID", reflect.TypeOf((*MockRepository)(nil).DeletePostByID), ctx, postID)
}

// GetPostByID mocks base method.
func (m *MockRepository) GetPostByID(ctx context.Context, postID int64) (*entity.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", ctx, postID)
	ret0, _ := ret[0].(*entity.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockRepositoryMockRecorder) GetPostByID(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockRepository)(nil).GetPostByID), ctx, postID)
}

// GetPostByUserID mocks base method.
func (m *MockRepository) GetPostByUserID(ctx context.Context, userID int64) ([]entity.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByUserID", ctx, userID)
	ret0, _ := ret[0].([]entity.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByUserID indicates an expected call of GetPostByUserID.
func (mr *MockRepositoryMockRecorder) GetPostByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByUserID", reflect.TypeOf((*MockRepository)(nil).GetPostByUserID), ctx, userID)
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

// CreatePost mocks base method.
func (m *MockUseCase) CreatePost(ctx context.Context, request *entity.PostRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockUseCaseMockRecorder) CreatePost(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockUseCase)(nil).CreatePost), ctx, request)
}

// DeletePost mocks base method.
func (m *MockUseCase) DeletePost(ctx context.Context, postID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockUseCaseMockRecorder) DeletePost(ctx, postID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockUseCase)(nil).DeletePost), ctx, postID)
}
