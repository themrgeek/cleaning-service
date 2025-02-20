package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type BookingPayload struct {
	Address       string `json:"address"`
	ServiceType   string `json:"service_type"`
	CarType       string `json:"car_type"`
	DateOfBooking string `json:"date_of_booking"`
}
type BookingServiceStatus struct {
	StatusOfBooking string `json:"status"`
	IsCanceled      bool   `json:"is_canceled"`
}

func InitDB(dataSourceName string) error {
	var err error
	DB, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

// ALTER TABLE bookings
// ADD COLUMN status VARCHAR(255),
// ADD COLUMN is_canceled BOOLEAN DEFAULT false;
var DB *sqlx.DB

func Init() {
	if err := InitDB(os.Getenv("DB_DSN")); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
}

func CreateBooking(booking *BookingPayload) {
	Address := booking.Address
	ServiceType := booking.ServiceType
	CarType := booking.CarType
	BookingDate := booking.DateOfBooking
	fmt.Println(Address, ServiceType)
	query := "INSERT INTO bookings (address, service_type,car_type,date_of_booking,status,is_canceled) VALUES (?, ?, ?, ?, ?,?)"
	Init()
	fmt.Println("Database initialized...")
	fmt.Println(query)
	_, err := DB.Exec(query, Address, ServiceType, CarType, BookingDate, "pending", true)
	if err != nil {
		log.Println("Error creating booking:", err)
	} else {
		log.Println("Booking created successfully")
	}
}

func DeleteBooking(id uint) error {
	query := "DELETE FROM bookings WHERE booking_id = ?"
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting booking:", err)
		return err
	}
	return nil
}
