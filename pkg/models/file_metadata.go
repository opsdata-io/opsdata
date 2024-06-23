package models

import (
	"time"

	"gorm.io/gorm"
)

type FileMetadata struct {
	ID        string         `gorm:"primaryKey"`
	LinkID    string         `gorm:"not null"`
	FileName  string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
