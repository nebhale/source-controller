// Code generated by MockGen. DO NOT EDIT.
// Source: gcp.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "cloud.google.com/go/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Bucket mocks base method.
func (m *MockClient) Bucket(arg0 string) *storage.BucketHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bucket", arg0)
	ret0, _ := ret[0].(*storage.BucketHandle)
	return ret0
}

// Bucket indicates an expected call of Bucket.
func (mr *MockClientMockRecorder) Bucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bucket", reflect.TypeOf((*MockClient)(nil).Bucket), arg0)
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// MockBucketHandle is a mock of BucketHandle interface.
type MockBucketHandle struct {
	ctrl     *gomock.Controller
	recorder *MockBucketHandleMockRecorder
}

// MockBucketHandleMockRecorder is the mock recorder for MockBucketHandle.
type MockBucketHandleMockRecorder struct {
	mock *MockBucketHandle
}

// NewMockBucketHandle creates a new mock instance.
func NewMockBucketHandle(ctrl *gomock.Controller) *MockBucketHandle {
	mock := &MockBucketHandle{ctrl: ctrl}
	mock.recorder = &MockBucketHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBucketHandle) EXPECT() *MockBucketHandleMockRecorder {
	return m.recorder
}

// Attrs mocks base method.
func (m *MockBucketHandle) Attrs(arg0 context.Context) (*storage.BucketAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attrs", arg0)
	ret0, _ := ret[0].(*storage.BucketAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attrs indicates an expected call of Attrs.
func (mr *MockBucketHandleMockRecorder) Attrs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*MockBucketHandle)(nil).Attrs), arg0)
}

// Create mocks base method.
func (m *MockBucketHandle) Create(arg0 context.Context, arg1 string, arg2 *storage.BucketAttrs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBucketHandleMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBucketHandle)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockBucketHandle) Delete(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBucketHandleMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBucketHandle)(nil).Delete), arg0)
}

// Object mocks base method.
func (m *MockBucketHandle) Object(arg0 string) *storage.ObjectHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Object", arg0)
	ret0, _ := ret[0].(*storage.ObjectHandle)
	return ret0
}

// Object indicates an expected call of Object.
func (mr *MockBucketHandleMockRecorder) Object(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Object", reflect.TypeOf((*MockBucketHandle)(nil).Object), arg0)
}

// Objects mocks base method.
func (m *MockBucketHandle) Objects(arg0 context.Context, arg1 *storage.Query) *storage.ObjectIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Objects", arg0, arg1)
	ret0, _ := ret[0].(*storage.ObjectIterator)
	return ret0
}

// Objects indicates an expected call of Objects.
func (mr *MockBucketHandleMockRecorder) Objects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Objects", reflect.TypeOf((*MockBucketHandle)(nil).Objects), arg0, arg1)
}

// MockObjectHandle is a mock of ObjectHandle interface.
type MockObjectHandle struct {
	ctrl     *gomock.Controller
	recorder *MockObjectHandleMockRecorder
}

// MockObjectHandleMockRecorder is the mock recorder for MockObjectHandle.
type MockObjectHandleMockRecorder struct {
	mock *MockObjectHandle
}

// NewMockObjectHandle creates a new mock instance.
func NewMockObjectHandle(ctrl *gomock.Controller) *MockObjectHandle {
	mock := &MockObjectHandle{ctrl: ctrl}
	mock.recorder = &MockObjectHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectHandle) EXPECT() *MockObjectHandleMockRecorder {
	return m.recorder
}

// Attrs mocks base method.
func (m *MockObjectHandle) Attrs(arg0 context.Context) (*storage.ObjectAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attrs", arg0)
	ret0, _ := ret[0].(*storage.ObjectAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attrs indicates an expected call of Attrs.
func (mr *MockObjectHandleMockRecorder) Attrs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*MockObjectHandle)(nil).Attrs), arg0)
}

// NewRangeReader mocks base method.
func (m *MockObjectHandle) NewRangeReader(arg0 context.Context, arg1, arg2 int64) (*storage.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRangeReader", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRangeReader indicates an expected call of NewRangeReader.
func (mr *MockObjectHandleMockRecorder) NewRangeReader(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRangeReader", reflect.TypeOf((*MockObjectHandle)(nil).NewRangeReader), arg0, arg1, arg2)
}
