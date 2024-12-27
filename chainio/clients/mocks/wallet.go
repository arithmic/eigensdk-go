// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/arithmic/eigensdk-go/chainio/clients/wallet (interfaces: Wallet)
//
// Generated by this command:
//
//	mockgen -destination=./clients/mocks/wallet.go -package=mocks github.com/arithmic/eigensdk-go/chainio/clients/wallet Wallet
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockWallet is a mock of Wallet interface.
type MockWallet struct {
	ctrl     *gomock.Controller
	recorder *MockWalletMockRecorder
}

// MockWalletMockRecorder is the mock recorder for MockWallet.
type MockWalletMockRecorder struct {
	mock *MockWallet
}

// NewMockWallet creates a new mock instance.
func NewMockWallet(ctrl *gomock.Controller) *MockWallet {
	mock := &MockWallet{ctrl: ctrl}
	mock.recorder = &MockWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWallet) EXPECT() *MockWalletMockRecorder {
	return m.recorder
}

// GetTransactionReceipt mocks base method.
func (m *MockWallet) GetTransactionReceipt(arg0 context.Context, arg1 string) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionReceipt", arg0, arg1)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionReceipt indicates an expected call of GetTransactionReceipt.
func (mr *MockWalletMockRecorder) GetTransactionReceipt(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionReceipt", reflect.TypeOf((*MockWallet)(nil).GetTransactionReceipt), arg0, arg1)
}

// SendTransaction mocks base method.
func (m *MockWallet) SendTransaction(arg0 context.Context, arg1 *types.Transaction) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockWalletMockRecorder) SendTransaction(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockWallet)(nil).SendTransaction), arg0, arg1)
}

// SenderAddress mocks base method.
func (m *MockWallet) SenderAddress(arg0 context.Context) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SenderAddress", arg0)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SenderAddress indicates an expected call of SenderAddress.
func (mr *MockWalletMockRecorder) SenderAddress(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SenderAddress", reflect.TypeOf((*MockWallet)(nil).SenderAddress), arg0)
}
