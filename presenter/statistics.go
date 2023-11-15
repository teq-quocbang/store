package presenter

import (
	"time"

	"github.com/google/uuid"
)

type SoldProduct struct {
	ProductIDs []uuid.UUID `json:"product_ids"`
	SoldQty    int64       `json:"sold_qty"`
	SoldStart  time.Time   `json:"sold_start"`
	SoldEnd    time.Time   `json:"sold_end"`
}

type ListStatisticsSoldProductChartResponseWrapper struct {
	Sold []SoldProduct          `json:"sold"`
	Meta map[string]interface{} `json:"meta"`
}
