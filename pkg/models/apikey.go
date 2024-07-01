package models

import (
	"time"
)

// APIKey represents an API key in the database
type APIKey struct {
	ID            uint      `json:"id" gorm:"primaryKey; autoIncrement"`
	AccessKey     string    `json:"accessKey" gorm:"not null; unique"`
	SecretKeyHash string    `json:"-" gorm:"not null"`
	Description   string    `json:"description"`
	CustomerID    uint      `json:"customerId"` // Foreign key to Customer (optional)
	UserID        uint      `json:"userId"`     // Foreign key to User (optional)
	IsAdminKey    bool      `json:"isAdminKey"`
	ExpiredAt     time.Time `json:"expiredAt"`
}
