package main

import (
	"encoding/json"
	"fmt"

	"github.com/jampajeen/go-cqrs-taxi/event"
	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}

func handleEvents(es *event.NatsEventBroker, hub *Hub) {

	err := es.OnServerCommand(func(m event.ServerCommandMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		if m.Event.IDUser != "" {
			hub.sendToIDUser(m, m.Event.IDUser)
		} else {
			hub.sendBroadcast(m, nil)
		}

	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnTaxiCancelProposal(func(m event.TaxiCancelProposalMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		// hub.broadcast(event.TaxiCancelProposalMessage{
		// 	// TODO:
		// }, nil)
	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnTaxiDropOff(func(m event.TaxiDropOffMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		// hub.broadcast(event.TaxiDropOffMessage{
		// 	// TODO:
		// }, nil)
	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnTaxiPickUp(func(m event.TaxiPickUpMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		// hub.broadcast(event.TaxiPickUpMessage{
		// 	// TODO:
		// }, nil)
	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnTaxiPropose(func(m event.TaxiProposeMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		// hub.broadcast(event.TaxiProposeMessage{
		// 	// TODO:
		// }, nil)
	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnUserAcceptProposal(func(m event.UserAcceptProposalMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		// hub.broadcast(event.UserAcceptProposalMessage{
		// 	// TODO:
		// }, nil)
	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnUserCancelBooking(func(m event.UserCancelBookingMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		// hub.broadcast(event.UserCancelBookingMessage{
		// 	// TODO:
		// }, nil)
	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnUserCancelRequestBooking(func(m event.UserCancelRequestBookingMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		// hub.broadcast(event.UserCancelBookingMessage{
		// 	// TODO:
		// }, nil)
	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnUserRequestBooking(func(m event.UserRequestBookingMessage) {
		data, _ := json.Marshal(m)
		Log.Info(string(data))
		// hub.broadcast(event.UserRequestBookingMessage{
		// 	// TODO:
		// }, nil)
	})
	if err != nil {
		Log.Error(err)
	}
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		Log.Fatal(err)
	}

	hub := newHub()

	es, err := event.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))
	if err != nil {
		Log.Fatal(err)
	}

	handleEvents(es, hub)

	event.SetEventStore(es)
	defer event.Close()

	go hub.run()

	e := echo.New()
	// e.Use(middleware.JWT([]byte("secret")))
	e.GET("/pusher", hub.handleWebSocket)

	port := fmt.Sprintf(":%d", 8083)
	e.Logger.Fatal(e.Start(port))
}
