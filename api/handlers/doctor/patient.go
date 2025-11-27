package doctor

import (
	"net/http"

	"github.com/YahiaJouini/chat-app-backend/api/middleware"
	"github.com/YahiaJouini/chat-app-backend/internal/db/queries"
	"github.com/YahiaJouini/chat-app-backend/pkg/auth"
	"github.com/YahiaJouini/chat-app-backend/pkg/response"
)

func GetPatients(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	patients, err := queries.GetDoctorPatients(claims.UserID)
	if err != nil {
		response.ServerError(w, "Failed to fetch patients")
		return
	}

	response.Success(w, patients, "Patients retrieved successfully")
}
