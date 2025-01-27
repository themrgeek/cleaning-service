package controllers

import (
	"net/http"

	"github.com/themrgeek/cleaning-service/pkg/model"
	"github.com/themrgeek/cleaning-service/pkg/utils"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		utils.RespondWithError(nil, w, http.StatusBadRequest, "Email parameter is missing")
		return
	}

	user, err := model.GetUserDetails(userEmail)
	if err != nil {
		utils.RespondWithError(nil, w, http.StatusInternalServerError, "Error fetching user details")
		return
	}

	if user == nil {
		utils.RespondWithError(nil, w, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}
