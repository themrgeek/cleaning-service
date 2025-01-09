package services

import (
	"errors"

	"github.com/themrgeek/cleaning-service/pkg/model"

	"sync"
)

var bookingStore = make(map[string]model.BookingRequest)

var mu sync.Mutex

func StoreBookingDetails(booking model.BookingRequest) error {

	mu.Lock()

	defer mu.Unlock()

	bookingStore[booking.Email] = booking

	return nil

}

func GetBookingDetails(email string) (model.BookingRequest, error) {

	mu.Lock()

	defer mu.Unlock()

	booking, exists := bookingStore[email]

	if !exists {

		return model.BookingRequest{}, errors.New("booking details not found")

	}

	return booking, nil

}
