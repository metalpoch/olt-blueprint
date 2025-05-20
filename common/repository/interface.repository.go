package repository

import (
	"context"

	"github.com/metalpoch/olt-blueprint/common/constants"
	"github.com/metalpoch/olt-blueprint/common/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InterfaceRepository interface {
	Get(ctx context.Context, id uint) (*entity.Interface, error)
	Upsert(ctx context.Context, element *entity.Interface) (uint64, error)
	GetAll(ctx context.Context) ([]*entity.Interface, error)
	GetAllByDevice(ctx context.Context, id uint64) ([]*entity.Interface, error)
	GetAllByDeviceAndIfindex(ctx context.Context, deviceID, idx uint64) (*entity.Interface, error)
}

type interfaceRepository struct {
	db *gorm.DB
}

func NewInterfaceRepository(db *gorm.DB) *interfaceRepository {
	return &interfaceRepository{db}
}

func (repo interfaceRepository) Upsert(ctx context.Context, element *entity.Interface) (uint64, error) {
	columns := []clause.Column{
		{Name: constants.INTERFACE_COLUMN_IF_INDEX},
		{Name: constants.INTERFACE_COLUMN_DEVICE_ID},
	}
	doUpdates := []string{
		constants.INTERFACE_COLUMN_IF_NAME,
		constants.INTERFACE_COLUMN_IF_DESCR,
		constants.INTERFACE_COLUMN_IF_ALIAS,
		constants.GLOBAL_COLUMN_UPDATED_AT,
	}

	err := repo.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   columns,
		DoUpdates: clause.AssignmentColumns(doUpdates),
	}).Create(element).Error

	return element.ID, err
}

func (repo interfaceRepository) Get(ctx context.Context, id uint) (*entity.Interface, error) {
	element := new(entity.Interface)
	err := repo.db.WithContext(ctx).First(element, id).Error
	return element, err
}

func (repo interfaceRepository) GetAll(ctx context.Context) ([]*entity.Interface, error) {
	var interfaces []*entity.Interface
	err := repo.db.WithContext(ctx).Find(&interfaces).Error
	return interfaces, err
}

func (repo interfaceRepository) GetAllByDevice(ctx context.Context, id uint64) ([]*entity.Interface, error) {
	var interfaces []*entity.Interface
	err := repo.db.WithContext(ctx).Where("device_id = ?", id).Find(&interfaces).Error
	return interfaces, err
}

func (repo interfaceRepository) GetAllByDeviceAndIfindex(ctx context.Context, deviceID, idx uint64) (*entity.Interface, error) {
	pon := new(entity.Interface)
	pon.DeviceID = deviceID
	pon.IfIndex = idx
	err := repo.db.WithContext(ctx).First(pon, idx).Error
	return pon, err
}
