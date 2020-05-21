package schema

import (
	"time"

	"github.com/olivere/elastic"
)

type (
	UserRequestBookingEventAppKind       string
	UserRequestBookingEventStatus        string
	UserRequestBookingEventPaymentMethod string
)

const (
	UserRequestBookingEventAppKindWeb                = "WEB"
	UserRequestBookingEventAppKindAndroid            = "ANDROID"
	UserRequestBookingEventAppKindIos                = "IOS"
	UserRequestBookingEventStatusUserSearching       = "CUSTOMER_SEARCHING"
	UserRequestBookingEventStatusTaxiProposed        = "TAXI_PROPOSED"
	UserRequestBookingEventStatusBookingConfirmed    = "BOOKING_CONFIRMED"
	UserRequestBookingEventStatusUserCancel          = "CUSTOMER_CANCEL"
	UserRequestBookingEventStatusTaxiCancel          = "TAXI_CANCEL"
	UserRequestBookingEventPaymentMethodAppCredit    = "APP_CREDIT"
	UserRequestBookingEventPaymentMethodCreditCard   = "CREDIT_CARD"
	UserRequestBookingEventPaymentMethodCash         = "CASH"
	UserRequestBookingEventPaymentMethodBankTransfer = "BANK_TRANSFER"
)

type UserRequestBookingEvent struct {
	ID              string                               `json:"id"`
	IDUser          string                               `json:"id_user"`
	PickUpTime      time.Time                            `json:"pick_up_time"`
	PickUpPlace     string                               `json:"pick_up_place"`
	PickUpPlaceLat  float64                              `json:"pick_up_place_lat"`
	PickUpPlaceLon  float64                              `json:"pick_up_place_lon"`
	DropOffPlace    string                               `json:"drop_off_place"`
	DropOffPlaceLat float64                              `json:"drop_off_place_lat"`
	DropOffPlaceLon float64                              `json:"drop_off_place_lon"`
	Conditions      string                               `json:"conditions"`
	Detail          string                               `json:"detail"`
	Lat             float64                              `json:"lat"`
	Lon             float64                              `json:"lon"`
	EstimatedKM     float32                              `json:"estimated_km"`
	EstimatedCharge float32                              `json:"estimated_charge"`
	Fee             float32                              `json:"fee"`
	AppKind         UserRequestBookingEventAppKind       `json:"app_kind"`
	PaymentMethod   UserRequestBookingEventPaymentMethod `json:"payment_method"`
	Status          UserRequestBookingEventStatus        `json:"status"`
	CreatedAt       time.Time                            `json:"created_at"`
	UpdatedAt       time.Time                            `json:"updated_at"`

	LocationPickUp  *elastic.GeoPoint `json:"location_pick_up" gorm:"-"`
	LocationDropOff *elastic.GeoPoint `json:"location_drop_off" gorm:"-"`
}
