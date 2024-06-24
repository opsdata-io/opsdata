package models

// User represents a user in the database
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
