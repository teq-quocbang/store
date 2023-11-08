package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Producer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uuid.UUID `json:"updated_by"`
}

func (Producer) TableName() string {
	return "producer"
}

func (p *Producer) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
