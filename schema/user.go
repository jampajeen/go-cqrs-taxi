package schema

import (
	"time"
)

type (
	UserKind   string
	UserStatus string
)

const (
	UserKindDriver     = "DRIVER"
	UserKindPassenger  = "PASSENGER"
	UserStatusActive   = "ACTIVE"
	UserStatusInactive = "INACTIVE"
)

type User struct {
	ID             string      `json:"id"`
	NationalIDCard string      `json:"national_id_card"`
	FirstName      string      `json:"first_name"`
	MiddleName     string      `json:"middle_name"`
	LastName       string      `json:"last_name"`
	Email          string      `json:"email"`
	Password       string      `json:"-"`
	Gender         string      `json:"gender"`
	Age            string      `json:"age"`
	Address        string      `json:"address"`
	Phone          string      `json:"phone"`
	Kind           UserKind    `json:"kind"`
	Status         UserStatus  `json:"status"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Data           interface{} `json:"data" gorm:"-"`
}
