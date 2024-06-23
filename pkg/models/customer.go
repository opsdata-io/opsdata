package models

type Customer struct {
	ID                 uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	CompanyName        string `json:"companyName" gorm:"not null; unique"`
	Address            string `json:"address" gorm:"not null"`
	ContactName        string `json:"contactName" gorm:"not null"`
	ContactTitle       string `json:"contactTitle" gorm:"not null"`
	ContactEmail       string `json:"contactEmail" gorm:"not null; unique"`
	ContactPhone       string `json:"contactPhone" gorm:"not null"`
	Notes              string `json:"notes"`
	SubscriptionStatus string `json:"subscriptionStatus" gorm:"not null"`
}
