package services

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/themrgeek/cleaning-service/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))
var otpStore = make(map[string]string)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GenerateOTP(email string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	otpStore[email] = otp
	return otp, nil
}

func VerifyOTP(email, otp string) (bool, error) {
	storedOTP, exists := otpStore[email]
	if !exists || storedOTP != otp {
		return false, nil
	}
	delete(otpStore, email)
	return true, nil
}

func SendEmail(to string, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "your_email@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", "Your OTP code is: "+otp)

	d := gomail.NewDialer("smtp.example.com", 587, "your_email@example.com", "your_email_password")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
