package routes

import (
	"github.com/YahiaJouini/chat-app-backend/api/handlers/admin"
	"github.com/gorilla/mux"
)

func InitAdminRoutes(router *mux.Router) {
	// specialities
	router.HandleFunc("/specialties", admin.CreateSpecialty).Methods("POST")
	router.HandleFunc("/specialties", admin.GetAllSpecialties).Methods("GET")
	router.HandleFunc("/specialties/{id}", admin.UpdateSpecialty).Methods("PUT")
	router.HandleFunc("/specialties/{id}", admin.DeleteSpecialty).Methods("DELETE")

	// user management
	router.HandleFunc("/users", admin.CreateUser).Methods("POST")
	router.HandleFunc("/users", admin.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", admin.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}/role", admin.UpdateUserRole).Methods("PATCH")
	router.HandleFunc("/doctors/{id}/verify", admin.VerifyDoctor).Methods("PATCH")

	// statistics
	router.HandleFunc("/stats", admin.GetDashboardOverview).Methods("GET")
}
