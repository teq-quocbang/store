package storage

import (
	"context"

	"github.com/teq-quocbang/store/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

func (r *pgRepository) UpsertStorage(ctx context.Context, req *model.Storage) error {
	return r.getDB(ctx).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{ // conflict with unique field
			"inventory_qty": gorm.Expr("inventory_qty + ?", req.InventoryQty),
		}),
	}).Create(&req).Error
}

func (r pgRepository) GetListStorageByLocat(ctx context.Context, locat string) ([]model.Storage, error) {
	storages := []model.Storage{}
	err := r.getDB(ctx).Where("locat = ?", locat).Find(&storages).Error
	if err != nil {
		return nil, err
	}
	return storages, nil
}
