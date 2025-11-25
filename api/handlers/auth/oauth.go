package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/YahiaJouini/chat-app-backend/internal/config"
	"github.com/YahiaJouini/chat-app-backend/internal/db/models"
	"github.com/YahiaJouini/chat-app-backend/internal/db/queries"
	"github.com/YahiaJouini/chat-app-backend/pkg/auth"
	"github.com/YahiaJouini/chat-app-backend/pkg/response"
	"google.golang.org/api/idtoken"
)

type GoogleLoginBody struct {
	IDToken string `json:"idToken" validate:"required"`
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	GoogleClientID, clientId_error := config.GetEnv("GOOGLE_CLIENT_ID")
	if clientId_error != nil {
		response.ServerError(w, "Google Client ID not configured")
		return
	}
	var body GoogleLoginBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.ServerError(w, "Invalid request body")
		return
	}

	payload, err := idtoken.Validate(context.Background(), body.IDToken, GoogleClientID)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "Invalid Google Token: "+err.Error())
		return
	}

	email := payload.Claims["email"].(string)
	emailVerified := payload.Claims["email_verified"].(bool)
	firstName := payload.Claims["given_name"].(string)
	lastName := payload.Claims["family_name"].(string)
	picture := payload.Claims["picture"].(string)

	if !emailVerified {
		response.Error(w, http.StatusForbidden, "Google email not verified")
		return
	}

	user, err := queries.GetUserByEmail(email)
	if err != nil {
		newUser := models.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Image:     picture,
			Verified:  true,
			Role:      "patient",
		}

		if err := queries.CreateUser(&newUser); err != nil {
			response.ServerError(w, "Could not create user: "+err.Error())
			return
		}
		user = &newUser
	} else {
		updateUserBody := queries.UpdateUserBody{
			Image: &picture,
		}
		queries.UpdateUser(user.ID, updateUserBody)
	}

	// same as login
	refreshToken := auth.GenerateToken(user, auth.RefreshToken)
	accessToken := auth.GenerateToken(user, auth.AccessToken)

	userAgent := r.Header.Get("User-Agent")
	if strings.Contains(userAgent, "Android") {
		mobileData := auth.MobileAuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			User:         user,
		}
		response.Success(w, mobileData, "Google Login successful")
		return
	}

	auth.SetAuthCookie(w, refreshToken, auth.Add)
	response.Success(w, accessToken, "Google Login successful")
}
