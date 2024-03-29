// Code generated by MockGen. DO NOT EDIT.
// Source: p2p/protocols/protocol.go

// Package quorum is a generated GoMock package.
package quorum

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocols "github.com/simcord-llc/bitbon-system-blockchain/p2p/protocols"
)

// MockMsgPauser is a mock of MsgPauser interface.
type MockMsgPauser struct {
	ctrl     *gomock.Controller
	recorder *MockMsgPauserMockRecorder
}

// MockMsgPauserMockRecorder is the mock recorder for MockMsgPauser.
type MockMsgPauserMockRecorder struct {
	mock *MockMsgPauser
}

// NewMockMsgPauser creates a new mock instance.
func NewMockMsgPauser(ctrl *gomock.Controller) *MockMsgPauser {
	mock := &MockMsgPauser{ctrl: ctrl}
	mock.recorder = &MockMsgPauserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMsgPauser) EXPECT() *MockMsgPauserMockRecorder {
	return m.recorder
}

// Pause mocks base method.
func (m *MockMsgPauser) Pause() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Pause")
}

// Pause indicates an expected call of Pause.
func (mr *MockMsgPauserMockRecorder) Pause() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockMsgPauser)(nil).Pause))
}

// Resume mocks base method.
func (m *MockMsgPauser) Resume() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Resume")
}

// Resume indicates an expected call of Resume.
func (mr *MockMsgPauserMockRecorder) Resume() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resume", reflect.TypeOf((*MockMsgPauser)(nil).Resume))
}

// Wait mocks base method.
func (m *MockMsgPauser) Wait() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Wait")
}

// Wait indicates an expected call of Wait.
func (mr *MockMsgPauserMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockMsgPauser)(nil).Wait))
}

// MockHook is a mock of Hook interface.
type MockHook struct {
	ctrl     *gomock.Controller
	recorder *MockHookMockRecorder
}

// MockHookMockRecorder is the mock recorder for MockHook.
type MockHookMockRecorder struct {
	mock *MockHook
}

// NewMockHook creates a new mock instance.
func NewMockHook(ctrl *gomock.Controller) *MockHook {
	mock := &MockHook{ctrl: ctrl}
	mock.recorder = &MockHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHook) EXPECT() *MockHookMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockHook) Apply(peer *protocols.Peer, costToLocalNode int64, size uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", peer, costToLocalNode, size)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockHookMockRecorder) Apply(peer, costToLocalNode, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockHook)(nil).Apply), peer, costToLocalNode, size)
}

// Validate mocks base method.
func (m *MockHook) Validate(peer *protocols.Peer, size uint32, msg interface{}, payer protocols.Payer) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", peer, size, msg, payer)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockHookMockRecorder) Validate(peer, size, msg, payer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockHook)(nil).Validate), peer, size, msg, payer)
}
