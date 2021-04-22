package models

import (
	"time"
)

type User struct {
	Id        uint      `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"firstname"`
	Email     string    `json:"email" gorm:"unique"`
	Password  []byte    `json:"-"`
	Status    string    `gorm:"default:pending" json:"-"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"-"`
}
