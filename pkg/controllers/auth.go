package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/themrgeek/cleaning-service/pkg/model"
	"github.com/themrgeek/cleaning-service/pkg/services"
	"github.com/themrgeek/cleaning-service/pkg/utils"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	user := model.User{
		Name:     r.URL.Query().Get("name"),
		Email:    r.URL.Query().Get("email"),
		Password: r.URL.Query().Get("password"),
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing required query parameters")
		return
	}

	hashedPassword, err := services.HashPassword(user.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}
	user.Password = hashedPassword

	err = model.CreateUser(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}
func Login(w http.ResponseWriter, r *http.Request) {
	var creds model.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := model.AuthenticateUser(creds)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := services.GenerateJWT(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
