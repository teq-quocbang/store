package checkout

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

func (r *pgRepository) UpsertCart(ctx context.Context, req *model.Cart) error {
	return r.getDB(ctx).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{ // conflict with unique field
			"qty": gorm.Expr("qty + ?", req.Qty),
		}),
	}).Create(&req).Error
}

func (r *pgRepository) GetCartByConstraint(ctx context.Context, accountID uuid.UUID, productID uuid.UUID) (model.Cart, error) {
	cart := model.Cart{}
	err := r.getDB(ctx).Where("account_id = ? and product_id = ?", accountID, productID).Take(&cart).Error
	if err != nil {
		return model.Cart{}, err
	}
	return cart, nil
}

func (r *pgRepository) GetListCart(ctx context.Context, accountID uuid.UUID) ([]model.Cart, error) {
	carts := []model.Cart{}
	err := r.getDB(ctx).Where("account_id = ?", accountID).Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}
