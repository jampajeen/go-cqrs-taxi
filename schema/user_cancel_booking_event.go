package schema

import (
	"time"
)

type UserCancelBookingEvent struct {
	ID                        string    `json:"id"`
	IDBooking                 string    `json:"id_booking"`
	IDUserRequestBookingEvent string    `json:"id_user_request_booking_event"` // required if no booking ID
	Lat                       float64   `json:"lat"`
	Lon                       float64   `json:"lon"`
	CreatedAt                 time.Time `json:"created_at"`
}
