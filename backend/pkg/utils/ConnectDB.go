package utils

import (
	"fmt"

	"github.com/opsdata-io/opsdata/backend/pkg/config"
	"github.com/opsdata-io/opsdata/backend/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB

// ConnectDB connects to the database and performs migrations
func ConnectDB() {
	var err error

	// Determine the server DSN
	serverDSN := config.CFG.DSN
	if serverDSN == "" {
		serverDSN = fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
			config.CFG.DBUser,
			config.CFG.DBPassword,
			config.CFG.DBHost,
			config.CFG.DBPort,
		)
	}

	// Connect to MySQL server
	serverDB, err := gorm.Open(mysql.Open(serverDSN), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to MySQL server: %v", err)
	}

	// Check if the database exists and create it if it doesn't
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.CFG.DBName)
	if err := serverDB.Exec(createDBSQL).Error; err != nil {
		logger.Fatalf("Failed to create database: %v", err)
	}

	// Determine the database DSN
	dbDSN := config.CFG.DSN
	if dbDSN == "" {
		dbDSN = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.CFG.DBUser,
			config.CFG.DBPassword,
			config.CFG.DBHost,
			config.CFG.DBPort,
			config.CFG.DBName,
		)
	}

	// Connect to the specific database
	DB, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(
		&models.User{},
		&models.UploadLink{},
		&models.FileMetadata{},
		&models.Customer{},
	)
	if err != nil {
		logger.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Check if admin user exists
	var admin models.User
	result := DB.Where("email = ?", config.CFG.AdminEmail).First(&admin)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// Create the admin user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.CFG.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			logger.Fatalf("Failed to hash password: %v", err)
		}
		admin = models.User{
			Email:    config.CFG.AdminEmail,
			Password: string(hashedPassword),
		}
		DB.Create(&admin)
		logger.Printf("Admin user created with email: %s and password: %s\n", config.CFG.AdminEmail, config.CFG.AdminPassword)
	}
}
