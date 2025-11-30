package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YahiaJouini/careflow/internal/db"
	"github.com/YahiaJouini/careflow/internal/db/models"
	"github.com/YahiaJouini/careflow/pkg/response"
	"github.com/YahiaJouini/careflow/pkg/utils"
	"github.com/gorilla/mux"
)

type SpecialtyInput struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"omitempty,max=255"`
	Icon        string `json:"icon" validate:"required,url"`
}

func CreateSpecialty(w http.ResponseWriter, r *http.Request) {
	var body SpecialtyInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := utils.Validate.Struct(body); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	specialty := models.Specialty{
		Name:        body.Name,
		Description: body.Description,
		Icon:        body.Icon,
	}

	if result := db.Db.Create(&specialty); result.Error != nil {
		response.ServerError(w, "Database error: "+result.Error.Error())
		return
	}

	response.Success(w, specialty, "Specialty created successfully")
}

func GetAllSpecialties(w http.ResponseWriter, r *http.Request) {
	var specialties []models.Specialty

	if result := db.Db.Find(&specialties); result.Error != nil {
		response.ServerError(w, "Could not fetch specialties")
		return
	}

	response.Success(w, specialties, "Specialties retrieved")
}

func UpdateSpecialty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var body SpecialtyInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := utils.Validate.Struct(body); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var specialty models.Specialty
	if result := db.Db.First(&specialty, id); result.Error != nil {
		response.Error(w, http.StatusNotFound, "Specialty not found")
		return
	}

	specialty.Name = body.Name
	specialty.Description = body.Description

	if result := db.Db.Save(&specialty); result.Error != nil {
		response.ServerError(w, "Failed to update specialty")
		return
	}

	response.Success(w, specialty, "Specialty updated successfully")
}

func DeleteSpecialty(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	result := db.Db.Delete(&models.Specialty{}, id)
	if result.Error != nil {
		response.ServerError(w, "Failed to delete specialty")
		return
	}

	if result.RowsAffected == 0 {
		response.Error(w, http.StatusNotFound, "Specialty not found")
		return
	}

	response.Success(w, nil, "Specialty deleted successfully")
}
