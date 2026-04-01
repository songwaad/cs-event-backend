package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB        DBConfig
	JWTSecret string
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func LoadConfig(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}
	return os.Getenv(key)
}

func Load() Config {
	dbPort, err := strconv.Atoi(LoadConfig("DB_PORT"))
	if err != nil {
		log.Fatal("Invalid DB_PORT value")
	}

	return Config{
		DB: DBConfig{
			Host:     LoadConfig("DB_HOST"),
			Port:     dbPort,
			User:     LoadConfig("DB_USER"),
			Password: LoadConfig("DB_PASSWORD"),
			Name:     LoadConfig("DB_NAME"),
		},
		JWTSecret: LoadConfig("JWT_SECRET_KEY"),
	}
}
