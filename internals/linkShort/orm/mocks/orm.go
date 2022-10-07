// Code generated by MockGen. DO NOT EDIT.
// Source: LinkShortening/internals/linkShort/orm (interfaces: LinkShortWrapperInterface,ConnectionPostgresInterface,TransactionInterface,ConnectionRedisInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

// MockLinkShortWrapperInterface is a mock of LinkShortWrapperInterface interface.
type MockLinkShortWrapperInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLinkShortWrapperInterfaceMockRecorder
}

// MockLinkShortWrapperInterfaceMockRecorder is the mock recorder for MockLinkShortWrapperInterface.
type MockLinkShortWrapperInterfaceMockRecorder struct {
	mock *MockLinkShortWrapperInterface
}

// NewMockLinkShortWrapperInterface creates a new mock instance.
func NewMockLinkShortWrapperInterface(ctrl *gomock.Controller) *MockLinkShortWrapperInterface {
	mock := &MockLinkShortWrapperInterface{ctrl: ctrl}
	mock.recorder = &MockLinkShortWrapperInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinkShortWrapperInterface) EXPECT() *MockLinkShortWrapperInterfaceMockRecorder {
	return m.recorder
}

// CreateLinkShort mocks base method.
func (m *MockLinkShortWrapperInterface) CreateLinkShort(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLinkShort", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLinkShort indicates an expected call of CreateLinkShort.
func (mr *MockLinkShortWrapperInterfaceMockRecorder) CreateLinkShort(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLinkShort", reflect.TypeOf((*MockLinkShortWrapperInterface)(nil).CreateLinkShort), arg0, arg1)
}

// CreateLinkShortPostgres mocks base method.
func (m *MockLinkShortWrapperInterface) CreateLinkShortPostgres(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLinkShortPostgres", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLinkShortPostgres indicates an expected call of CreateLinkShortPostgres.
func (mr *MockLinkShortWrapperInterfaceMockRecorder) CreateLinkShortPostgres(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLinkShortPostgres", reflect.TypeOf((*MockLinkShortWrapperInterface)(nil).CreateLinkShortPostgres), arg0, arg1)
}

// CreateLinkShortRedis mocks base method.
func (m *MockLinkShortWrapperInterface) CreateLinkShortRedis(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLinkShortRedis", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLinkShortRedis indicates an expected call of CreateLinkShortRedis.
func (mr *MockLinkShortWrapperInterfaceMockRecorder) CreateLinkShortRedis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLinkShortRedis", reflect.TypeOf((*MockLinkShortWrapperInterface)(nil).CreateLinkShortRedis), arg0, arg1)
}

// TakeLinkFull mocks base method.
func (m *MockLinkShortWrapperInterface) TakeLinkFull(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeLinkFull", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeLinkFull indicates an expected call of TakeLinkFull.
func (mr *MockLinkShortWrapperInterfaceMockRecorder) TakeLinkFull(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeLinkFull", reflect.TypeOf((*MockLinkShortWrapperInterface)(nil).TakeLinkFull), arg0)
}

// TakeLinkFullPostgres mocks base method.
func (m *MockLinkShortWrapperInterface) TakeLinkFullPostgres(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeLinkFullPostgres", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeLinkFullPostgres indicates an expected call of TakeLinkFullPostgres.
func (mr *MockLinkShortWrapperInterfaceMockRecorder) TakeLinkFullPostgres(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeLinkFullPostgres", reflect.TypeOf((*MockLinkShortWrapperInterface)(nil).TakeLinkFullPostgres), arg0)
}

// TakeLinkFullRedis mocks base method.
func (m *MockLinkShortWrapperInterface) TakeLinkFullRedis(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeLinkFullRedis", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeLinkFullRedis indicates an expected call of TakeLinkFullRedis.
func (mr *MockLinkShortWrapperInterfaceMockRecorder) TakeLinkFullRedis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeLinkFullRedis", reflect.TypeOf((*MockLinkShortWrapperInterface)(nil).TakeLinkFullRedis), arg0)
}

// MockConnectionPostgresInterface is a mock of ConnectionPostgresInterface interface.
type MockConnectionPostgresInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionPostgresInterfaceMockRecorder
}

// MockConnectionPostgresInterfaceMockRecorder is the mock recorder for MockConnectionPostgresInterface.
type MockConnectionPostgresInterfaceMockRecorder struct {
	mock *MockConnectionPostgresInterface
}

// NewMockConnectionPostgresInterface creates a new mock instance.
func NewMockConnectionPostgresInterface(ctrl *gomock.Controller) *MockConnectionPostgresInterface {
	mock := &MockConnectionPostgresInterface{ctrl: ctrl}
	mock.recorder = &MockConnectionPostgresInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionPostgresInterface) EXPECT() *MockConnectionPostgresInterfaceMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockConnectionPostgresInterface) Begin(arg0 context.Context) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockConnectionPostgresInterfaceMockRecorder) Begin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockConnectionPostgresInterface)(nil).Begin), arg0)
}

// Exec mocks base method.
func (m *MockConnectionPostgresInterface) Exec(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockConnectionPostgresInterfaceMockRecorder) Exec(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockConnectionPostgresInterface)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockConnectionPostgresInterface) Query(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockConnectionPostgresInterfaceMockRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockConnectionPostgresInterface)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockConnectionPostgresInterface) QueryRow(arg0 context.Context, arg1 string, arg2 ...interface{}) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockConnectionPostgresInterfaceMockRecorder) QueryRow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockConnectionPostgresInterface)(nil).QueryRow), varargs...)
}

// MockTransactionInterface is a mock of TransactionInterface interface.
type MockTransactionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionInterfaceMockRecorder
}

// MockTransactionInterfaceMockRecorder is the mock recorder for MockTransactionInterface.
type MockTransactionInterfaceMockRecorder struct {
	mock *MockTransactionInterface
}

// NewMockTransactionInterface creates a new mock instance.
func NewMockTransactionInterface(ctrl *gomock.Controller) *MockTransactionInterface {
	mock := &MockTransactionInterface{ctrl: ctrl}
	mock.recorder = &MockTransactionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionInterface) EXPECT() *MockTransactionInterfaceMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockTransactionInterface) Begin(arg0 context.Context) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockTransactionInterfaceMockRecorder) Begin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockTransactionInterface)(nil).Begin), arg0)
}

// BeginFunc mocks base method.
func (m *MockTransactionInterface) BeginFunc(arg0 context.Context, arg1 func(pgx.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginFunc", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginFunc indicates an expected call of BeginFunc.
func (mr *MockTransactionInterfaceMockRecorder) BeginFunc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginFunc", reflect.TypeOf((*MockTransactionInterface)(nil).BeginFunc), arg0, arg1)
}

// Commit mocks base method.
func (m *MockTransactionInterface) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionInterfaceMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransactionInterface)(nil).Commit), arg0)
}

// Conn mocks base method.
func (m *MockTransactionInterface) Conn() *pgx.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn")
	ret0, _ := ret[0].(*pgx.Conn)
	return ret0
}

// Conn indicates an expected call of Conn.
func (mr *MockTransactionInterfaceMockRecorder) Conn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockTransactionInterface)(nil).Conn))
}

// CopyFrom mocks base method.
func (m *MockTransactionInterface) CopyFrom(arg0 context.Context, arg1 pgx.Identifier, arg2 []string, arg3 pgx.CopyFromSource) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFrom", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CopyFrom indicates an expected call of CopyFrom.
func (mr *MockTransactionInterfaceMockRecorder) CopyFrom(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFrom", reflect.TypeOf((*MockTransactionInterface)(nil).CopyFrom), arg0, arg1, arg2, arg3)
}

// Exec mocks base method.
func (m *MockTransactionInterface) Exec(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockTransactionInterfaceMockRecorder) Exec(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockTransactionInterface)(nil).Exec), varargs...)
}

// LargeObjects mocks base method.
func (m *MockTransactionInterface) LargeObjects() pgx.LargeObjects {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LargeObjects")
	ret0, _ := ret[0].(pgx.LargeObjects)
	return ret0
}

// LargeObjects indicates an expected call of LargeObjects.
func (mr *MockTransactionInterfaceMockRecorder) LargeObjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LargeObjects", reflect.TypeOf((*MockTransactionInterface)(nil).LargeObjects))
}

// Prepare mocks base method.
func (m *MockTransactionInterface) Prepare(arg0 context.Context, arg1, arg2 string) (*pgconn.StatementDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", arg0, arg1, arg2)
	ret0, _ := ret[0].(*pgconn.StatementDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare.
func (mr *MockTransactionInterfaceMockRecorder) Prepare(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockTransactionInterface)(nil).Prepare), arg0, arg1, arg2)
}

// Query mocks base method.
func (m *MockTransactionInterface) Query(arg0 context.Context, arg1 string, arg2 ...interface{}) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockTransactionInterfaceMockRecorder) Query(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockTransactionInterface)(nil).Query), varargs...)
}

// QueryFunc mocks base method.
func (m *MockTransactionInterface) QueryFunc(arg0 context.Context, arg1 string, arg2, arg3 []interface{}, arg4 func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFunc", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFunc indicates an expected call of QueryFunc.
func (mr *MockTransactionInterfaceMockRecorder) QueryFunc(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFunc", reflect.TypeOf((*MockTransactionInterface)(nil).QueryFunc), arg0, arg1, arg2, arg3, arg4)
}

// QueryRow mocks base method.
func (m *MockTransactionInterface) QueryRow(arg0 context.Context, arg1 string, arg2 ...interface{}) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockTransactionInterfaceMockRecorder) QueryRow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockTransactionInterface)(nil).QueryRow), varargs...)
}

// Rollback mocks base method.
func (m *MockTransactionInterface) Rollback(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionInterfaceMockRecorder) Rollback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransactionInterface)(nil).Rollback), arg0)
}

// SendBatch mocks base method.
func (m *MockTransactionInterface) SendBatch(arg0 context.Context, arg1 *pgx.Batch) pgx.BatchResults {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendBatch", arg0, arg1)
	ret0, _ := ret[0].(pgx.BatchResults)
	return ret0
}

// SendBatch indicates an expected call of SendBatch.
func (mr *MockTransactionInterfaceMockRecorder) SendBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatch", reflect.TypeOf((*MockTransactionInterface)(nil).SendBatch), arg0, arg1)
}

// MockConnectionRedisInterface is a mock of ConnectionRedisInterface interface.
type MockConnectionRedisInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionRedisInterfaceMockRecorder
}

// MockConnectionRedisInterfaceMockRecorder is the mock recorder for MockConnectionRedisInterface.
type MockConnectionRedisInterfaceMockRecorder struct {
	mock *MockConnectionRedisInterface
}

// NewMockConnectionRedisInterface creates a new mock instance.
func NewMockConnectionRedisInterface(ctrl *gomock.Controller) *MockConnectionRedisInterface {
	mock := &MockConnectionRedisInterface{ctrl: ctrl}
	mock.recorder = &MockConnectionRedisInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionRedisInterface) EXPECT() *MockConnectionRedisInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConnectionRedisInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionRedisInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnectionRedisInterface)(nil).Close))
}

// Do mocks base method.
func (m *MockConnectionRedisInterface) Do(arg0 string, arg1 ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockConnectionRedisInterfaceMockRecorder) Do(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockConnectionRedisInterface)(nil).Do), varargs...)
}

// Err mocks base method.
func (m *MockConnectionRedisInterface) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockConnectionRedisInterfaceMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockConnectionRedisInterface)(nil).Err))
}

// Flush mocks base method.
func (m *MockConnectionRedisInterface) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockConnectionRedisInterfaceMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockConnectionRedisInterface)(nil).Flush))
}

// Receive mocks base method.
func (m *MockConnectionRedisInterface) Receive() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockConnectionRedisInterfaceMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockConnectionRedisInterface)(nil).Receive))
}

// Send mocks base method.
func (m *MockConnectionRedisInterface) Send(arg0 string, arg1 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockConnectionRedisInterfaceMockRecorder) Send(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockConnectionRedisInterface)(nil).Send), varargs...)
}
