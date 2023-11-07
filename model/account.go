package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID           uint      `json:"ID,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	HashPassword []byte    `json:"hash_password,omitempty"`
	IsVerified   bool      `json:"is_verified,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

func (Account) TableName() string {
	return "account"
}

type AccountVerify struct {
	AccountID  uint      `json:"account_id"`
	SecretCode uuid.UUID `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
}

func (AccountVerify) TableName() string {
	return "account_verify"
}

type Session struct {
	ID           uuid.UUID `json:"ID"`
	AccountID    string    `json:"account_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIP     string    `json:"client_ip"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (Session) TableName() string {
	return "session"
}
