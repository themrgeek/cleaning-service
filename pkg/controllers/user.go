package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/themrgeek/cleaning-service/pkg/model"
	"github.com/themrgeek/cleaning-service/pkg/services"
	"github.com/themrgeek/cleaning-service/pkg/utils"

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

// UpdateStatus updates the status of a cleaner

func UpdateStatus(w http.ResponseWriter, r *http.Request) {

	// Implementation of the function

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Status updated"))

}

// ViewInquiries handles the GET /cleaner/inquiries request

func ViewInquiries(w http.ResponseWriter, r *http.Request) {

	// Implement the logic to view inquiries

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Inquiries viewed successfully"))

}

// CompleteInquiry handles the completion of an inquiry

func CompleteInquiry(w http.ResponseWriter, r *http.Request) {

	// Implement the logic to complete an inquiry

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Inquiry completed"))

}

// ViewPerformance handles the performance view request

func ViewPerformance(w http.ResponseWriter, r *http.Request) {

	// Implementation of the function

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Performance data"))

}

// ViewRevenue handles the request to view revenue

func ViewRevenue(w http.ResponseWriter, r *http.Request) {

	// Implementation for viewing revenue

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Revenue details"))

}
