package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	User          string
	Host          string
	Port          int
	Password      string
	Name          string
	EnableSSLMode bool
}

var dbConfigurations *DbConfig

var configurations *Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DbConfig
}

func LoadConfig() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to the env variables:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)

	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	if jwtSecretKey == "" {
		fmt.Println("Secret Key is Required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")

	if dbUser == "" {
		fmt.Println("DB User  is Required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		fmt.Println("DB Host  is Required")
		os.Exit(1)
	}
	dbPass := os.Getenv("DB_PASSWORD")

	if dbPass == "" {
		fmt.Println("DB Password  is Required")
		os.Exit(1)
	}

	db_port := os.Getenv("DB_PORT")

	if db_port == "" {
		fmt.Println("DB Port  is Required")
		os.Exit(1)
	}
	dbPort, err := strconv.ParseInt(db_port, 10, 64)

	if err != nil {
		fmt.Println("DB Port must be number")
		os.Exit(1)

	}

	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		fmt.Println("DB Name  is Required")
		os.Exit(1)
	}

	SSLMode := os.Getenv("DB_ENABLE_SSL_MODE")

	dbSSLMode, err := strconv.ParseBool(SSLMode)
	if err != nil {
		fmt.Println("Invalid enable ssl mode value")
		os.Exit(1)
	}
	dbConfigurations = &DbConfig{
		User:          dbUser,
		Host:          dbHost,
		Port:          int(dbPort),
		Password:      dbPass,
		Name:          dbName,
		EnableSSLMode: dbSSLMode,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfigurations,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		//first time
		LoadConfig()
	}
	return configurations
}
