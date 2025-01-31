package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/themrgeek/cleaning-service/pkg/model"
	"github.com/themrgeek/cleaning-service/pkg/utils"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.RespondWithError(nil, w, http.StatusUnauthorized, "Authorization header is missing")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims := &utils.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if err != nil || !token.Valid {
		utils.RespondWithError(nil, w, http.StatusUnauthorized, "Invalid token")
		return
	}

	userEmail := claims.Email
	user, err := model.GetUserDetails(userEmail)
	if err != nil {
		utils.RespondWithError(nil, w, http.StatusInternalServerError, "Error fetching user details")
		return
	}

	if user == nil {
		utils.RespondWithError(nil, w, http.StatusNotFound, "User not found")
		return
	}

	bookings, err := model.GetUserBookings(userEmail)
	if err != nil {
		utils.RespondWithError(nil, w, http.StatusInternalServerError, "Error fetching bookings")
		return
	}

	// Simulate asynchronous processing
	go func() {
		// Simulate a long-running task
		// You can replace this with actual asynchronous processing logic
		// For example, fetching additional data, processing images, etc.
	}()

	w.WriteHeader(http.StatusAccepted)
	userProfile := struct {
		User     *model.User      `json:"user"`
		Bookings []*model.Booking `json:"bookings"`
	}{
		User:     user,
		Bookings: bookings,
	}

	utils.RespondWithJSON(w, http.StatusAccepted, userProfile)
}
