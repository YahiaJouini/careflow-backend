package routes

import (
	"github.com/YahiaJouini/chat-app-backend/api/handlers/public"
	"github.com/gorilla/mux"
)

func InitPublicRoutes(router *mux.Router) {
	// Public Specialties (No Auth) - Used for registration forms/landing pages
	router.HandleFunc("/specialties", public.GetSpecialties).Methods("GET")
	router.HandleFunc("/doctors", public.GetDoctors).Methods("GET")
}
