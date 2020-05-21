package event

import (
	"time"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type UserRequestBookingMessage struct {
	ID        string
	CreatedAt time.Time
	Event     schema.UserRequestBookingEvent
}

func (m *UserRequestBookingMessage) Key() string {
	return "user.request.booking"
}
