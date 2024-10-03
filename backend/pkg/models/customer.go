package models

// Customer represents a customer in the database
type Customer struct {
	ID                 uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	CompanyName        string `json:"companyName" gorm:"not null; unique"`
	Address            string `json:"address" gorm:"not null"`
	Notes              string `json:"notes"`
	SubscriptionStatus string `json:"subscriptionStatus" gorm:"not null"`
}
