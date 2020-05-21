package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type UserAcceptProposalMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.UserAcceptProposalEvent
}

func (m *UserAcceptProposalMessage) Key() string {
	return "user.accept.proposal"
}
