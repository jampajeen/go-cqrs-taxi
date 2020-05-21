package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type TaxiCancelProposalMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.TaxiCancelProposalEvent
}

func (m *TaxiCancelProposalMessage) Key() string {
	return "taxi.cancel.proposal"
}
