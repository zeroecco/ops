// Code generated by MockGen. DO NOT EDIT.
// Source: upcloud/provider.go

// Package mock_upcloud is a generated GoMock package.
package mock_upcloud

import (
	upcloud "github.com/UpCloudLtd/upcloud-go-api/upcloud"
	request "github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetServerConfigurations mocks base method
func (m *MockService) GetServerConfigurations() (*upcloud.ServerConfigurations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerConfigurations")
	ret0, _ := ret[0].(*upcloud.ServerConfigurations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerConfigurations indicates an expected call of GetServerConfigurations
func (mr *MockServiceMockRecorder) GetServerConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerConfigurations", reflect.TypeOf((*MockService)(nil).GetServerConfigurations))
}

// GetServers mocks base method
func (m *MockService) GetServers() (*upcloud.Servers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServers")
	ret0, _ := ret[0].(*upcloud.Servers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServers indicates an expected call of GetServers
func (mr *MockServiceMockRecorder) GetServers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServers", reflect.TypeOf((*MockService)(nil).GetServers))
}

// GetServerDetails mocks base method
func (m *MockService) GetServerDetails(r *request.GetServerDetailsRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerDetails", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerDetails indicates an expected call of GetServerDetails
func (mr *MockServiceMockRecorder) GetServerDetails(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerDetails", reflect.TypeOf((*MockService)(nil).GetServerDetails), r)
}

// CreateServer mocks base method
func (m *MockService) CreateServer(r *request.CreateServerRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateServer indicates an expected call of CreateServer
func (mr *MockServiceMockRecorder) CreateServer(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockService)(nil).CreateServer), r)
}

// WaitForServerState mocks base method
func (m *MockService) WaitForServerState(r *request.WaitForServerStateRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForServerState", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForServerState indicates an expected call of WaitForServerState
func (mr *MockServiceMockRecorder) WaitForServerState(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForServerState", reflect.TypeOf((*MockService)(nil).WaitForServerState), r)
}

// StartServer mocks base method
func (m *MockService) StartServer(r *request.StartServerRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartServer", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartServer indicates an expected call of StartServer
func (mr *MockServiceMockRecorder) StartServer(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartServer", reflect.TypeOf((*MockService)(nil).StartServer), r)
}

// StopServer mocks base method
func (m *MockService) StopServer(r *request.StopServerRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopServer", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopServer indicates an expected call of StopServer
func (mr *MockServiceMockRecorder) StopServer(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopServer", reflect.TypeOf((*MockService)(nil).StopServer), r)
}

// RestartServer mocks base method
func (m *MockService) RestartServer(r *request.RestartServerRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartServer", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestartServer indicates an expected call of RestartServer
func (mr *MockServiceMockRecorder) RestartServer(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartServer", reflect.TypeOf((*MockService)(nil).RestartServer), r)
}

// ModifyServer mocks base method
func (m *MockService) ModifyServer(r *request.ModifyServerRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyServer", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyServer indicates an expected call of ModifyServer
func (mr *MockServiceMockRecorder) ModifyServer(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyServer", reflect.TypeOf((*MockService)(nil).ModifyServer), r)
}

// DeleteServer mocks base method
func (m *MockService) DeleteServer(r *request.DeleteServerRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServer", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServer indicates an expected call of DeleteServer
func (mr *MockServiceMockRecorder) DeleteServer(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServer", reflect.TypeOf((*MockService)(nil).DeleteServer), r)
}

// DeleteServerAndStorages mocks base method
func (m *MockService) DeleteServerAndStorages(r *request.DeleteServerAndStoragesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServerAndStorages", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServerAndStorages indicates an expected call of DeleteServerAndStorages
func (mr *MockServiceMockRecorder) DeleteServerAndStorages(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServerAndStorages", reflect.TypeOf((*MockService)(nil).DeleteServerAndStorages), r)
}

// GetStorages mocks base method
func (m *MockService) GetStorages(r *request.GetStoragesRequest) (*upcloud.Storages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorages", r)
	ret0, _ := ret[0].(*upcloud.Storages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorages indicates an expected call of GetStorages
func (mr *MockServiceMockRecorder) GetStorages(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorages", reflect.TypeOf((*MockService)(nil).GetStorages), r)
}

// GetStorageDetails mocks base method
func (m *MockService) GetStorageDetails(r *request.GetStorageDetailsRequest) (*upcloud.StorageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageDetails", r)
	ret0, _ := ret[0].(*upcloud.StorageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageDetails indicates an expected call of GetStorageDetails
func (mr *MockServiceMockRecorder) GetStorageDetails(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageDetails", reflect.TypeOf((*MockService)(nil).GetStorageDetails), r)
}

// CreateStorage mocks base method
func (m *MockService) CreateStorage(r *request.CreateStorageRequest) (*upcloud.StorageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStorage", r)
	ret0, _ := ret[0].(*upcloud.StorageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStorage indicates an expected call of CreateStorage
func (mr *MockServiceMockRecorder) CreateStorage(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStorage", reflect.TypeOf((*MockService)(nil).CreateStorage), r)
}

// ModifyStorage mocks base method
func (m *MockService) ModifyStorage(r *request.ModifyStorageRequest) (*upcloud.StorageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyStorage", r)
	ret0, _ := ret[0].(*upcloud.StorageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyStorage indicates an expected call of ModifyStorage
func (mr *MockServiceMockRecorder) ModifyStorage(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyStorage", reflect.TypeOf((*MockService)(nil).ModifyStorage), r)
}

// AttachStorage mocks base method
func (m *MockService) AttachStorage(r *request.AttachStorageRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachStorage", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachStorage indicates an expected call of AttachStorage
func (mr *MockServiceMockRecorder) AttachStorage(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachStorage", reflect.TypeOf((*MockService)(nil).AttachStorage), r)
}

// DetachStorage mocks base method
func (m *MockService) DetachStorage(r *request.DetachStorageRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachStorage", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachStorage indicates an expected call of DetachStorage
func (mr *MockServiceMockRecorder) DetachStorage(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachStorage", reflect.TypeOf((*MockService)(nil).DetachStorage), r)
}

// CloneStorage mocks base method
func (m *MockService) CloneStorage(r *request.CloneStorageRequest) (*upcloud.StorageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneStorage", r)
	ret0, _ := ret[0].(*upcloud.StorageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneStorage indicates an expected call of CloneStorage
func (mr *MockServiceMockRecorder) CloneStorage(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneStorage", reflect.TypeOf((*MockService)(nil).CloneStorage), r)
}

// TemplatizeStorage mocks base method
func (m *MockService) TemplatizeStorage(r *request.TemplatizeStorageRequest) (*upcloud.StorageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemplatizeStorage", r)
	ret0, _ := ret[0].(*upcloud.StorageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TemplatizeStorage indicates an expected call of TemplatizeStorage
func (mr *MockServiceMockRecorder) TemplatizeStorage(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemplatizeStorage", reflect.TypeOf((*MockService)(nil).TemplatizeStorage), r)
}

// WaitForStorageState mocks base method
func (m *MockService) WaitForStorageState(r *request.WaitForStorageStateRequest) (*upcloud.StorageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForStorageState", r)
	ret0, _ := ret[0].(*upcloud.StorageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForStorageState indicates an expected call of WaitForStorageState
func (mr *MockServiceMockRecorder) WaitForStorageState(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForStorageState", reflect.TypeOf((*MockService)(nil).WaitForStorageState), r)
}

// LoadCDROM mocks base method
func (m *MockService) LoadCDROM(r *request.LoadCDROMRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadCDROM", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadCDROM indicates an expected call of LoadCDROM
func (mr *MockServiceMockRecorder) LoadCDROM(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadCDROM", reflect.TypeOf((*MockService)(nil).LoadCDROM), r)
}

// EjectCDROM mocks base method
func (m *MockService) EjectCDROM(r *request.EjectCDROMRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EjectCDROM", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EjectCDROM indicates an expected call of EjectCDROM
func (mr *MockServiceMockRecorder) EjectCDROM(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EjectCDROM", reflect.TypeOf((*MockService)(nil).EjectCDROM), r)
}

// CreateBackup mocks base method
func (m *MockService) CreateBackup(r *request.CreateBackupRequest) (*upcloud.StorageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBackup", r)
	ret0, _ := ret[0].(*upcloud.StorageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBackup indicates an expected call of CreateBackup
func (mr *MockServiceMockRecorder) CreateBackup(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackup", reflect.TypeOf((*MockService)(nil).CreateBackup), r)
}

// RestoreBackup mocks base method
func (m *MockService) RestoreBackup(r *request.RestoreBackupRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreBackup", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreBackup indicates an expected call of RestoreBackup
func (mr *MockServiceMockRecorder) RestoreBackup(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreBackup", reflect.TypeOf((*MockService)(nil).RestoreBackup), r)
}

// CreateStorageImport mocks base method
func (m *MockService) CreateStorageImport(r *request.CreateStorageImportRequest) (*upcloud.StorageImportDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStorageImport", r)
	ret0, _ := ret[0].(*upcloud.StorageImportDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStorageImport indicates an expected call of CreateStorageImport
func (mr *MockServiceMockRecorder) CreateStorageImport(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStorageImport", reflect.TypeOf((*MockService)(nil).CreateStorageImport), r)
}

// GetStorageImportDetails mocks base method
func (m *MockService) GetStorageImportDetails(r *request.GetStorageImportDetailsRequest) (*upcloud.StorageImportDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageImportDetails", r)
	ret0, _ := ret[0].(*upcloud.StorageImportDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageImportDetails indicates an expected call of GetStorageImportDetails
func (mr *MockServiceMockRecorder) GetStorageImportDetails(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageImportDetails", reflect.TypeOf((*MockService)(nil).GetStorageImportDetails), r)
}

// WaitForStorageImportCompletion mocks base method
func (m *MockService) WaitForStorageImportCompletion(r *request.WaitForStorageImportCompletionRequest) (*upcloud.StorageImportDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForStorageImportCompletion", r)
	ret0, _ := ret[0].(*upcloud.StorageImportDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForStorageImportCompletion indicates an expected call of WaitForStorageImportCompletion
func (mr *MockServiceMockRecorder) WaitForStorageImportCompletion(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForStorageImportCompletion", reflect.TypeOf((*MockService)(nil).WaitForStorageImportCompletion), r)
}

// DeleteStorage mocks base method
func (m *MockService) DeleteStorage(arg0 *request.DeleteStorageRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStorage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStorage indicates an expected call of DeleteStorage
func (mr *MockServiceMockRecorder) DeleteStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStorage", reflect.TypeOf((*MockService)(nil).DeleteStorage), arg0)
}

// GetTags mocks base method
func (m *MockService) GetTags() (*upcloud.Tags, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags")
	ret0, _ := ret[0].(*upcloud.Tags)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags
func (mr *MockServiceMockRecorder) GetTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockService)(nil).GetTags))
}

// CreateTag mocks base method
func (m *MockService) CreateTag(r *request.CreateTagRequest) (*upcloud.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTag", r)
	ret0, _ := ret[0].(*upcloud.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTag indicates an expected call of CreateTag
func (mr *MockServiceMockRecorder) CreateTag(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockService)(nil).CreateTag), r)
}

// ModifyTag mocks base method
func (m *MockService) ModifyTag(r *request.ModifyTagRequest) (*upcloud.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyTag", r)
	ret0, _ := ret[0].(*upcloud.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyTag indicates an expected call of ModifyTag
func (mr *MockServiceMockRecorder) ModifyTag(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyTag", reflect.TypeOf((*MockService)(nil).ModifyTag), r)
}

// DeleteTag mocks base method
func (m *MockService) DeleteTag(r *request.DeleteTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTag", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTag indicates an expected call of DeleteTag
func (mr *MockServiceMockRecorder) DeleteTag(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockService)(nil).DeleteTag), r)
}

// TagServer mocks base method
func (m *MockService) TagServer(r *request.TagServerRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagServer", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagServer indicates an expected call of TagServer
func (mr *MockServiceMockRecorder) TagServer(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagServer", reflect.TypeOf((*MockService)(nil).TagServer), r)
}

// UntagServer mocks base method
func (m *MockService) UntagServer(r *request.UntagServerRequest) (*upcloud.ServerDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagServer", r)
	ret0, _ := ret[0].(*upcloud.ServerDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagServer indicates an expected call of UntagServer
func (mr *MockServiceMockRecorder) UntagServer(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagServer", reflect.TypeOf((*MockService)(nil).UntagServer), r)
}

// GetAccount mocks base method
func (m *MockService) GetAccount() (*upcloud.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount")
	ret0, _ := ret[0].(*upcloud.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockServiceMockRecorder) GetAccount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockService)(nil).GetAccount))
}
