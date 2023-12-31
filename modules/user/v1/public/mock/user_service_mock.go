// Code generated by MockGen. DO NOT EDIT.
// Source: modules/user/v1/public/user_service_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/muhammadchandra19/akira-go-proto/modules/user/v1/public"
	grpc "google.golang.org/grpc"
)

// MockUserProtoClient is a mock of UserProtoClient interface.
type MockUserProtoClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserProtoClientMockRecorder
}

// MockUserProtoClientMockRecorder is the mock recorder for MockUserProtoClient.
type MockUserProtoClientMockRecorder struct {
	mock *MockUserProtoClient
}

// NewMockUserProtoClient creates a new mock instance.
func NewMockUserProtoClient(ctrl *gomock.Controller) *MockUserProtoClient {
	mock := &MockUserProtoClient{ctrl: ctrl}
	mock.recorder = &MockUserProtoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserProtoClient) EXPECT() *MockUserProtoClientMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method.
func (m *MockUserProtoClient) RegisterUser(ctx context.Context, in *v1.User, opts ...grpc.CallOption) (*v1.TokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterUser", varargs...)
	ret0, _ := ret[0].(*v1.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserProtoClientMockRecorder) RegisterUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserProtoClient)(nil).RegisterUser), varargs...)
}

// SearchUser mocks base method.
func (m *MockUserProtoClient) SearchUser(ctx context.Context, in *v1.SearchParams, opts ...grpc.CallOption) (*v1.SearchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchUser", varargs...)
	ret0, _ := ret[0].(*v1.SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUser indicates an expected call of SearchUser.
func (mr *MockUserProtoClientMockRecorder) SearchUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockUserProtoClient)(nil).SearchUser), varargs...)
}

// SignIn mocks base method.
func (m *MockUserProtoClient) SignIn(ctx context.Context, in *v1.SignInRequest, opts ...grpc.CallOption) (*v1.TokenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignIn", varargs...)
	ret0, _ := ret[0].(*v1.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserProtoClientMockRecorder) SignIn(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUserProtoClient)(nil).SignIn), varargs...)
}

// MockUserProtoServer is a mock of UserProtoServer interface.
type MockUserProtoServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserProtoServerMockRecorder
}

// MockUserProtoServerMockRecorder is the mock recorder for MockUserProtoServer.
type MockUserProtoServerMockRecorder struct {
	mock *MockUserProtoServer
}

// NewMockUserProtoServer creates a new mock instance.
func NewMockUserProtoServer(ctrl *gomock.Controller) *MockUserProtoServer {
	mock := &MockUserProtoServer{ctrl: ctrl}
	mock.recorder = &MockUserProtoServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserProtoServer) EXPECT() *MockUserProtoServerMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method.
func (m *MockUserProtoServer) RegisterUser(arg0 context.Context, arg1 *v1.User) (*v1.TokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0, arg1)
	ret0, _ := ret[0].(*v1.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserProtoServerMockRecorder) RegisterUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserProtoServer)(nil).RegisterUser), arg0, arg1)
}

// SearchUser mocks base method.
func (m *MockUserProtoServer) SearchUser(arg0 context.Context, arg1 *v1.SearchParams) (*v1.SearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUser", arg0, arg1)
	ret0, _ := ret[0].(*v1.SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUser indicates an expected call of SearchUser.
func (mr *MockUserProtoServerMockRecorder) SearchUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUser", reflect.TypeOf((*MockUserProtoServer)(nil).SearchUser), arg0, arg1)
}

// SignIn mocks base method.
func (m *MockUserProtoServer) SignIn(arg0 context.Context, arg1 *v1.SignInRequest) (*v1.TokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", arg0, arg1)
	ret0, _ := ret[0].(*v1.TokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserProtoServerMockRecorder) SignIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUserProtoServer)(nil).SignIn), arg0, arg1)
}

// mustEmbedUnimplementedUserProtoServer mocks base method.
func (m *MockUserProtoServer) mustEmbedUnimplementedUserProtoServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserProtoServer")
}

// mustEmbedUnimplementedUserProtoServer indicates an expected call of mustEmbedUnimplementedUserProtoServer.
func (mr *MockUserProtoServerMockRecorder) mustEmbedUnimplementedUserProtoServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserProtoServer", reflect.TypeOf((*MockUserProtoServer)(nil).mustEmbedUnimplementedUserProtoServer))
}

// MockUnsafeUserProtoServer is a mock of UnsafeUserProtoServer interface.
type MockUnsafeUserProtoServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserProtoServerMockRecorder
}

// MockUnsafeUserProtoServerMockRecorder is the mock recorder for MockUnsafeUserProtoServer.
type MockUnsafeUserProtoServerMockRecorder struct {
	mock *MockUnsafeUserProtoServer
}

// NewMockUnsafeUserProtoServer creates a new mock instance.
func NewMockUnsafeUserProtoServer(ctrl *gomock.Controller) *MockUnsafeUserProtoServer {
	mock := &MockUnsafeUserProtoServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserProtoServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserProtoServer) EXPECT() *MockUnsafeUserProtoServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserProtoServer mocks base method.
func (m *MockUnsafeUserProtoServer) mustEmbedUnimplementedUserProtoServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserProtoServer")
}

// mustEmbedUnimplementedUserProtoServer indicates an expected call of mustEmbedUnimplementedUserProtoServer.
func (mr *MockUnsafeUserProtoServerMockRecorder) mustEmbedUnimplementedUserProtoServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserProtoServer", reflect.TypeOf((*MockUnsafeUserProtoServer)(nil).mustEmbedUnimplementedUserProtoServer))
}
