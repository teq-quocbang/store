package presenter

import "github.com/teq-quocbang/store/model"

type ListStatisticsResponseWrapper struct {
	CustomerOrder []model.CustomerOrder  `json:"customer_order"`
	Meta          map[string]interface{} `json:"meta"`
}
