package model

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ProductType string    `json:"product_type"`
	ProducerID  uuid.UUID `json:"producer_id"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   uuid.UUID `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
}

func (Product) TableName() string {
	return "product"
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}

func (c Product) BuildUpdateFields() map[string]interface{} {
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
