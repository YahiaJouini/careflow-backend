package routes

import (
	"github.com/YahiaJouini/careflow/api/handlers/public"
	"github.com/gorilla/mux"
)

func InitPublicRoutes(router *mux.Router) {
	router.HandleFunc("/specialties", public.GetSpecialties).Methods("GET")
	router.HandleFunc("/doctors", public.GetDoctors).Methods("GET")
}
