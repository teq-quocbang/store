package storage

import (
	"context"

	"github.com/google/uuid"
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

func (r *pgRepository) GetInventoryQty(ctx context.Context, productID uuid.UUID) (int, error) {
	var NResult struct {
		N int
	}
	err := r.getDB(ctx).Model(&model.Storage{}).Where("product_id = ?", productID).Select("sum(inventory_qty) as N").Scan(&NResult).Error
	if err != nil {
		return 0, err
	}
	return NResult.N, nil
}
