package main

import (
	"encoding/json"

	"net/http"
	"time"

	"github.com/gorilla/websocket"
	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	clients     map[*Client]bool
	userClients map[string]*Client
	broadcast   chan []byte
	register    chan *Client
	unregister  chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:   make(chan []byte),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
		userClients: make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.onConnected(client)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.userClients, client.IDUser)
				delete(h.clients, client)
				close(client.send)
				h.onDisconnected(client)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
					h.onSent(client, message)
				default:
					close(client.send)
					delete(h.userClients, client.IDUser)
					delete(h.clients, client)
					h.onDisconnected(client)
				}
			}
		}

	}
}

func (hub *Hub) sendBroadcast(message interface{}, ignore *Client) {
	data, _ := json.Marshal(message)
	hub.broadcast <- data
}

func (hub *Hub) send(message interface{}, client *Client) {
	data, _ := json.Marshal(message)
	client.send <- data
	hub.onSent(client, data)
}

func (hub *Hub) sendToIDUser(message interface{}, id string) {
	data, _ := json.Marshal(message)
	if client, ok := hub.userClients[id]; ok {
		client.send <- data
		hub.onSent(client, data)
	}
}

func (hub *Hub) handleWebSocket(c echo.Context) error {
	socket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		Log.Error(err)
		return err
	}

	client := &Client{hub: hub, conn: socket, send: make(chan []byte)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()

	return err
}

func (hub *Hub) handleCommand(client *Client, cmd RequestDto) {
	var request struct {
		IDUser  string
		AppKind string
	}

	if cmd.Cmd == CmdWhoAreYouResponse {
		jsonbody, err := json.Marshal(cmd.Data)

		err = json.Unmarshal(jsonbody, &request)
		if err != nil {
			Log.Warn("client implementation error")
		}

		Log.Info("IDUser = %s, AppKind = %s", request.IDUser, request.AppKind)
		client.IDUser = request.IDUser
		client.AppKind = request.AppKind

		hub.userClients[client.IDUser] = client
	}

	if cmd.Cmd == CmdAreYouThereRequest {
		createdAt := time.Now().UTC()
		id, err := uuid.NewV4()
		if err == nil {
			cmd := RequestDto{
				Cmd:       CmdAreYouThereResponse,
				ID:        id.String(),
				IDRelated: cmd.ID, // request ID
				CreatedAt: createdAt,
			}
			hub.send(cmd, client)
		}

	}

}

func (hub *Hub) onReceived(client *Client, data []byte) {
	Log.Info("client received %s", string(data))

	var request RequestDto
	err := json.Unmarshal(data, &request)
	if err != nil {
		Log.Warn("client implementation error")
	}

	hub.handleCommand(client, request)
}

func (hub *Hub) onSent(client *Client, data []byte) {
	Log.Info("client sent %s", string(data))
}

func (hub *Hub) onConnected(client *Client) {
	Log.Info("client connected: ", client.conn.RemoteAddr())

	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		hub.unregister <- client
		return
	}
	cmd := RequestDto{
		Cmd:       CmdWhoAreYouRequest,
		ID:        id.String(),
		CreatedAt: createdAt,
	}
	hub.send(cmd, client)
}

func (hub *Hub) onDisconnected(client *Client) {
	Log.Info("client disconnected: ", client.conn.RemoteAddr())
}
