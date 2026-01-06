package doctor

import (
	"net/http"

	"github.com/YahiaJouini/careflow/api/middleware"
	"github.com/YahiaJouini/careflow/internal/db/queries"
	"github.com/YahiaJouini/careflow/pkg/auth"
	"github.com/YahiaJouini/careflow/pkg/response"
)

func GetPatients(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	patients, err := queries.GetDoctorPatients(claims.UserID)
	if err != nil {
		response.ServerError(w, err.Error())
		return
	}

	response.Success(w, patients, "Patients retrieved successfully")
}
