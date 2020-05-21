package schema

import (
	"time"
)

type TaxiProposeEventStatus string

const (
	TaxiProposeEventStatusUserCancel       = "CUSTOMER_CANCEL"
	TaxiProposeEventStatusTaxiCancel       = "TAXI_CANCEL"
	TaxiProposeEventStatusBookingConfirmed = "BOOKING_CONFIRMED"
)

type TaxiProposeEvent struct {
	ID                        string                 `json:"id"`
	IDUserRequestBookingEvent string                 `json:"id_user_request_booking_event"`
	IDTaxi                    string                 `json:"id_taxi"`
	Detail                    string                 `json:"detail"`
	Lat                       float64                `json:"lat"`
	Lon                       float64                `json:"lon"`
	Status                    TaxiProposeEventStatus `json:"status"`
	CreatedAt                 time.Time              `json:"created_at"`
	UpdatedAt                 time.Time              `json:"updated_at"`
}
