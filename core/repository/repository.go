package repository

import (
	"context"

	"github.com/metalpoch/olt-blueprint/common/entity"
	"github.com/metalpoch/olt-blueprint/common/model"
)

type InfoRepository interface {
	GetDevice(ctx context.Context, id uint) (*entity.Device, error)
	GetDeviceByIP(ctx context.Context, ip string) (*entity.Device, error)
	GetDeviceBySysname(ctx context.Context, sysname string) (*entity.Device, error)
	GetAllDevice(ctx context.Context) ([]*entity.Device, error)
	GetDeviceByState(ctx context.Context, state string) ([]*entity.Device, error)
	GetDeviceByCounty(ctx context.Context, state, county string) ([]*entity.Device, error)
	GetDeviceByMunicipality(ctx context.Context, state, county, municipality string) ([]*entity.Device, error)
	GetInterface(ctx context.Context, id uint) (*entity.Interface, error)
	GetInterfacesByDevice(ctx context.Context, id uint) ([]*entity.Interface, error)
	GetInterfacesByDeviceAndPorts(ctx context.Context, id uint, pattern string) ([]*entity.Interface, error)
	GetLocationStates(ctx context.Context) ([]*string, error)
	GetLocationCounties(ctx context.Context, state string) ([]*string, error)
	GetLocationMunicipalities(ctx context.Context, state, county string) ([]*string, error)
	GetLocation(ctx context.Context, id uint) (*entity.Location, error)
	GetFat(ctx context.Context, id uint) (*entity.Fat, error)
	GetODN(ctx context.Context, odn string) ([]*entity.FatInterface, error)
	GetODNStates(ctx context.Context, state string) ([]*string, error)
	GetODNStatesContries(ctx context.Context, state, country string) ([]*string, error)
	GetODNStatesContriesMunicipality(ctx context.Context, state, country, municipality string) ([]*string, error)
	GetODNDevice(ctx context.Context, id uint) ([]*string, error)
	GetODNDevicePort(ctx context.Context, id uint, pattern string) ([]*string, error)
	GetAllODN(ctx context.Context) ([]*string, error)
	GetDevicesByOND(ctx context.Context, odn string) ([]*uint, error)
}

type TrafficRepository interface {
	GetTrafficByInterface(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error)
	GetTrafficByDevice(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error)
	GetTrafficByFat(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error)
	GetTrafficByLocationID(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error)
	GetTrafficByState(ctx context.Context, state string, date *model.TranficRangeDate) ([]*entity.Traffic, error)
	GetTrafficByCounty(ctx context.Context, state, county string, date *model.TranficRangeDate) ([]*entity.Traffic, error)
	GetTrafficByMunicipality(ctx context.Context, state, county, municipality string, date *model.TranficRangeDate) ([]*entity.Traffic, error)
	GetTrafficByODN(ctx context.Context, odn string, date *model.TranficRangeDate) ([]*entity.Traffic, error)
	GetTotalTrafficByState(ctx context.Context, ids []uint, month string) (*model.TrafficState, error)
	GetTotalTrafficByODN(ctx context.Context, ids []uint, month string) (*model.TrafficODN, error)
}
