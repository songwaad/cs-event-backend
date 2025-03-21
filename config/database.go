package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDatabase() *gorm.DB {
	if db == nil {
		return ConnectDatabase()
	}
	return db
}

func ConnectDatabase() *gorm.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		dbHost := os.Getenv("DB_HOST")
		dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			log.Fatal("Invalid DB_PORT value")
		}
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName)

		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level: Silent, Error, Warn, Info
				Colorful:      true,
			},
		)

		dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Fatalf("Failed to connect database: %v", err)
		}

		db = dbInstance
		fmt.Println("Database connected Successfully!")
	})

	return db
}
