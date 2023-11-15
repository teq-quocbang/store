package model

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID       `json:"id" yaml:"id"`
	Name        string          `json:"name" yaml:"name"`
	ProductType string          `json:"product_type" yaml:"product_type"`
	ProducerID  uuid.UUID       `json:"producer_id" yaml:"producer_id"`
	Price       decimal.Decimal `json:"price" yaml:"price"`
	CreatedAt   time.Time       `json:"created_at" yaml:"created_at"`
	CreatedBy   uuid.UUID       `json:"created_by" yaml:"created_by"`
	UpdatedAt   time.Time       `json:"updated_at" yaml:"updated_at"`
	UpdatedBy   uuid.UUID       `json:"updated_by" yaml:"updated_by"`
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
