package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type TaxiLocationUpdateMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.Taxi
}

func (m *TaxiLocationUpdateMessage) Key() string {
	return "taxi.location.update"
}
