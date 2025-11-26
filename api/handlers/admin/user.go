package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YahiaJouini/chat-app-backend/internal/db"
	"github.com/YahiaJouini/chat-app-backend/internal/db/models"
	"github.com/YahiaJouini/chat-app-backend/internal/db/queries"
	"github.com/YahiaJouini/chat-app-backend/pkg/auth"
	"github.com/YahiaJouini/chat-app-backend/pkg/response"
	"github.com/YahiaJouini/chat-app-backend/pkg/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ChangeRoleInput struct {
	Role          string  `json:"role" validate:"required,oneof=admin doctor patient"`
	SpecialtyID   *uint   `json:"specialtyId" validate:"required_if=Role doctor"`
	LicenseNumber *string `json:"licenseNumber" validate:"required_if=Role doctor"`
}

type CreateUserInput struct {
	FirstName string `json:"firstName" validate:"required,min=3"`
	LastName  string `json:"lastName" validate:"required,min=3"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Role      string `json:"role" validate:"required,oneof=admin patient"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	roleFilter := r.URL.Query().Get("role")

	users, err := queries.GetAllUsers(roleFilter)
	if err != nil {
		response.ServerError(w, "Could not fetch users")
		return
	}

	response.Success(w, users, "Users retrieved successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := queries.DeleteUser(uint(id)); err != nil {
		response.ServerError(w, "Failed to delete user: "+err.Error())
		return
	}

	response.Success(w, nil, "User deleted successfully")
}

func VerifyDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := queries.VerifyDoctor(uint(id)); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, nil, "Doctor verified successfully")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := utils.Validate.Struct(body); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(body.Password)
	if err != nil {
		response.ServerError(w, "Error hashing password")
		return
	}

	newUser := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  hashedPassword,
		Role:      body.Role,
		Verified:  true,
	}

	if err := queries.CreateUser(db.Db, &newUser); err != nil {
		response.ServerError(w, "Database error: "+err.Error())
		return
	}

	response.Success(w, newUser, "User created successfully")
}

func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var body ChangeRoleInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := utils.Validate.Struct(body); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := db.Db.First(&user, id).Error; err != nil {
		response.Error(w, http.StatusNotFound, "User not found")
		return
	}

	// transaction needed because we might touch the doctors table
	err = db.Db.Transaction(func(tx *gorm.DB) error {
		user.Role = body.Role
		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		if body.Role == "doctor" {
			var count int64
			tx.Model(&models.Doctor{}).Where("user_id = ?", user.ID).Count(&count)

			if count == 0 {
				doctor := models.Doctor{
					UserID:        user.ID,
					SpecialtyID:   *body.SpecialtyID,
					LicenseNumber: *body.LicenseNumber,
				}
				if err := queries.CreateDoctor(tx, &doctor); err != nil {
					return err
				}
			}
		}

		if body.Role != "doctor" {
			tx.Where("user_id = ?", user.ID).Delete(&models.Doctor{})
		}

		return nil
	})
	if err != nil {
		response.ServerError(w, "Failed to update role: "+err.Error())
		return
	}

	response.Success(w, user, "User role updated successfully")
}
