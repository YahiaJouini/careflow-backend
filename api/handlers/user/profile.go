package user

import (
	"github.com/YahiaJouini/chat-app-backend/api/middleware"
	"github.com/YahiaJouini/chat-app-backend/internal/db/queries"
	"github.com/YahiaJouini/chat-app-backend/pkg/auth"
	"github.com/YahiaJouini/chat-app-backend/pkg/response"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value(middleware.UserClaimsKey).(*auth.Claims)

	data, err := queries.GetUserByEmail(claims.Email)
	if err != nil {
		auth.SetAuthCookie(w, "", auth.Remove)
		response.Unauthorized(w, "User not found. Logged out.")
		return
	}

	response.Success(w, data, "User retrieved successfully")
}
