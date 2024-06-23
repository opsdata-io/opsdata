package utils

import (
	"fmt"
	"log"

	"github.com/opsdata-io/opsdata/pkg/config"
	"github.com/opsdata-io/opsdata/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.CFG.DBUser,
		config.CFG.DBPassword,
		config.CFG.DBHost,
		config.CFG.DBPort,
		config.CFG.DBName,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	// Migrate the schema
	DB.AutoMigrate(
		&models.User{},
		&models.UploadLink{},
		&models.FileMetadata{},
		&models.Customer{},
	)

	// Check if admin user exists
	var admin models.User
	result := DB.Where("email = ?", config.CFG.AdminEmail).First(&admin)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// Create the admin user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.CFG.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}
		admin = models.User{
			Email:    config.CFG.AdminEmail,
			Password: string(hashedPassword),
		}
		DB.Create(&admin)
		fmt.Printf("Admin user created with email: %s and password: %s\n", config.CFG.AdminEmail, config.CFG.AdminPassword)
	}
}

func VerifyCredentials(email, password string) bool {
	var user models.User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func SaveUploadLink(link *models.UploadLink) error {
	return DB.Create(link).Error
}

func GetUploadLink(linkID string) (*models.UploadLink, error) {
	var link models.UploadLink
	result := DB.Where("id = ?", linkID).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func SaveFileMetadata(file *models.FileMetadata) error {
	return DB.Create(file).Error
}

func GetFilesByUserID(userID uint) ([]models.FileMetadata, error) {
	var files []models.FileMetadata
	result := DB.Where("user_id = ?", userID).Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}
	return files, nil
}
