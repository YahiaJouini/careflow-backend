package auth

import (
	"github.com/YahiaJouini/chat-app-backend/pkg/response"
	"net/http"
)

// this will run on every authenticated page load
func Authenticated(w http.ResponseWriter, r *http.Request) {
	response.Success(w, nil, "User authenticated")
}
