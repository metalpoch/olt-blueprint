package repository

import (
	"context"

	"github.com/metalpoch/olt-blueprint/common/entity"
	"github.com/metalpoch/olt-blueprint/common/model"
	"gorm.io/gorm"
)

type trafficRepository struct {
	db *gorm.DB
}

func NewTrafficRepository(db *gorm.DB) *trafficRepository {
	return &trafficRepository{db}
}

func (repo trafficRepository) Add(ctx context.Context, traffic *entity.Traffic) error {
	return repo.db.WithContext(ctx).Create(traffic).Error
}

func (repo trafficRepository) GetTrafficByInterface(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error) {
	var traffic []*entity.Traffic
	err := repo.db.WithContext(ctx).
		Where("interface_id = ? AND date BETWEEN ? AND ?", id, date.InitDate, date.EndDate).
		Order("date").
		Find(&traffic).
		Error
	return traffic, err
}

func (repo trafficRepository) GetTrafficByDevice(ctx context.Context, id uint, date *model.TranficRangeDate) ([]*entity.Traffic, error) {
	var traffic []*entity.Traffic
	err := repo.db.WithContext(ctx).
		Model(&entity.Traffic{}).
		Select("date, SUM(\"in\") AS \"in\", SUM(out) AS out, SUM(bandwidth) as bandwidth").
		Joins("JOIN interfaces ON interfaces.id = traffics.interface_id").
		Where("interfaces.device_id = ? AND traffics.date BETWEEN ? AND ?", id, date.InitDate, date.EndDate).
		Group("date").
		Order("date").
		Find(&traffic).
		Error
	return traffic, err
}