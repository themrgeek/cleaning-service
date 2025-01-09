package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/themrgeek/cleaning-service/pkg/model"

	_ "github.com/lib/pq"
)

func StoreBookingToTable(booking model.BookingRequest) error {
	connStr := "user=username dbname=cleaning_service sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	query := `
		INSERT INTO Booking (customer_name, service_type, booking_date, address)
		VALUES ($1, $2, $3, $4)
	`
	_, err = db.Exec(query, booking.Name, booking.ServiceType, booking.DateOfBooking, booking.Address)
	if err != nil {
		return fmt.Errorf("failed to store booking: %v", err)
	}

	return nil
}
