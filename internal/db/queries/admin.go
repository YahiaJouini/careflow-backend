package queries

import (
	"errors"

	"github.com/YahiaJouini/chat-app-backend/internal/db"
	"github.com/YahiaJouini/chat-app-backend/internal/db/models"
)

// GetAllUsers retrieves users, optionally filtered by role
func GetAllUsers(roleFilter string) ([]models.User, error) {
	var users []models.User

	// Start query
	query := db.Db.Model(&models.User{}).Preload("Doctor").Preload("Doctor.Specialty")

	// Apply filter if provided
	if roleFilter != "" {
		query = query.Where("role = ?", roleFilter)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// VerifyDoctor sets IsVerified to true for a specific user's doctor profile
func VerifyDoctor(userID uint) error {
	var user models.User

	// Check if user exists and is a doctor
	if err := db.Db.Preload("Doctor").First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Role != "doctor" || user.Doctor == nil {
		return errors.New("user is not a doctor")
	}

	// Update the Doctor table specifically
	if err := db.Db.Model(&models.Doctor{}).Where("user_id = ?", userID).Update("is_verified", true).Error; err != nil {
		return err
	}

	return nil
}
