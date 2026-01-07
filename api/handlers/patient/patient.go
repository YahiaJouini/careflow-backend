package patient

import (
	"encoding/json"
	"net/http"

	"github.com/YahiaJouini/careflow/api/middleware"
	"github.com/YahiaJouini/careflow/internal/db/queries"
	"github.com/YahiaJouini/careflow/pkg/auth"
	"github.com/YahiaJouini/careflow/pkg/response"
	"github.com/YahiaJouini/careflow/pkg/utils"
)

func GetPatient(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)

	patient, err := queries.GetPatientByUserID(claims.UserID)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Patient profile not found")
		return
	}

	response.Success(w, patient, "Patient profile retrieved successfully")
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)

	var body queries.UpdatePatientBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.ServerError(w, err.Error())
		return
	}

	if err := utils.Validate.Struct(body); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedPatient, err := queries.UpdatePatient(claims.UserID, body)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, updatedPatient, "Patient profile updated successfully")
}
