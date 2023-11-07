package model

import (
	"reflect"
	"time"
)

type Course struct {
	ID         string    `json:"id"`
	SemesterID string    `json:"semester_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  *uint     `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  *uint     `json:"updated_by"`
}

func (Course) TableName() string {
	return "course"
}

func (c Course) BuildUpdateFields() map[string]interface{} {
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
