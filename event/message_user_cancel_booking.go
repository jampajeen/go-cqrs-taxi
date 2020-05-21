package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type UserCancelBookingMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.UserCancelBookingEvent
}

func (m *UserCancelBookingMessage) Key() string {
	return "user.cancel.booking"
}
