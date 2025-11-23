package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found", err)
	}
}

func GetEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if exists {
		return value, nil
	}
	return "", errors.New("key : " + key + "not found")
}
