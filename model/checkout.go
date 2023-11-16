package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Cart struct {
	AccountID uuid.UUID `json:"account_id"`
	ProductID uuid.UUID `json:"product_id"`
	Qty       int64     `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
}

func (Cart) TableName() string {
	return "cart"
}

type CustomerOrder struct {
	AccountID uuid.UUID `json:"account_id"`
	ProductID uuid.UUID `json:"product_id"`
	// price of per product
	PriceOfPer decimal.Decimal `json:"price_of_per"`
	SoldQty    int64           `json:"sold_qty"`
	CreatedAt  time.Time       `json:"created_at"`
}

func (CustomerOrder) TableName() string {
	return "customer_order"
}
