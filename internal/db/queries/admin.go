package queries

import (
	"errors"

	"github.com/YahiaJouini/careflow/internal/db"
	"github.com/YahiaJouini/careflow/internal/db/models"
)

func GetAllUsers(roleFilter string) ([]models.User, error) {
	var users []models.User

	query := db.Db.Model(&models.User{}).Preload("Doctor").Preload("Doctor.Specialty")

	if roleFilter != "" {
		query = query.Where("role = ?", roleFilter)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func VerifyDoctor(doctorID uint) error {
	var doctor models.Doctor

	if err := db.Db.First(&doctor, doctorID).Error; err != nil {
		return errors.New("doctor not found with this ID")
	}

	doctor.IsVerified = true

	if err := db.Db.Save(&doctor).Error; err != nil {
		return err
	}

	return nil
}
