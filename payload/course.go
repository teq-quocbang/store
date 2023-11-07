package payload

import "github.com/go-playground/validator/v10"

type CreateCourseRequest struct {
	ID         string `json:"id" validate:"required"`
	SemesterID string `json:"semester_id" validate:"required"`
}

type ListCourseBySemesterRequest struct {
	SemesterID string `json:"semester_id" validate:"required"`
}

func (c *CreateCourseRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *ListCourseBySemesterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

type UpdateCourseRequest struct {
	ID         string `json:"id" validate:"required"`
	SemesterID string `json:"semester_id" validate:"required"`
}

func (c *UpdateCourseRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
