package checkout

import (
	"context"
	"fmt"
	"time"

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

func (r *pgRepository) RemoveFromCart(ctx context.Context, accountID uuid.UUID, productID uuid.UUID, qty int64) error {
	return r.getDB(ctx).Model(&model.Cart{}).Where("account_id = ? and product_id = ?", accountID, productID).Update("qty", gorm.Expr("qty - ?", qty)).Error
}

func (r *pgRepository) CreateCustomerOrder(ctx context.Context, cdr *model.CustomerOrder) error {
	tx := r.getDB(ctx).Begin()

	// create customer order
	err := tx.Create(&cdr).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// decrease qty from storage
	storage := model.Storage{}
	err = tx.Model(&storage).Where("product_id = ?", cdr.ProductID).Update("inventory_qty", gorm.Expr("inventory_qty - ?", cdr.SoldQty)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// check if request qty is out of storage
	if storage.InventoryQty < 0 {
		tx.Rollback()
		return fmt.Errorf("order qty is out of inventory qty")
	}

	// delete from cart
	err = tx.Model(&model.Cart{}).Where("account_id = ? and product_id = ?", cdr.AccountID, cdr.ProductID).Delete("1=1").Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *pgRepository) GetListOrdered(
	ctx context.Context,
	accountID uuid.UUID,
	startTime time.Time,
	endTime time.Time,
	order []string) ([]model.CustomerOrder, error) {
	var (
		db   = r.getDB(ctx).Model(&model.CustomerOrder{})
		cdrs []model.CustomerOrder
	)

	for i := range order {
		db = db.Order(order[i])
	}

	err := db.Where("account_id = ? and created_at >= ? and created_at <= ?", accountID, startTime, endTime).Find(&cdrs).Error
	if err != nil {
		return nil, err
	}

	return cdrs, nil
}
