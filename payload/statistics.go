package payload

import "github.com/teq-quocbang/store/codetype"

type TimeDuration string

const (
	DEFAULT_TIME_DURATION TimeDuration = "day"
	WEEK                  TimeDuration = "week"
	Month                 TimeDuration = "month"
)

type GetChartRequest struct {
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	ProductType  string `json:"product_type" query:"product_type"`
	SortBy       codetype.SortType
	OrderBy      string `json:"order_by,omitempty" query:"order_by"`
	TimeDuration string `json:"time_duration" query:"time_duration"`
}

func (p GetChartRequest) IsNeedToFilter() bool {
	return p.ProductType != ""
}

func (p *GetChartRequest) Format() {
	if p.TimeDuration == "" {
		p.TimeDuration = string(DEFAULT_TIME_DURATION)
	}
}
