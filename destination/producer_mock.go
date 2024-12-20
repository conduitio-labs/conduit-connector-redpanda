// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio-labs/conduit-connector-redpanda/destination (interfaces: Producer)
//
// Generated by this command:
//
//	mockgen -destination producer_mock.go -package destination -mock_names=Producer=MockProducer . Producer
//

// Package destination is a generated GoMock package.
package destination

import (
	context "context"
	reflect "reflect"

	opencdc "github.com/conduitio/conduit-commons/opencdc"
	gomock "go.uber.org/mock/gomock"
)

// MockProducer is a mock of Producer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
	isgomock struct{}
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockProducer) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockProducerMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProducer)(nil).Close), arg0)
}

// Produce mocks base method.
func (m *MockProducer) Produce(arg0 context.Context, arg1 []opencdc.Record) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce.
func (mr *MockProducerMockRecorder) Produce(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockProducer)(nil).Produce), arg0, arg1)
}
