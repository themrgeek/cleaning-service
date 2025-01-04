package model

import (
	"cleaning-service/pkg/config"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BookingRequest struct {
	ServiceType string `json:"service_type"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
}

type OTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type Appointment struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	return err
}

func AuthenticateUser(creds Credentials) (*User, error) {
	var user User
	err := config.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", creds.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare hashed passwords
	if !services.CheckPasswordHash(creds.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func AllocateCleaner(request OTPRequest) (*Cleaner, error) {
	// Implementation for allocating cleaner based on proximity and availability
}

func DeleteAppointment(appointmentID string) error {
	var status string
	err := config.DB.QueryRow("SELECT status FROM appointments WHERE id = ?", appointmentID).Scan(&status)
	if err != nil {
		return errors.New("appointment not found")
	}

	if status != "pending" {
		return errors.New("appointment cannot be deleted as it is already in processing")
	}

	_, err = config.DB.Exec("DELETE FROM appointments WHERE id = ?", appointmentID)
	if err != nil {
		return errors.New("error deleting appointment")
	}

	return nil
}
