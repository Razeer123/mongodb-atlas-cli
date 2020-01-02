// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/clusters.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

// MockClusterCreator is a mock of ClusterCreator interface
type MockClusterCreator struct {
	ctrl     *gomock.Controller
	recorder *MockClusterCreatorMockRecorder
}

// MockClusterCreatorMockRecorder is the mock recorder for MockClusterCreator
type MockClusterCreatorMockRecorder struct {
	mock *MockClusterCreator
}

// NewMockClusterCreator creates a new mock instance
func NewMockClusterCreator(ctrl *gomock.Controller) *MockClusterCreator {
	mock := &MockClusterCreator{ctrl: ctrl}
	mock.recorder = &MockClusterCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterCreator) EXPECT() *MockClusterCreatorMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockClusterCreator) CreateCluster(arg0 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockClusterCreatorMockRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockClusterCreator)(nil).CreateCluster), arg0)
}

// MockClusterStore is a mock of ClusterStore interface
type MockClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStoreMockRecorder
}

// MockClusterStoreMockRecorder is the mock recorder for MockClusterStore
type MockClusterStoreMockRecorder struct {
	mock *MockClusterStore
}

// NewMockClusterStore creates a new mock instance
func NewMockClusterStore(ctrl *gomock.Controller) *MockClusterStore {
	mock := &MockClusterStore{ctrl: ctrl}
	mock.recorder = &MockClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterStore) EXPECT() *MockClusterStoreMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockClusterStore) CreateCluster(arg0 *mongodbatlas.Cluster) (*mongodbatlas.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", arg0)
	ret0, _ := ret[0].(*mongodbatlas.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockClusterStoreMockRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockClusterStore)(nil).CreateCluster), arg0)
}