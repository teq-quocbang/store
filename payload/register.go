package payload

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/teq-quocbang/store/codetype"
)

type CreateRegisterRequest struct {
	SemesterID string `json:"semester_id" validate:"required"`
	ClassID    string `json:"class_id" validate:"required"`
	CourseID   string `json:"course_id" validate:"required"`
}

func (r *CreateRegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type ListRegisterInformationRequest struct {
	SemesterID string `json:"semester_id" validate:"required"`
	codetype.Paginator
	SortBy  codetype.SortType `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy string            `json:"order_by,omitempty" query:"order_by"`
}

func (s *ListRegisterInformationRequest) Format() {
	s.Paginator.Format()
	s.SortBy.Format()
	s.OrderBy = strings.ToLower(strings.TrimSpace(s.OrderBy))

	for i := range orderByExample {
		if s.OrderBy == orderByExample[i] {
			return
		}
	}

	s.OrderBy = ""
}

func (s *ListRegisterInformationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

type ListRegisteredHistories struct {
	SemesterID string `json:"semester_id"`
	codetype.Paginator
	SortBy  codetype.SortType `json:"sort_by,omitempty" query:"sort_by"`
	OrderBy string            `json:"order_by,omitempty" query:"order_by"`
}

func (s *ListRegisteredHistories) Format() {
	s.Paginator.Format()
	s.SortBy.Format()
	s.OrderBy = strings.ToLower(strings.TrimSpace(s.OrderBy))

	for i := range orderByExample {
		if s.OrderBy == orderByExample[i] {
			return
		}
	}

	s.OrderBy = ""
}

func (r *ListRegisteredHistories) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type UnRegisterRequest struct {
	SemesterID string `json:"semester_id" validate:"required"`
	ClassID    string `json:"class_id" validate:"required"`
	CourseID   string `json:"course_id" validate:"required"`
}

func (u *UnRegisterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
