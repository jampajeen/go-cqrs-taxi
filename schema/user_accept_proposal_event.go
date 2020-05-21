package schema

import (
	"time"
)

type UserAcceptProposalEvent struct {
	ID                        string    `json:"id"`
	IDTaxiProposeEvent        string    `json:"id_taxi_propose_event"`
	IDUserRequestBookingEvent string    `json:"id_user_request_booking_event"`
	Lat                       float64   `json:"lat"`
	Lon                       float64   `json:"lon"`
	CreatedAt                 time.Time `json:"created_at"`
}
