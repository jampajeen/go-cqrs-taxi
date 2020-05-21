package event

import (
	"time"
)

type ServerCommand struct {
	IDUser string `json:"id_user"`
	Body   string `json:"body" validate:"required"`
}

type ServerCommandMessage struct {
	ID        string
	CreatedAt time.Time
	Event     ServerCommand
}

func (m *ServerCommandMessage) Key() string {
	return "server.command"
}
