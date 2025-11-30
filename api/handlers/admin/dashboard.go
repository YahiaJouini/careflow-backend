package admin

import (
	"net/http"

	"github.com/YahiaJouini/careflow/internal/db/queries"
	"github.com/YahiaJouini/careflow/pkg/response"
)

func GetDashboardOverview(w http.ResponseWriter, r *http.Request) {
	stats, err := queries.GetAdminStats()
	if err != nil {
		response.ServerError(w, "Failed to calculate admin stats")
		return
	}

	response.Success(w, stats, "Admin dashboard stats retrieved")
}
