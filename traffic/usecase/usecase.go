package usecase

import "github.com/metalpoch/olt-blueprint/common/model"

type FeedUsecase interface {
	GetDevice(id uint) (*model.Device, error)
	GetAllDevice() ([]*model.DeviceLite, error)
	GetInterface(id uint) (*model.Interface, error)
	GetInterfacesByDevice(id uint) ([]*model.InterfaceLite, error)
	GetLocationStates() ([]*string, error)
	GetLocationCounties(state string) ([]*string, error)
	GetLocationMunicipalities(state, county string) ([]*string, error)
}

type TrafficUsecase interface {
	GetTrafficByInterface(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
	GetTrafficByDevice(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
	GetTrafficByFat(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
	GetTrafficByLocationID(id uint, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
	GetTrafficByState(state string, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
	GetTrafficByCounty(state, county string, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
	GetTrafficByMunicipality(state, county, municipality string, date *model.TranficRangeDate) ([]*model.TrafficResponse, error)
}
