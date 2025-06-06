// Code generated by MockGen. DO NOT EDIT.
// Source: jwt.go
//
// Generated by this command:
//
//	mockgen -source=jwt.go -destination=./mock/jwt/mock_jwt.go -package=mock_jwt
//

// Package mock_jwt is a generated GoMock package.
package mock_jwt

import (
	reflect "reflect"

	jwt "github.com/rednek21/go-toolkit/jwt"
	gomock "go.uber.org/mock/gomock"
)

// MockTokenManager is a mock of TokenManager interface.
type MockTokenManager struct {
	ctrl     *gomock.Controller
	recorder *MockTokenManagerMockRecorder
	isgomock struct{}
}

// MockTokenManagerMockRecorder is the mock recorder for MockTokenManager.
type MockTokenManagerMockRecorder struct {
	mock *MockTokenManager
}

// NewMockTokenManager creates a new mock instance.
func NewMockTokenManager(ctrl *gomock.Controller) *MockTokenManager {
	mock := &MockTokenManager{ctrl: ctrl}
	mock.recorder = &MockTokenManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenManager) EXPECT() *MockTokenManagerMockRecorder {
	return m.recorder
}

// GenerateTokenPair mocks base method.
func (m *MockTokenManager) GenerateTokenPair(username string) (*jwt.TokenPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokenPair", username)
	ret0, _ := ret[0].(*jwt.TokenPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTokenPair indicates an expected call of GenerateTokenPair.
func (mr *MockTokenManagerMockRecorder) GenerateTokenPair(username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokenPair", reflect.TypeOf((*MockTokenManager)(nil).GenerateTokenPair), username)
}

// ParseToken mocks base method.
func (m *MockTokenManager) ParseToken(tokenString string, isRefresh bool) (*jwt.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", tokenString, isRefresh)
	ret0, _ := ret[0].(*jwt.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockTokenManagerMockRecorder) ParseToken(tokenString, isRefresh any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockTokenManager)(nil).ParseToken), tokenString, isRefresh)
}
