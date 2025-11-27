package routes

import (
	"github.com/YahiaJouini/chat-app-backend/api/handlers/doctor"
	"github.com/gorilla/mux"
)

func InitDoctorRoutes(router *mux.Router) {
	router.HandleFunc("/stats", doctor.GetDashboardOverview).Methods("GET")

	// --- Appointments ---
	// Read
	router.HandleFunc("/appointments", doctor.GetAppointments).Methods("GET")

	// Validate (Confirm/Complete)
	router.HandleFunc("/appointments/{id}/validate", doctor.ValidateAppointment).Methods("PATCH")

	// Update (Edit Notes/Date)
	router.HandleFunc("/appointments/{id}", doctor.UpdateAppointment).Methods("PUT")

	// Delete (Cancel)
	router.HandleFunc("/appointments/{id}", doctor.CancelAppointment).Methods("DELETE")
}
