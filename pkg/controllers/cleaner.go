package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/themrgeek/cleaning-service/pkg/model"
	"github.com/themrgeek/cleaning-service/pkg/utils"
)

type CleanerController struct {
	// Add any dependencies here, e.g., a service layer
}

func NewCleanerController() *CleanerController {
	return &CleanerController{}
}

func (c *CleanerController) GetCleaners(w http.ResponseWriter, r *http.Request) {
	cleaners, err := models.GetCleaners(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, cleaners)
}

func (c *CleanerController) GetCleaner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid cleaner ID")
		return
	}
	cleaner, err := models.GetCleanerByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, cleaner)
}

func (c *CleanerController) CreateCleaner(w http.ResponseWriter, r *http.Request) {
	var cleaner models.Cleaner
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cleaner); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	createdCleaner, err := models.CreateCleaner(r.Context(), cleaner)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, createdCleaner)
}

func (c *CleanerController) UpdateCleaner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid cleaner ID")
		return
	}
	var cleaner models.Cleaner
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cleaner); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	cleaner.ID = id
	updatedCleaner, err := models.UpdateCleaner(r.Context(), id, cleaner)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, updatedCleaner)
}

func (c *CleanerController) DeleteCleaner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid cleaner ID")
		return
	}
	if err := models.DeleteCleaner(r.Context(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
