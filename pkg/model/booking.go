package model

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type Booking struct {
	ID          uint      `json:"id" db:"id" gorm:"primaryKey"`
	Address     string    `json:"address" db:"address" binding:"required"`
	Date        string    `json:"date" db:"date" binding:"required"`
	ServiceType string    `json:"service_type" db:"service_type" binding:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Details     string    `json:"details"`
	Review      string    `json:"review"`
}

var DB *sqlx.DB

func CreateBooking(booking *Booking) error {
	query := "INSERT INTO Booking (Address, Date, ServiceType, CreatedAt) VALUES (?, DATE('now'), ?,NOW())"
	_, err := DB.Exec(query, booking.Address, booking.Date, time.Now())
	if err != nil {
		log.Println("Error creating booking:", err)
		return err
	}
	return nil
}
func DeleteBooking(id uint) error {
	query := "DELETE FROM bookings WHERE id = ?"
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting booking:", err)
		return err
	}
	return nil
}
