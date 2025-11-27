package public

import (
	"net/http"
	"strconv"

	"github.com/YahiaJouini/chat-app-backend/internal/db/queries"
	"github.com/YahiaJouini/chat-app-backend/pkg/response"
)

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("specialityId")
	var specialtyID uint

	if queryParam != "" {
		id, err := strconv.Atoi(queryParam)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid specialityId")
			return
		}
		specialtyID = uint(id)
	}

	data, err := queries.GetPublicDoctors(specialtyID)
	if err != nil {
		response.ServerError(w, "Failed to fetch doctors")
		return
	}

	// 3. Return Response
	response.Success(w, data, "Doctors retrieved successfully")
}
