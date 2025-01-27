package utils

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type contextKey string

const userIDKey contextKey = "userID"

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			RespondWithError(nil, w, http.StatusUnauthorized, "Missing Authorization Header")
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			RespondWithError(nil, w, http.StatusUnauthorized, "Invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
