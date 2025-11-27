package routes

import (
	"github.com/YahiaJouini/chat-app-backend/api/handlers/patient"
	"github.com/gorilla/mux"
)

func InitPatientRoutes(router *mux.Router) {
	router.HandleFunc("/appointments", patient.CreateAppointment).Methods("POST")
	router.HandleFunc("/appointments", patient.GetAppointments).Methods("GET")
	router.HandleFunc("/appointments/{id}", patient.UpdateAppointment).Methods("PUT")
	router.HandleFunc("/appointments/{id}", patient.CancelAppointment).Methods("DELETE")
}
