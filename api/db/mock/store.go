// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ot07/next-bazaar/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	db "github.com/ot07/next-bazaar/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CountProducts mocks base method.
func (m *MockStore) CountProducts(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProducts", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProducts indicates an expected call of CountProducts.
func (mr *MockStoreMockRecorder) CountProducts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProducts", reflect.TypeOf((*MockStore)(nil).CountProducts), arg0)
}

// CreateCartProduct mocks base method.
func (m *MockStore) CreateCartProduct(arg0 context.Context, arg1 db.CreateCartProductParams) (db.CartProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCartProduct", arg0, arg1)
	ret0, _ := ret[0].(db.CartProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCartProduct indicates an expected call of CreateCartProduct.
func (mr *MockStoreMockRecorder) CreateCartProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCartProduct", reflect.TypeOf((*MockStore)(nil).CreateCartProduct), arg0, arg1)
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteCategory mocks base method.
func (m *MockStore) DeleteCategory(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockStoreMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockStore)(nil).DeleteCategory), arg0, arg1)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), arg0, arg1)
}

// GetCartProductByUserIdAndProductId mocks base method.
func (m *MockStore) GetCartProductByUserIdAndProductId(arg0 context.Context, arg1 db.GetCartProductByUserIdAndProductIdParams) (db.CartProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartProductByUserIdAndProductId", arg0, arg1)
	ret0, _ := ret[0].(db.CartProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartProductByUserIdAndProductId indicates an expected call of GetCartProductByUserIdAndProductId.
func (mr *MockStoreMockRecorder) GetCartProductByUserIdAndProductId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartProductByUserIdAndProductId", reflect.TypeOf((*MockStore)(nil).GetCartProductByUserIdAndProductId), arg0, arg1)
}

// GetCartProductsByUserId mocks base method.
func (m *MockStore) GetCartProductsByUserId(arg0 context.Context, arg1 uuid.UUID) ([]db.CartProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartProductsByUserId", arg0, arg1)
	ret0, _ := ret[0].([]db.CartProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartProductsByUserId indicates an expected call of GetCartProductsByUserId.
func (mr *MockStoreMockRecorder) GetCartProductsByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartProductsByUserId", reflect.TypeOf((*MockStore)(nil).GetCartProductsByUserId), arg0, arg1)
}

// GetCategoriesByIDs mocks base method.
func (m *MockStore) GetCategoriesByIDs(arg0 context.Context, arg1 []uuid.UUID) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesByIDs", arg0, arg1)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoriesByIDs indicates an expected call of GetCategoriesByIDs.
func (mr *MockStoreMockRecorder) GetCategoriesByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesByIDs", reflect.TypeOf((*MockStore)(nil).GetCategoriesByIDs), arg0, arg1)
}

// GetCategory mocks base method.
func (m *MockStore) GetCategory(arg0 context.Context, arg1 uuid.UUID) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockStoreMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockStore)(nil).GetCategory), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStore) GetProduct(arg0 context.Context, arg1 uuid.UUID) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStore)(nil).GetProduct), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 uuid.UUID) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUsersByIDs mocks base method.
func (m *MockStore) GetUsersByIDs(arg0 context.Context, arg1 []uuid.UUID) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByIDs", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByIDs indicates an expected call of GetUsersByIDs.
func (mr *MockStoreMockRecorder) GetUsersByIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByIDs", reflect.TypeOf((*MockStore)(nil).GetUsersByIDs), arg0, arg1)
}

// ListCategories mocks base method.
func (m *MockStore) ListCategories(arg0 context.Context, arg1 db.ListCategoriesParams) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", arg0, arg1)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCategories indicates an expected call of ListCategories.
func (mr *MockStoreMockRecorder) ListCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockStore)(nil).ListCategories), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStore) ListProducts(arg0 context.Context, arg1 db.ListProductsParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStore)(nil).ListProducts), arg0, arg1)
}

// TruncateCartProductsTable mocks base method.
func (m *MockStore) TruncateCartProductsTable(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TruncateCartProductsTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TruncateCartProductsTable indicates an expected call of TruncateCartProductsTable.
func (mr *MockStoreMockRecorder) TruncateCartProductsTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateCartProductsTable", reflect.TypeOf((*MockStore)(nil).TruncateCartProductsTable), arg0)
}

// TruncateCategoriesTable mocks base method.
func (m *MockStore) TruncateCategoriesTable(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TruncateCategoriesTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TruncateCategoriesTable indicates an expected call of TruncateCategoriesTable.
func (mr *MockStoreMockRecorder) TruncateCategoriesTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateCategoriesTable", reflect.TypeOf((*MockStore)(nil).TruncateCategoriesTable), arg0)
}

// TruncateProductsTable mocks base method.
func (m *MockStore) TruncateProductsTable(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TruncateProductsTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TruncateProductsTable indicates an expected call of TruncateProductsTable.
func (mr *MockStoreMockRecorder) TruncateProductsTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateProductsTable", reflect.TypeOf((*MockStore)(nil).TruncateProductsTable), arg0)
}

// TruncateSessionsTable mocks base method.
func (m *MockStore) TruncateSessionsTable(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TruncateSessionsTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TruncateSessionsTable indicates an expected call of TruncateSessionsTable.
func (mr *MockStoreMockRecorder) TruncateSessionsTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateSessionsTable", reflect.TypeOf((*MockStore)(nil).TruncateSessionsTable), arg0)
}

// TruncateUsersTable mocks base method.
func (m *MockStore) TruncateUsersTable(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TruncateUsersTable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TruncateUsersTable indicates an expected call of TruncateUsersTable.
func (mr *MockStoreMockRecorder) TruncateUsersTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TruncateUsersTable", reflect.TypeOf((*MockStore)(nil).TruncateUsersTable), arg0)
}

// UpdateCartProduct mocks base method.
func (m *MockStore) UpdateCartProduct(arg0 context.Context, arg1 db.UpdateCartProductParams) (db.CartProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartProduct", arg0, arg1)
	ret0, _ := ret[0].(db.CartProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartProduct indicates an expected call of UpdateCartProduct.
func (mr *MockStoreMockRecorder) UpdateCartProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartProduct", reflect.TypeOf((*MockStore)(nil).UpdateCartProduct), arg0, arg1)
}
