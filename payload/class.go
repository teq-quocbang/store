package payload

import "github.com/go-playground/validator/v10"

type CreateClassRequest struct {
	ID         string `json:"id" validate:"required"`
	CourseID   string `json:"course_id" validate:"required"`
	SemesterID string `json:"semester_id" validate:"required"`
	StartTime  string `json:"start_time" validate:"required"`
	EndTime    string `json:"end_time" validate:"required"`
	Credits    int    `json:"credits" validate:"required,min=1"`
	MaxSlot    int    `json:"max_slot" validate:"required,min=1"`
}

func (c *CreateClassRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

type ListClassBySemesterRequest struct {
	SemesterID string `json:"semester_id" validate:"required"`
}

func (c *ListClassBySemesterRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

type UpdateClassRequest struct {
	ID         string `json:"id"`
	CourseID   string `json:"course_id"`
	SemesterID string `json:"semester_id"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Credits    int    `json:"credits" validate:"min=0"`
	MaxSlot    int    `json:"max_slot" validate:"min=0"`
}

func (c *UpdateClassRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
