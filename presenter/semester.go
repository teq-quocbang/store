package presenter

import "github.com/teq-quocbang/store/model"

type SemesterResponseWrapper struct {
	Semester model.Semester `json:"semester"`
}

type ListSemesterResponseWrapper struct {
	Semester []model.Semester `json:"semester"`
	Meta     interface{}      `json:"meta"`
}
