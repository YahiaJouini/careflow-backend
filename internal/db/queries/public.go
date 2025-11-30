package queries

import (
	"github.com/YahiaJouini/careflow/internal/db"
	"github.com/YahiaJouini/careflow/internal/db/models"
)

func GetPublicDoctors(specialtyID uint) ([]models.Doctor, error) {
	var doctors []models.Doctor

	query := db.Db.
		Preload("User").
		Preload("Specialty").
		Where("is_verified = ?", true)

	if specialtyID > 0 {
		query = query.Where("specialty_id = ?", specialtyID)
	}

	err := query.Find(&doctors).Error
	return doctors, err
}
