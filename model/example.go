package model

import (
	"time"

	"gorm.io/gorm"
)

type Example struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	CreatedBy int64           `json:"created_by"`
	UpdatedBy *int64          `json:"updated_by"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}
