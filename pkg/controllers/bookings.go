package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/themrgeek/cleaning-service/pkg/model"
	"github.com/themrgeek/cleaning-service/pkg/utils"
)

func CreateBooking(c *gin.Context) {
	var booking model.Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		utils.RespondWithError(c, nil, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := model.CreateBooking(&booking); err != nil {
		utils.RespondWithError(c, nil, http.StatusInternalServerError, "Failed to save booking")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully"})
}
