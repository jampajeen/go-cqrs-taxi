package event

import (
	"bytes"
	"encoding/gob"
	"time"

	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	"github.com/jampajeen/go-cqrs-taxi/schema"
	"github.com/nats-io/go-nats"
	uuid "github.com/satori/go.uuid"
)

type NatsEventBroker struct {
	nc *nats.Conn

	onTaxiCancelProposalSubscription       *nats.Subscription
	onTaxiDropOffSubscription              *nats.Subscription
	onTaxiPickUpSubscription               *nats.Subscription
	onTaxiProposeSubscription              *nats.Subscription
	onUserAcceptProposalSubscription       *nats.Subscription
	onUserCancelBookingSubscription        *nats.Subscription
	onUserCancelRequestBookingSubscription *nats.Subscription
	onUserRequestBookingSubscription       *nats.Subscription
	onTaxiInsertedSubscription             *nats.Subscription
	onTaxiUpdatedSubscription              *nats.Subscription
	onTaxiLocationUpdatedSubscription      *nats.Subscription
	onServerCommandSubscription            *nats.Subscription
}

func NewNats(url string) (*NatsEventBroker, error) {
	Log.Info("Nats: %s", url)
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	gob.Register(TaxiCancelProposalMessage{})
	gob.Register(TaxiDropOffMessage{})
	gob.Register(TaxiPickUpMessage{})
	gob.Register(TaxiProposeMessage{})
	gob.Register(UserAcceptProposalMessage{})
	gob.Register(UserCancelBookingMessage{})
	gob.Register(UserCancelRequestBookingMessage{})
	gob.Register(UserRequestBookingMessage{})
	gob.Register(ServerCommandMessage{})

	return &NatsEventBroker{nc: nc}, nil
}

func (e *NatsEventBroker) Close() {
	if e.nc != nil {
		e.nc.Close()
	}

	if e.onTaxiCancelProposalSubscription != nil {
		e.onTaxiCancelProposalSubscription.Unsubscribe()
	}
	if e.onTaxiDropOffSubscription != nil {
		e.onTaxiDropOffSubscription.Unsubscribe()
	}
	if e.onTaxiPickUpSubscription != nil {
		e.onTaxiPickUpSubscription.Unsubscribe()
	}
	if e.onTaxiProposeSubscription != nil {
		e.onTaxiProposeSubscription.Unsubscribe()
	}
	if e.onUserAcceptProposalSubscription != nil {
		e.onUserAcceptProposalSubscription.Unsubscribe()
	}
	if e.onUserCancelBookingSubscription != nil {
		e.onUserCancelBookingSubscription.Unsubscribe()
	}
	if e.onUserCancelRequestBookingSubscription != nil {
		e.onUserCancelRequestBookingSubscription.Unsubscribe()
	}
	if e.onUserRequestBookingSubscription != nil {
		e.onUserRequestBookingSubscription.Unsubscribe()
	}

	if e.onTaxiInsertedSubscription != nil {
		e.onTaxiInsertedSubscription.Unsubscribe()
	}

	if e.onTaxiUpdatedSubscription != nil {
		e.onTaxiUpdatedSubscription.Unsubscribe()
	}

	if e.onTaxiLocationUpdatedSubscription != nil {
		e.onTaxiLocationUpdatedSubscription.Unsubscribe()
	}

	if e.onServerCommandSubscription != nil {
		e.onServerCommandSubscription.Unsubscribe()
	}
}

func (mq *NatsEventBroker) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	b.Reset()
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (mq *NatsEventBroker) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Reset()
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}

/*
* Implementations
 */
func (e *NatsEventBroker) PublishTaxiCancelProposal(ev schema.TaxiCancelProposalEvent) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := TaxiCancelProposalMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishTaxiDropOff(ev schema.TaxiDropOffEvent) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := TaxiDropOffMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishTaxiPickUp(ev schema.TaxiPickUpEvent) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := TaxiPickUpMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishTaxiPropose(ev schema.TaxiProposeEvent) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := TaxiProposeMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishUserAcceptProposal(ev schema.UserAcceptProposalEvent) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := UserAcceptProposalMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishUserCancelBooking(ev schema.UserCancelBookingEvent) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := UserCancelBookingMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishUserCancelRequestBooking(ev schema.UserCancelRequestBookingEvent) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := UserCancelRequestBookingMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishUserRequestBooking(ev schema.UserRequestBookingEvent) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := UserRequestBookingMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishTaxiInserted(ev schema.Taxi) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := TaxiInsertMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishTaxiUpdated(ev schema.Taxi) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := TaxiUpdateMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishTaxiLocationUpdated(ev schema.Taxi) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := TaxiLocationUpdateMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) PublishServerCommand(ev ServerCommand) error {
	createdAt := time.Now().UTC()
	id, err := uuid.NewV4()
	if err != nil {
		Log.Error(err)
		return err
	}

	m := ServerCommandMessage{
		ID:        id.String(),
		CreatedAt: createdAt,
		Event:     ev,
	}

	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

func (e *NatsEventBroker) OnTaxiCancelProposal(f func(TaxiCancelProposalMessage)) (err error) {
	m := TaxiCancelProposalMessage{}
	e.onTaxiCancelProposalSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := TaxiCancelProposalMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}

func (e *NatsEventBroker) OnTaxiDropOff(f func(TaxiDropOffMessage)) (err error) {
	m := TaxiDropOffMessage{}
	e.onTaxiDropOffSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := TaxiDropOffMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}
func (e *NatsEventBroker) OnTaxiPickUp(f func(TaxiPickUpMessage)) (err error) {
	m := TaxiPickUpMessage{}
	e.onTaxiPickUpSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := TaxiPickUpMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}
func (e *NatsEventBroker) OnTaxiPropose(f func(TaxiProposeMessage)) (err error) {
	m := TaxiProposeMessage{}
	e.onTaxiProposeSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := TaxiProposeMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}
func (e *NatsEventBroker) OnUserAcceptProposal(f func(UserAcceptProposalMessage)) (err error) {
	m := UserAcceptProposalMessage{}
	e.onUserAcceptProposalSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := UserAcceptProposalMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}
func (e *NatsEventBroker) OnUserCancelBooking(f func(UserCancelBookingMessage)) (err error) {
	m := UserCancelBookingMessage{}
	e.onUserCancelBookingSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := UserCancelBookingMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}
func (e *NatsEventBroker) OnUserCancelRequestBooking(f func(UserCancelRequestBookingMessage)) (err error) {
	m := UserCancelRequestBookingMessage{}
	e.onUserCancelRequestBookingSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := UserCancelRequestBookingMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}
func (e *NatsEventBroker) OnUserRequestBooking(f func(UserRequestBookingMessage)) (err error) {
	m := UserRequestBookingMessage{}
	e.onUserRequestBookingSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := UserRequestBookingMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}

func (e *NatsEventBroker) OnTaxiInserted(f func(TaxiInsertMessage)) (err error) {
	m := TaxiInsertMessage{}
	e.onTaxiInsertedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := TaxiInsertMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}

func (e *NatsEventBroker) OnTaxiUpdated(f func(TaxiUpdateMessage)) (err error) {
	m := TaxiUpdateMessage{}
	e.onTaxiUpdatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := TaxiUpdateMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}

func (e *NatsEventBroker) OnTaxiLocationUpdated(f func(TaxiLocationUpdateMessage)) (err error) {
	m := TaxiLocationUpdateMessage{}
	e.onTaxiLocationUpdatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := TaxiLocationUpdateMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}

func (e *NatsEventBroker) OnServerCommand(f func(ServerCommandMessage)) (err error) {
	m := ServerCommandMessage{}
	e.onServerCommandSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		b := ServerCommandMessage{}
		e.readMessage(msg.Data, &b)
		f(b)
	})
	return
}
