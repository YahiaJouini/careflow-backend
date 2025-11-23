package db

import (
	"fmt"
	"github.com/YahiaJouini/chat-app-backend/internal/db/models"
	"log"
)

func Migrate() {
	err := Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("migrations applied successfully")
}
