package queries

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/YahiaJouini/careflow/internal/config"
)

type HealthAssistanceRequest struct {
	Symptom string `json:"symptom"`
}

type HealthAssistanceResponse struct {
	Status     string   `json:"status"`
	Answer     string   `json:"answer"`
	Tags       []string `json:"tags"`
	Confidence float64  `json:"confidence"`
}

func GetHealthAssistance(req HealthAssistanceRequest) (*HealthAssistanceResponse, error) {
	url, err := config.GetEnv("ASSISTANCE_MODEL_API")
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(req)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var assistanceResp HealthAssistanceResponse

	if err := json.NewDecoder(resp.Body).Decode(&assistanceResp); err != nil {
		return nil, err
	}

	return &assistanceResp, nil
}
