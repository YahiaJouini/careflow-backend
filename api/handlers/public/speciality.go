package public

import (
	"net/http"

	"github.com/YahiaJouini/careflow/internal/db"
	"github.com/YahiaJouini/careflow/internal/db/models"
	"github.com/YahiaJouini/careflow/pkg/response"
)

type PublicSpecialty struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func GetSpecialties(w http.ResponseWriter, r *http.Request) {
	var specialties []models.Specialty

	if result := db.Db.Find(&specialties); result.Error != nil {
		response.ServerError(w, "Could not fetch specialties")
		return
	}

	publicResponse := make([]PublicSpecialty, len(specialties))
	for i, s := range specialties {
		publicResponse[i] = PublicSpecialty{
			ID:          s.ID,
			Name:        s.Name,
			Description: s.Description,
			Icon:        s.Icon,
		}
	}

	response.Success(w, publicResponse, "Specialties retrieved")
}
