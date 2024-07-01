package models

import "time"

// Server represents a server in the database
type Server struct {
	ID          uint      `json:"id" gorm:"primaryKey; autoIncrement"`
	CustomerID  uint      `json:"customerId" gorm:"index"` // Foreign key to Customer
	Name        string    `json:"name" gorm:"not null"`
	DeviceType  string    `json:"deviceType"`
	IPAddress   string    `json:"ipAddress"`
	Description string    `json:"description"`
	LastPing    time.Time `json:"lastPing" gorm: "current_timestamp"`
}
