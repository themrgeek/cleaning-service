package controllers

import (
	"net/http"

	"github.com/themrgeek/cleaning-service/pkg/model"
	"github.com/themrgeek/cleaning-service/pkg/utils"
)

// func CreateBooking(c *gin.Context) {
// 	var booking model.Booking

// 	if err := c.ShouldBindJSON(&booking); err != nil {
// 		utils.RespondWithError(c, nil, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	if err := model.CreateBooking(&booking); err != nil {
// 		utils.RespondWithError(c, nil, http.StatusInternalServerError, "Failed to save booking")
// 		return
// 	}

//		c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully"})
//	}
//

// CreateBooking handles the creation of a booking
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	// Extract and sanitize input parameters
	address := utils.SanitizeString(r.URL.Query().Get("address"))
	serviceType := utils.SanitizeString(r.URL.Query().Get("service_type"))
	// Validate required parameters
	if address == "" || serviceType == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Both 'address' and 'service_type' parameters are required")
		return
	}

	// Validate service type
	if !utils.IsValidServiceType(serviceType) {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid service type")
		return
	}

	// Create booking payload
	booking := model.BookingPayload{
		Address:     address,
		ServiceType: serviceType,
	}

	// Create database record
	model.CreateBooking(&booking)
	// Success response
	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Booking created successfully",
		"booking": booking,
	})
}
