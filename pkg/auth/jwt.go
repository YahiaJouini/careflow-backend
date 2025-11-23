package auth

import (
	"errors"
	"fmt"
	"github.com/YahiaJouini/chat-app-backend/internal/config"
	"github.com/YahiaJouini/chat-app-backend/internal/db/models"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
type Key string

const (
	RefreshToken Key = "REFRESH_TOKEN"
	AccessToken  Key = "ACCESS_TOKEN"
)

func GenerateToken(user *models.User, key Key) string {
	secret, _ := config.GetEnv(string(key))
	var tokenTime time.Duration
	if key == "REFRESH_TOKEN" {
		tokenTime = time.Hour * 24 * 30 // one month
	} else {
		tokenTime = time.Minute * 15 // 15 minutes
	}
	claims := Claims{
		UserID:   user.ID,
		Email:    user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal("An error generating token", err)
	}

	return tokenString
}

func VerifyToken(tokenString string, key Key) (*Claims, error) {
	secret, err := config.GetEnv(string(key))
	if err != nil {
		return nil, fmt.Errorf("error getting secret key: %w", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func GetRefreshToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return "", errors.New("auth_token cookie not fount")
	}
	token := cookie.Value
	return token, nil
}
