package queries

import (
	"errors"
	"fmt"

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
func VerifyDoctor(doctorID uint) error {
	var doctor models.Doctor

	// DEBUG LOGS -------------------------
	fmt.Printf("--- DEBUG: VerifyDoctor Called ---\n")
	fmt.Printf("Received Doctor ID: %d\n", doctorID)
	// ------------------------------------

	// FIX: We search by Primary Key (ID), NOT user_id
	if err := db.Db.First(&doctor, doctorID).Error; err != nil {
		fmt.Printf("--- DEBUG: Doctor lookup failed: %v ---\n", err) // Log error
		return errors.New("doctor not found with this ID")
	}

	fmt.Printf("--- DEBUG: Found Doctor for UserID: %d. Verifying... ---\n", doctor.UserID)

	doctor.IsVerified = true

	if err := db.Db.Save(&doctor).Error; err != nil {
		return err
	}

	return nil
}
