package checkout

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

func (r *pgRepository) UpsertCart(ctx context.Context, req *model.Cart) error {
	return r.getDB(ctx).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{ // conflict with unique field
			"qty": gorm.Expr("qty + ?", req.Qty),
		}),
	}).Create(&req).Error
}
