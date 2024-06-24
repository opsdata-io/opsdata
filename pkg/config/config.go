package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/opsdata-io/opsdata/pkg/version"
)

// AppConfig structure for environment-based configurations.
type AppConfig struct {
	Debug               bool   `json:"debug"`
	MetricsPort         int    `json:"metricsPort"`
	ServerPort          int    `json:"serverPort"`
	Version             bool   `json:"version"`
	JWTSecret           string `json:"jwtSecret"`
	EncryptionKey       string `json:"encryptionKey"`
	S3Bucket            string `json:"s3Bucket"`
	S3Region            string `json:"s3Region"`
	S3Endpoint          string `json:"s3Endpoint"`
	S3AccessKey         string `json:"s3AccessKey"`
	S3SecretKey         string `json:"s3SecretKey"`
	DSN                 string `json:"dsn"`
	DBHost              string `json:"dbHost"`
	DBPort              int    `json:"dbPort"`
	DBUser              string `json:"dbUser"`
	DBPassword          string `json:"dbPassword"`
	DBName              string `json:"dbName"`
	AdminEmail          string `json:"adminEmail"`
	AdminPassword       string `json:"adminPassword"`
	SendGridAPIKey      string `json:"sendGridAPIKey"`
	SendGridSenderEmail string `json:"sendGridSenderEmail"`
	SendGridSenderName  string `json:"sendGridSenderName"`
}

// CFG is the global configuration instance populated by LoadConfiguration.
var CFG AppConfig

// LoadConfiguration loads the configuration from the environment variables and command line flags.
func LoadConfiguration() {
	debug := flag.Bool("debug", parseEnvBool("DEBUG", false), "Enable debug mode")
	metricsPort := flag.Int("metricsPort", parseEnvInt("METRICS_PORT", 9000), "Port for metrics server")
	serverPort := flag.Int("serverPort", parseEnvInt("SERVER_PORT", 8080), "Port for the server")
	showVersion := flag.Bool("version", false, "Show version and exit")
	jwtSecret := getEnvOrDefault("JWT_SECRET", "secret")
	encryptionKey := getEnvOrDefault("ENCRYPTION_KEY", "secret")
	s3Bucket := getEnvOrDefault("S3_BUCKET", "opsdata")
	s3Region := getEnvOrDefault("S3_REGION", "us-east-1")
	s3Endpoint := getEnvOrDefault("S3_ENDPOINT", "https://s3.wasabisys.com")
	s3AccessKey := getEnvOrDefault("S3_ACCESS_KEY", "your-access-key")
	s3SecretKey := getEnvOrDefault("S3_SECRET_KEY", "your-secret-key")
	dsn := getEnvOrDefault("DSN", "")
	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbPort := parseEnvInt("DB_PORT", 3306)
	dbUser := getEnvOrDefault("DB_USER", "root")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "password")
	dbName := getEnvOrDefault("DB_NAME", "opsdata")
	adminEmail := getEnvOrDefault("ADMIN_EMAIL", "admin@support.tools")
	adminPassword := getEnvOrDefault("ADMIN_PASSWORD", "password")
	sendGridAPIKey := getEnvOrDefault("SENDGRID_API_KEY", "your-sendgrid-api-key")
	sendGridSenderEmail := getEnvOrDefault("SENDGRID_SENDER_EMAIL", "<EMAIL>")
	sendGridSenderName := getEnvOrDefault("SENDGRID_SENDER_NAME", "<NAME>")

	flag.Parse()

	CFG.Debug = *debug
	CFG.MetricsPort = *metricsPort
	CFG.ServerPort = *serverPort
	CFG.Version = *showVersion
	CFG.JWTSecret = jwtSecret
	CFG.EncryptionKey = encryptionKey
	CFG.S3Bucket = s3Bucket
	CFG.S3Region = s3Region
	CFG.S3Endpoint = s3Endpoint
	CFG.S3AccessKey = s3AccessKey
	CFG.S3SecretKey = s3SecretKey
	CFG.DSN = dsn
	CFG.DBHost = dbHost
	CFG.DBPort = dbPort
	CFG.DBUser = dbUser
	CFG.DBPassword = dbPassword
	CFG.DBName = dbName
	CFG.AdminEmail = adminEmail
	CFG.AdminPassword = adminPassword
	CFG.SendGridAPIKey = sendGridAPIKey
	CFG.SendGridSenderEmail = sendGridSenderEmail
	CFG.SendGridSenderName = sendGridSenderName

	if CFG.Version {
		fmt.Printf("Version: %s\nGit Commit: %s\nBuild Time: %s\n", version.Version, version.GitCommit, version.BuildTime)
		os.Exit(0)
	}
}

// getEnvOrDefault returns the value of the environment variable with the given key or the default value if the key is not set.
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// parseEnvInt parses the environment variable with the given key and returns its integer representation or the default value if the key is not set.
func parseEnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error parsing %s as int: %v. Using default value: %d", key, err, defaultValue)
		return defaultValue
	}
	return intValue
}

// parseEnvBool parses the environment variable with the given key and returns its boolean representation or the default value if the key is not set.
func parseEnvBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Printf("Error parsing %s as bool: %v. Using default value: %t", key, err, defaultValue)
		return defaultValue
	}
	return boolValue
}
