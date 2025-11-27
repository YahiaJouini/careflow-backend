package doctor

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/YahiaJouini/chat-app-backend/api/middleware"
	"github.com/YahiaJouini/chat-app-backend/internal/db/queries"
	"github.com/YahiaJouini/chat-app-backend/pkg/auth"
	"github.com/YahiaJouini/chat-app-backend/pkg/response"
	"github.com/gorilla/mux"
)

// GET /doctor/appointments
func GetAppointments(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)

	data, err := queries.GetDoctorAppointments(claims.UserID)
	if err != nil {
		response.ServerError(w, "Failed to fetch appointments")
		return
	}
	response.Success(w, data, "Doctor appointments retrieved")
}

// PATCH /doctor/appointments/{id}/validate
func ValidateAppointment(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var req queries.ValidateAppointmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	data, err := queries.ValidateAppointment(claims.UserID, uint(id), req.Status)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, data, "Appointment status updated")
}

// PUT /doctor/appointments/{id}
func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var req queries.DoctorUpdateAppointmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	data, err := queries.UpdateAppointmentDoctor(claims.UserID, uint(id), req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, data, "Appointment updated successfully")
}

// DELETE /doctor/appointments/{id}
func CancelAppointment(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := queries.CancelAppointmentDoctor(claims.UserID, uint(id))
	if err != nil {
		response.Error(w, http.StatusNotFound, "Appointment not found")
		return
	}
	response.Success(w, nil, "Appointment cancelled")
}
