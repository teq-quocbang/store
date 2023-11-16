package presenter

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type SoldProduct struct {
	ProductIDs []uuid.UUID     `json:"product_ids,omitempty"`
	SoldQty    int64           `json:"sold_qty,omitempty"`
	TotalPrice decimal.Decimal `json:"total_price,omitempty"`
	SoldStart  time.Time       `json:"sold_start,omitempty"`
	SoldEnd    time.Time       `json:"sold_end,omitempty"`
}

type ListStatisticsSoldProductChartResponseWrapper struct {
	Sold []SoldProduct          `json:"sold"`
	Meta map[string]interface{} `json:"meta"`
}
