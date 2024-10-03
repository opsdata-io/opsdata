package models

// FileMetadata represents metadata for a file uploaded to the server and stored in the database
type FileMetadata struct {
	ID       string `gorm:"primaryKey"`
	LinkID   string `gorm:"not null"`
	FileName string `gorm:"not null"`
}
