package routes

import (
	"github.com/YahiaJouini/chat-app-backend/api/handlers/user"
	"github.com/gorilla/mux"
)

func InitUserRoutes(router *mux.Router) {
	router.HandleFunc("", user.GetUser).Methods("GET")
}
