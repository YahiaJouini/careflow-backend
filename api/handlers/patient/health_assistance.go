package patient

import (
	"encoding/json"
	"net/http"

	"github.com/YahiaJouini/careflow/internal/db/queries"
	"github.com/YahiaJouini/careflow/pkg/response"
)

func HealthAssistance(w http.ResponseWriter, r *http.Request) {
	var req queries.HealthAssistanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	data, err := queries.GetHealthAssistance(req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, data, "Health assistance retrieved successfully")
}
