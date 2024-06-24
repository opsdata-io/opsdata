package models

import (
	"time"

	"gorm.io/gorm"
)

// UploadLink represents a link to upload files to the server and store in the database
type UploadLink struct {
	ID         string         `gorm:"primaryKey"`
	UserID     string         `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Customer   string         `gorm:"not null"`
	CaseNumber string         `gorm:"not null"`
	Subject    string         `gorm:"not null"`
	Notes      string         `gorm:"not null"`
}
