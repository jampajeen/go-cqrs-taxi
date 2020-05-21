package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type TaxiPickUpMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.TaxiPickUpEvent
}

func (m *TaxiPickUpMessage) Key() string {
	return "taxi.pickup"
}
