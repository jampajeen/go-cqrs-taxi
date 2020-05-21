package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type UserCancelRequestBookingMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.UserCancelRequestBookingEvent
}

func (m *UserCancelRequestBookingMessage) Key() string {
	return "user.cancel.request.booking"
}
