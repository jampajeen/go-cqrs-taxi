package schema

import (
	"time"
)

type PaymentStatus string

const (
	PaymentStatusPaid     = "PAID"
	PaymentStatusRefunded = "REFUNDED"
)

type Payment struct {
	ID        string        `json:"id"`
	IDBooking string        `json:"id_booking"`
	Status    PaymentStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Data      interface{}   `json:"data" gorm:"-"`
}
