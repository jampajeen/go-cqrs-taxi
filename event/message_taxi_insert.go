package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type TaxiInsertMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.Taxi
}

func (m *TaxiInsertMessage) Key() string {
	return "taxi.insert"
}
