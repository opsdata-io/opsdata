package database

import (
	"fmt"
	"os"

	"github.com/mattmattox/opsdata/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DSN has exported fields representing MySQL connection parameters
type DSN struct {
	Username string
	Password string
	Hostname string
	Port     string
	Socket   string
	//
	DefaultsFile string
	Protocol     string
	//
	DefaultDb string
	Params    []string
}

func Connect() {

	databaseHostname := os.Getenv("DB_HOST")
	if databaseHostname == "" {
		databaseHostname = "localhost"
	}

	databasePort := os.Getenv("DB_PORT")
	if databasePort == "" {
		databasePort = "3306"
	}

	databaseUsername := os.Getenv("DB_USER")
	if databaseUsername == "" {
		databaseUsername = "opsdata"
	}

	databasePassword := os.Getenv("DB_PASS")
	if databasePassword == "" {
		databasePassword = "opsdata"
	}

	databaseName := os.Getenv("DB_NAME")
	if databaseName == "" {
		databaseName = "opsdata"
	}

	//dsn := "opsdata:opsdata@tcp(localhost:3306)/opsdata"
	dsn := fmt.Sprint(databaseUsername, ":", databasePassword, "@tcp(", databaseHostname, ":", databasePort, ")/", databaseName, "?parseTime=true")

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Access{})
}
