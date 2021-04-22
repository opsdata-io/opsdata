package models

import (
	"time"
)

type User struct {
	Id        uint      `json:"-"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"firstname"`
	Email     string    `json:"email" gorm:"unique"`
	Password  []byte    `json:"-"`
	Status    string    `gorm:"default:pending" json:"-"`
	CreatedAt time.Time `json:"created" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
