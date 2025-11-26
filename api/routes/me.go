package routes

import (
	"github.com/YahiaJouini/chat-app-backend/api/handlers/me"
	"github.com/gorilla/mux"
)

func InitUserRoutes(router *mux.Router) {
	router.HandleFunc("", me.GetUser).Methods("GET")
	router.HandleFunc("", me.UpdateUser).Methods("PUT")
	router.HandleFunc("", me.DeleteUser).Methods("DELETE")
}
