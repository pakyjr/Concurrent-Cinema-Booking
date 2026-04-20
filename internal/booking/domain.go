package booking

import (
	"errors"
	"time"
)

var ( //TODO syntax
	ErrSeatAlreadyBooked = errors.New("seat is already taken")
)

type Booking struct {
	ID string
	MovieID string
	SeatID string
	UserID string
	Status string
	ExpiredAt time.Time
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}
