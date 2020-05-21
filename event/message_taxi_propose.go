package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type TaxiProposeMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.TaxiProposeEvent
}

func (m *TaxiProposeMessage) Key() string {
	return "taxi.propose"
}
