package model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID           uuid.UUID `json:"ID,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	HashPassword []byte    `json:"hash_password,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

func (Account) TableName() string {
	return "account"
}
