// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/smancke/guble/server (interfaces: PubSubSource)

package gubled

import (
	gomock "github.com/golang/mock/gomock"
	server "github.com/smancke/guble/server"
	auth "github.com/smancke/guble/server/auth"
	store "github.com/smancke/guble/store"
)

// Mock of PubSubSource interface
type MockPubSubSource struct {
	ctrl     *gomock.Controller
	recorder *_MockPubSubSourceRecorder
}

// Recorder for MockPubSubSource (not exported)
type _MockPubSubSourceRecorder struct {
	mock *MockPubSubSource
}

func NewMockPubSubSource(ctrl *gomock.Controller) *MockPubSubSource {
	mock := &MockPubSubSource{ctrl: ctrl}
	mock.recorder = &_MockPubSubSourceRecorder{mock}
	return mock
}

func (_m *MockPubSubSource) EXPECT() *_MockPubSubSourceRecorder {
	return _m.recorder
}

func (_m *MockPubSubSource) AccessManager() (auth.AccessManager, error) {
	ret := _m.ctrl.Call(_m, "AccessManager")
	ret0, _ := ret[0].(auth.AccessManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPubSubSourceRecorder) AccessManager() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AccessManager")
}

func (_m *MockPubSubSource) KVStore() (store.KVStore, error) {
	ret := _m.ctrl.Call(_m, "KVStore")
	ret0, _ := ret[0].(store.KVStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPubSubSourceRecorder) KVStore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KVStore")
}

func (_m *MockPubSubSource) MessageStore() (store.MessageStore, error) {
	ret := _m.ctrl.Call(_m, "MessageStore")
	ret0, _ := ret[0].(store.MessageStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPubSubSourceRecorder) MessageStore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MessageStore")
}

func (_m *MockPubSubSource) Subscribe(_param0 *server.Route) (*server.Route, error) {
	ret := _m.ctrl.Call(_m, "Subscribe", _param0)
	ret0, _ := ret[0].(*server.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPubSubSourceRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Subscribe", arg0)
}

func (_m *MockPubSubSource) Unsubscribe(_param0 *server.Route) {
	_m.ctrl.Call(_m, "Unsubscribe", _param0)
}

func (_mr *_MockPubSubSourceRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unsubscribe", arg0)
}
