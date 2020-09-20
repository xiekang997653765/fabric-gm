// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	ledger "github.com/hyperledger/fabric/core/ledger"
	mock "github.com/stretchr/testify/mock"

	peer "github.com/hyperledger/fabric-protos-go/peer"
)

// LedgerResources is an autogenerated mock type for the LedgerResources type
type LedgerResources struct {
	mock.Mock
}

// GetTransactionByID provides a mock function with given fields: txID
func (_m *LedgerResources) GetTransactionByID(txID string) (*peer.ProcessedTransaction, error) {
	ret := _m.Called(txID)

	var r0 *peer.ProcessedTransaction
	if rf, ok := ret.Get(0).(func(string) *peer.ProcessedTransaction); ok {
		r0 = rf(txID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*peer.ProcessedTransaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(txID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryExecutor provides a mock function with given fields:
func (_m *LedgerResources) NewQueryExecutor() (ledger.QueryExecutor, error) {
	ret := _m.Called()

	var r0 ledger.QueryExecutor
	if rf, ok := ret.Get(0).(func() ledger.QueryExecutor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ledger.QueryExecutor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
