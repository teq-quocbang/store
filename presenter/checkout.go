package presenter

import "github.com/teq-quocbang/store/model"

type CartResponseWrapper struct {
	Cart *model.Cart `json:"product"`
}

type ListCartResponseWrapper struct {
	Cart []model.Cart `json:"product" yaml:"product"`
	Meta interface{}  `json:"meta"`
}

type CustomerOrderResponseWrapper struct {
	CustomerOrder *model.CustomerOrder `json:"product"`
}

type ListCustomerOrderResponseWrapper struct {
	CustomerOrder []model.CustomerOrder `json:"product" yaml:"product"`
	Meta          interface{}           `json:"meta"`
}
