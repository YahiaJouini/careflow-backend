package doctor

import (
	"net/http"

	"github.com/YahiaJouini/careflow/api/middleware"
	"github.com/YahiaJouini/careflow/internal/db/queries"
	"github.com/YahiaJouini/careflow/pkg/auth"
	"github.com/YahiaJouini/careflow/pkg/response"
)

func GetDashboardOverview(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)

	stats, err := queries.GetDoctorStats(claims.UserID)
	if err != nil {
		response.ServerError(w, err.Error())
		return
	}

	response.Success(w, stats, "Doctor dashboard stats retrieved")
}
