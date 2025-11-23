package db

import (
	"database/sql"
	"github.com/YahiaJouini/chat-app-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitializeDB() *sql.DB {
	dsn, err := config.GetEnv("DATABASE_URL")
	if err != nil {
		log.Fatal("Failed to get environment variable DATABASE_URL: ", err)
	}
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the db: ", err)
	}
	log.Println("Database connected successfully")

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatal("Failed to get *sql.DB instance: ", err)
	}
	return sqlDB
}
