package schema

import (
	"time"

	"github.com/olivere/elastic"
)

type TaxiStatus string

const (
	TaxiStatusAvailable = "AVAILABLE"
	TaxiStatusPassenger = "PASSENGER"
	TaxiStatusDeny      = "DENY"
	TaxiStatusOffline   = "OFFLINE"
)

type Taxi struct {
	ID        string            `json:"id"`
	IDCar     string            `json:"id_car"`
	IDCenter  string            `json:"id_center"`
	IDDriver  string            `json:"id_driver"`
	Body      string            `json:"body"`
	Lat       float64           `json:"lat"`
	Lon       float64           `json:"lon"`
	Status    TaxiStatus        `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Data      interface{}       `json:"data" gorm:"-"`
	Location  *elastic.GeoPoint `json:"location" gorm:"-"`
}
