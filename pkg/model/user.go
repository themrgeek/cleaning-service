package model

import (
	"errors"

	"github.com/themrgeek/cleaning-service/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CheckPasswordHash compares a plain password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type BookingRequest struct {
	ServiceType   string `json:"service_type"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	DateOfBooking string `json:"date_of_booking"`
}

type OTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type Appointment struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type Cleaner struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
		return nil, err
	}
	if !CheckPasswordHash(creds.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}

func AllocateCleaner(request OTPRequest) (*Cleaner, error) {
	var cleaner Cleaner
	err := config.DB.QueryRow("SELECT id, name, email FROM cleaners WHERE available = 1 ORDER BY proximity LIMIT 1").Scan(&cleaner.ID, &cleaner.Name, &cleaner.Email)
	if err != nil {
		return nil, errors.New("no available cleaners found")
	}

	_, err = config.DB.Exec("UPDATE cleaners SET available = 0 WHERE id = ?", cleaner.ID)
	if err != nil {
		return nil, errors.New("error updating cleaner availability")
	}

	return &cleaner, nil
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
