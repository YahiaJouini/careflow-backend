package queries

import (
	"errors"
	"github.com/YahiaJouini/chat-app-backend/pkg/email"
	"time"

	"github.com/YahiaJouini/chat-app-backend/internal/db"
	"github.com/YahiaJouini/chat-app-backend/internal/db/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := db.Db.Take(&user, "email = ?", email)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func MarkAsVerified(email string) error {
	result := db.Db.Model(models.User{}).Where("email = ?", email).Update("verified", true)
	return result.Error
}

func UpdateVerificationCode(userEmail string) (string, int, error) {
	user, err := GetUserByEmail(userEmail)
	if err != nil {
		return "", 404, err
	}
	newCode, err := email.GenerateVerificationCode()
	if err != nil {
		return "", 500, err
	}
	expiresAt := time.Now().Add(15 * time.Minute)
	user.VerificationCode = newCode
	user.CodeExpirationTime = expiresAt
	db.Db.Save(&user)
	return newCode, 200, nil
}
