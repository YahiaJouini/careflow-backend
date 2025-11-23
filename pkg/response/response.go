package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(w http.ResponseWriter, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Response{
		Success: true,
		Message: message,
	}
	if data != nil {
		response.Data = data
	}

	json.NewEncoder(w).Encode(response)
}

func Error(w http.ResponseWriter, statusCode int, errorMessage string) {
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   errorMessage,
	})
}

func Unauthorized(w http.ResponseWriter, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   errorMessage,
	})
}

func ServerError(w http.ResponseWriter, errors ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	errorMsg := "An error occurred"
	if len(errors) > 0 && errors[0] != "" {
		errorMsg = errors[0]
	}
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   errorMsg,
	})
}
