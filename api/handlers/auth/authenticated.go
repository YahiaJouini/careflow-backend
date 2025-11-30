package auth

import (
	"net/http"

	"github.com/YahiaJouini/careflow/pkg/response"
)

// if this handler is reached, it means the user is authenticated with appropriate roles
func Authenticated(w http.ResponseWriter, r *http.Request) {
	response.Success(w, nil, "User authenticated")
}
