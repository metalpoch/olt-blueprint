package usecase

import "github.com/metalpoch/olt-blueprint/common/model"

type DeviceUsecase interface {
	Add(device *model.AddDevice) error
	Check(device *model.Device) error
	GetAll() ([]*model.Device, error)
	GetDeviceWithOIDRows() ([]*model.DeviceWithOID, error)
	Update(id uint, device *model.AddDevice) error
	Delete(id uint) error
}

type InterfaceUsecase interface {
	Upsert(element *model.Interface) error
	GetAll() ([]*model.Interface, error)
	GetAllByDevice(id uint) ([]*model.Interface, error)
}

type TrafficUsecase interface {
	Add(measurement *model.Traffic) error
	GetTrafficByInterface(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
	GetTrafficByDevice(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
}