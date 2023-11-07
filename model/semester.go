package model

import (
	"reflect"
	"time"
)

type Semester struct {
	ID                string    `json:"id"`
	MinCredits        int       `json:"min_credits"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	RegisterStartAt   time.Time `json:"register_start_at"`
	RegisterExpiresAt time.Time `json:"register_expires_at"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         *uint     `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         *uint     `json:"updated_by"`
}

func (Semester) TableName() string {
	return "semester"
}

func (s Semester) BuildUpdateFields() map[string]interface{} {
	values := reflect.ValueOf(s)
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
