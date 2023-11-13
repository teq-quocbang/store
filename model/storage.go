package model

import (
	"time"

	"github.com/google/uuid"
)

type Storage struct {
	Locat        string    `json:"locat"`
	ProductID    uuid.UUID `json:"product_id"`
	InventoryQty int64     `json:"inventory_qty"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    uuid.UUID `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
}

func (Storage) TableName() string {
	return "storage"
}
