// Code generated by MockGen. DO NOT EDIT.
// Source: modules/chat/v1/public/chat_service_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/muhammadchandra19/akira-go-proto/modules/chat/v1/public"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockChatServiceClient is a mock of ChatServiceClient interface.
type MockChatServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceClientMockRecorder
}

// MockChatServiceClientMockRecorder is the mock recorder for MockChatServiceClient.
type MockChatServiceClientMockRecorder struct {
	mock *MockChatServiceClient
}

// NewMockChatServiceClient creates a new mock instance.
func NewMockChatServiceClient(ctrl *gomock.Controller) *MockChatServiceClient {
	mock := &MockChatServiceClient{ctrl: ctrl}
	mock.recorder = &MockChatServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServiceClient) EXPECT() *MockChatServiceClientMockRecorder {
	return m.recorder
}

// AddUserToRoom mocks base method.
func (m *MockChatServiceClient) AddUserToRoom(ctx context.Context, in *v1.UserRoom, opts ...grpc.CallOption) (*v1.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUserToRoom", varargs...)
	ret0, _ := ret[0].(*v1.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUserToRoom indicates an expected call of AddUserToRoom.
func (mr *MockChatServiceClientMockRecorder) AddUserToRoom(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToRoom", reflect.TypeOf((*MockChatServiceClient)(nil).AddUserToRoom), varargs...)
}

// CreateRoom mocks base method.
func (m *MockChatServiceClient) CreateRoom(ctx context.Context, in *v1.Room, opts ...grpc.CallOption) (*v1.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRoom", varargs...)
	ret0, _ := ret[0].(*v1.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockChatServiceClientMockRecorder) CreateRoom(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockChatServiceClient)(nil).CreateRoom), varargs...)
}

// CreateStream mocks base method.
func (m *MockChatServiceClient) CreateStream(ctx context.Context, in *v1.StreamConnect, opts ...grpc.CallOption) (v1.ChatService_CreateStreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateStream", varargs...)
	ret0, _ := ret[0].(v1.ChatService_CreateStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStream indicates an expected call of CreateStream.
func (mr *MockChatServiceClientMockRecorder) CreateStream(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStream", reflect.TypeOf((*MockChatServiceClient)(nil).CreateStream), varargs...)
}

// SendMessage mocks base method.
func (m *MockChatServiceClient) SendMessage(ctx context.Context, in *v1.ContentMessage, opts ...grpc.CallOption) (*v1.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessage", varargs...)
	ret0, _ := ret[0].(*v1.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockChatServiceClientMockRecorder) SendMessage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockChatServiceClient)(nil).SendMessage), varargs...)
}

// MockChatService_CreateStreamClient is a mock of ChatService_CreateStreamClient interface.
type MockChatService_CreateStreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatService_CreateStreamClientMockRecorder
}

// MockChatService_CreateStreamClientMockRecorder is the mock recorder for MockChatService_CreateStreamClient.
type MockChatService_CreateStreamClientMockRecorder struct {
	mock *MockChatService_CreateStreamClient
}

// NewMockChatService_CreateStreamClient creates a new mock instance.
func NewMockChatService_CreateStreamClient(ctrl *gomock.Controller) *MockChatService_CreateStreamClient {
	mock := &MockChatService_CreateStreamClient{ctrl: ctrl}
	mock.recorder = &MockChatService_CreateStreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatService_CreateStreamClient) EXPECT() *MockChatService_CreateStreamClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockChatService_CreateStreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockChatService_CreateStreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockChatService_CreateStreamClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockChatService_CreateStreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockChatService_CreateStreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChatService_CreateStreamClient)(nil).Context))
}

// Header mocks base method.
func (m *MockChatService_CreateStreamClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockChatService_CreateStreamClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockChatService_CreateStreamClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockChatService_CreateStreamClient) Recv() (*v1.ResponseStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.ResponseStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockChatService_CreateStreamClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockChatService_CreateStreamClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockChatService_CreateStreamClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockChatService_CreateStreamClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChatService_CreateStreamClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockChatService_CreateStreamClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockChatService_CreateStreamClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChatService_CreateStreamClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockChatService_CreateStreamClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockChatService_CreateStreamClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockChatService_CreateStreamClient)(nil).Trailer))
}

// MockChatServiceServer is a mock of ChatServiceServer interface.
type MockChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceServerMockRecorder
}

// MockChatServiceServerMockRecorder is the mock recorder for MockChatServiceServer.
type MockChatServiceServerMockRecorder struct {
	mock *MockChatServiceServer
}

// NewMockChatServiceServer creates a new mock instance.
func NewMockChatServiceServer(ctrl *gomock.Controller) *MockChatServiceServer {
	mock := &MockChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServiceServer) EXPECT() *MockChatServiceServerMockRecorder {
	return m.recorder
}

// AddUserToRoom mocks base method.
func (m *MockChatServiceServer) AddUserToRoom(arg0 context.Context, arg1 *v1.UserRoom) (*v1.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToRoom", arg0, arg1)
	ret0, _ := ret[0].(*v1.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUserToRoom indicates an expected call of AddUserToRoom.
func (mr *MockChatServiceServerMockRecorder) AddUserToRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToRoom", reflect.TypeOf((*MockChatServiceServer)(nil).AddUserToRoom), arg0, arg1)
}

// CreateRoom mocks base method.
func (m *MockChatServiceServer) CreateRoom(arg0 context.Context, arg1 *v1.Room) (*v1.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", arg0, arg1)
	ret0, _ := ret[0].(*v1.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockChatServiceServerMockRecorder) CreateRoom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockChatServiceServer)(nil).CreateRoom), arg0, arg1)
}

// CreateStream mocks base method.
func (m *MockChatServiceServer) CreateStream(arg0 *v1.StreamConnect, arg1 v1.ChatService_CreateStreamServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStream indicates an expected call of CreateStream.
func (mr *MockChatServiceServerMockRecorder) CreateStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStream", reflect.TypeOf((*MockChatServiceServer)(nil).CreateStream), arg0, arg1)
}

// SendMessage mocks base method.
func (m *MockChatServiceServer) SendMessage(arg0 context.Context, arg1 *v1.ContentMessage) (*v1.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0, arg1)
	ret0, _ := ret[0].(*v1.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockChatServiceServerMockRecorder) SendMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockChatServiceServer)(nil).SendMessage), arg0, arg1)
}

// mustEmbedUnimplementedChatServiceServer mocks base method.
func (m *MockChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatServiceServer")
}

// mustEmbedUnimplementedChatServiceServer indicates an expected call of mustEmbedUnimplementedChatServiceServer.
func (mr *MockChatServiceServerMockRecorder) mustEmbedUnimplementedChatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatServiceServer", reflect.TypeOf((*MockChatServiceServer)(nil).mustEmbedUnimplementedChatServiceServer))
}

// MockUnsafeChatServiceServer is a mock of UnsafeChatServiceServer interface.
type MockUnsafeChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeChatServiceServerMockRecorder
}

// MockUnsafeChatServiceServerMockRecorder is the mock recorder for MockUnsafeChatServiceServer.
type MockUnsafeChatServiceServerMockRecorder struct {
	mock *MockUnsafeChatServiceServer
}

// NewMockUnsafeChatServiceServer creates a new mock instance.
func NewMockUnsafeChatServiceServer(ctrl *gomock.Controller) *MockUnsafeChatServiceServer {
	mock := &MockUnsafeChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeChatServiceServer) EXPECT() *MockUnsafeChatServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedChatServiceServer mocks base method.
func (m *MockUnsafeChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatServiceServer")
}

// mustEmbedUnimplementedChatServiceServer indicates an expected call of mustEmbedUnimplementedChatServiceServer.
func (mr *MockUnsafeChatServiceServerMockRecorder) mustEmbedUnimplementedChatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatServiceServer", reflect.TypeOf((*MockUnsafeChatServiceServer)(nil).mustEmbedUnimplementedChatServiceServer))
}

// MockChatService_CreateStreamServer is a mock of ChatService_CreateStreamServer interface.
type MockChatService_CreateStreamServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatService_CreateStreamServerMockRecorder
}

// MockChatService_CreateStreamServerMockRecorder is the mock recorder for MockChatService_CreateStreamServer.
type MockChatService_CreateStreamServerMockRecorder struct {
	mock *MockChatService_CreateStreamServer
}

// NewMockChatService_CreateStreamServer creates a new mock instance.
func NewMockChatService_CreateStreamServer(ctrl *gomock.Controller) *MockChatService_CreateStreamServer {
	mock := &MockChatService_CreateStreamServer{ctrl: ctrl}
	mock.recorder = &MockChatService_CreateStreamServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatService_CreateStreamServer) EXPECT() *MockChatService_CreateStreamServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockChatService_CreateStreamServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockChatService_CreateStreamServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChatService_CreateStreamServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockChatService_CreateStreamServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockChatService_CreateStreamServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChatService_CreateStreamServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockChatService_CreateStreamServer) Send(arg0 *v1.ResponseStream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockChatService_CreateStreamServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChatService_CreateStreamServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockChatService_CreateStreamServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockChatService_CreateStreamServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockChatService_CreateStreamServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockChatService_CreateStreamServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockChatService_CreateStreamServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChatService_CreateStreamServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockChatService_CreateStreamServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockChatService_CreateStreamServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockChatService_CreateStreamServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockChatService_CreateStreamServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockChatService_CreateStreamServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockChatService_CreateStreamServer)(nil).SetTrailer), arg0)
}
