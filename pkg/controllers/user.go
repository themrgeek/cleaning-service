package controllers

import (
	"cleaning-service/pkg/model"
	"cleaning-service/pkg/services"
	"cleaning-service/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func BookService(w http.ResponseWriter, r *http.Request) {
	var booking model.BookingRequest
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	otp, err := services.GenerateOTP(booking.Email)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error generating OTP")
		return
	}

	// Send OTP via email
	go services.SendEmail(booking.Email, "OTP for booking", otp)

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "OTP sent to email"})
}

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var otpRequest model.OTPRequest
	err := json.NewDecoder(r.Body).Decode(&otpRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	isValid, err := services.VerifyOTP(otpRequest.Email, otpRequest.OTP)
	if err != nil || !isValid {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	// Allocate cleaner and send email
	cleaner, err := model.AllocateCleaner(otpRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error allocating cleaner")
		return
	}

	go services.SendEmail(otpRequest.Email, "Cleaner allocated", cleaner.Name)

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Cleaner allocated", "cleaner": cleaner.Name})
}

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	// Implementation for processing payment
}

func SubmitReview(w http.ResponseWriter, r *http.Request) {
	// Implementation for submitting review
}

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appointmentID := vars["id"]

	err := model.DeleteAppointment(appointmentID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Appointment deleted successfully"})
}
