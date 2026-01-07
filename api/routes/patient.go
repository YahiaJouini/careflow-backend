package routes

import (
	"github.com/YahiaJouini/careflow/api/handlers/patient"
	"github.com/gorilla/mux"
)

func InitPatientRoutes(router *mux.Router) {
	router.HandleFunc("/me", patient.GetPatient).Methods("GET")
	router.HandleFunc("/me", patient.UpdatePatient).Methods("PUT")
}
