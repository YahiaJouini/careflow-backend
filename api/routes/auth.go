package routes

import (
	"net/http"

	"github.com/YahiaJouini/chat-app-backend/api/handlers/auth"
	"github.com/YahiaJouini/chat-app-backend/api/middleware"
	"github.com/gorilla/mux"
)

func InitAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.HandleFunc("/register", auth.Register).Methods("POST")
	router.HandleFunc("/verify-email", auth.ValidateCode).Methods("POST")
	router.HandleFunc("/resend-verification", auth.ResendCode).Methods("POST")
	router.HandleFunc("/google-login", auth.GoogleLogin).Methods("POST")
	// after login
	router.HandleFunc("/logout", auth.Logout).Methods("POST")
	router.HandleFunc("/refresh-token", auth.RefreshToken).Methods("POST")
	// check if user is authenticated
	router.Handle("/verify", middleware.AuthMiddleware(middleware.All)(http.HandlerFunc(auth.Authenticated))).Methods("GET")
	router.Handle("/verify-admin", middleware.AuthMiddleware(middleware.Admin)(http.HandlerFunc(auth.Authenticated))).Methods("GET")
	router.Handle("/verify-doctor", middleware.AuthMiddleware(middleware.Doctor)(http.HandlerFunc(auth.Authenticated))).Methods("GET")
}
