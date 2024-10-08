package repository

import (
	"context"

	"github.com/metalpoch/olt-blueprint/common/entity"
)

type FeedRepository interface {
	GetDevice(ctx context.Context, id uint) (*entity.Device, error)
	GetAllDevice(ctx context.Context) ([]*entity.Device, error)
	GetInterface(ctx context.Context, id uint) (*entity.Interface, error)
	GetInterfacesByDevice(ctx context.Context, id uint) ([]*entity.Interface, error)
}
