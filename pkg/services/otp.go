package services

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOTP(email string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))

	// Store OTP in database with expiration time

	return otp, nil
}

func VerifyOTP(email, otp string) (bool, error) {
	// Check OTP from database

	return true, nil
}
