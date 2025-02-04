package model

import (
	"errors"
	"log"

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

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	log.Println("Error", err)
	return err
}
func GetUserDetails(userEmail string) (*User, error) {
	var user User
	err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE email = ?", userEmail).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
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
