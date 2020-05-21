package schema

import (
	"time"
)

type BookingStatus string

const (
	BookingStatusDone   BookingStatus = "DONE"
	BookingStatusCancel BookingStatus = "CANCEL"
)

type Booking struct {
	ID                        string        `json:"id"`
	IDTaxiProposeEvent        string        `json:"id_taxi_propose_event"`
	IDUserRequestBookingEvent string        `json:"id_user_request_booking_event"`
	Status                    BookingStatus `json:"status"`
	CreatedAt                 time.Time     `json:"created_at"`
	UpdatedAt                 time.Time     `json:"updated_at"`
	Data                      interface{}   `json:"data" gorm:"-"`
}
