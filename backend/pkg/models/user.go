package models

// User represents a user in the database
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     Role
}

// Role represents the role model
type Role string

// Standard list of Roles
const (
	RoleRoot             Role = "root"
	RoleTech             Role = "tech"
	RoleCustomerManager  Role = "customer_manager"
	RoleCustomerUser     Role = "customer_user"
	RoleCustomerReadonly Role = "customer_readonly"
)
