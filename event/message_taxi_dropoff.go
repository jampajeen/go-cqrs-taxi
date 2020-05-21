package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type TaxiDropOffMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.TaxiDropOffEvent
}

func (m *TaxiDropOffMessage) Key() string {
	return "taxi.dropoff"
}
