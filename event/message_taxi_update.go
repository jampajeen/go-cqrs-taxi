package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type TaxiUpdateMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.Taxi
}

func (m *TaxiUpdateMessage) Key() string {
	return "taxi.update"
}
