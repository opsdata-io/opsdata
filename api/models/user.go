package models

import (
	"time"
)

type User struct {
	Id        uint      `json:"-"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  []byte    `json:"-"`
	Status    string    `gorm:"default:pending" json:"-"`
	CreatedAt time.Time `json:"created" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type Key struct {
	Id        uint      `json:"-"`
	Uuid      string    `json:"uuid"`
	Accesskey string    `json:"accesskey" gorm:"unique"`
	Secretkey string    `json:"secretkey"`
	Status    string    `gorm:"default:pending" json:"-"`
	CreatedAt time.Time `json:"created" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated" gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Expires   time.Time `json:"expires" gorm:"type:DATETIME DEFAULT NULL"`
}
