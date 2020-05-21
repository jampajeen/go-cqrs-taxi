package schema

import (
	"time"
)

type (
	DriverStatus string
)

const (
	DriverStatusActive   DriverStatus = "ACTIVE"
	DriverStatusInactive DriverStatus = "INACTIVE"
)

type Driver struct {
	ID        string       `json:"id"`
	IDUser    string       `json:"id_user"`
	Status    DriverStatus `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Data      interface{}  `json:"data" gorm:"-"`
}
