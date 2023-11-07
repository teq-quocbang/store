package presenter

import "github.com/teq-quocbang/store/model"

type RegisterResponseCustom struct {
	AccountID uint           `json:"account_id"`
	Semester  model.Semester `json:"semester"`
	Class     model.Class    `json:"class"`
	Course    model.Course   `json:"course"`
}

type RegisterResponseWrapper struct {
	Register RegisterResponseCustom `json:"register"`
}

type ListRegisterResponseWrapper struct {
	Register []RegisterResponseCustom `json:"register"`
	Meta     interface{}              `json:"meta"`
}
