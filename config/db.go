package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
	// DB_PROD *gorm.DB
}

func InitialDB() Database {

	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DBHost, DBUser, DBPass, DBName, DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(40)

	return Database{
		DB: db,
	}
}
