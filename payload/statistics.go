package payload

import "github.com/teq-quocbang/store/codetype"

type GetProductSoldChartRequest struct {
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	ProductType string `json:"product_type" query:"product_type"`
	SortBy      codetype.SortType
	OrderBy     string `json:"order_by,omitempty" query:"order_by"`
}

func (p GetProductSoldChartRequest) IsNeedToFilter() bool {
	return p.ProductType != ""
}
