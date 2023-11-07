package model

import (
	"reflect"
	"time"
)

type Class struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	SemesterID  string    `json:"semester_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Credits     uint      `json:"creates"`
	MaxSlot     uint      `json:"max_slot"`
	CurrentSlot uint      `json:"current_slot"`
	CanCancel   bool      `json:"can_cancel"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   *uint     `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   *uint     `json:"updated_by"`
}

func (Class) TableName() string {
	return "class"
}

func (c Class) BuildUpdateFields() map[string]interface{} {
	values := reflect.ValueOf(c)
	result := make(map[string]interface{}, values.NumField())

	for i := 0; i < values.NumField(); i++ {
		filed := values.Field(i)
		fieldName := values.Type().Field(i).Name

		if !filed.IsZero() {
			result[fieldName] = filed.Interface()
		}
	}

	return result
}
