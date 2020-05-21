package event

import "github.com/jampajeen/go-cqrs-taxi/schema"

type EventBroker interface {
	Close()

	PublishTaxiCancelProposal(data schema.TaxiCancelProposalEvent) error
	PublishTaxiDropOff(data schema.TaxiDropOffEvent) error
	PublishTaxiPickUp(data schema.TaxiPickUpEvent) error
	PublishTaxiPropose(data schema.TaxiProposeEvent) error
	PublishUserAcceptProposal(data schema.UserAcceptProposalEvent) error
	PublishUserCancelBooking(data schema.UserCancelBookingEvent) error
	PublishUserCancelRequestBooking(data schema.UserCancelRequestBookingEvent) error
	PublishUserRequestBooking(data schema.UserRequestBookingEvent) error

	PublishTaxiInserted(data schema.Taxi) error
	PublishTaxiUpdated(data schema.Taxi) error
	PublishTaxiLocationUpdated(data schema.Taxi) error
	PublishServerCommand(data ServerCommand) error

	OnTaxiCancelProposal(f func(TaxiCancelProposalMessage)) error
	OnTaxiDropOff(f func(TaxiDropOffMessage)) error
	OnTaxiPickUp(f func(TaxiPickUpMessage)) error
	OnTaxiPropose(f func(TaxiProposeMessage)) error
	OnUserAcceptProposal(f func(UserAcceptProposalMessage)) error
	OnUserCancelBooking(f func(UserCancelBookingMessage)) error
	OnUserCancelRequestBooking(f func(UserCancelRequestBookingMessage)) error
	OnUserRequestBooking(f func(UserRequestBookingMessage)) error

	OnTaxiInserted(f func(TaxiInsertMessage)) error
	OnTaxiUpdated(f func(TaxiUpdateMessage)) error
	OnTaxiLocationUpdated(f func(TaxiLocationUpdateMessage)) error
	OnServerCommand(f func(ServerCommandMessage)) error
}

var impl EventBroker

func SetEventStore(es EventBroker) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishTaxiInserted(data schema.Taxi) error {
	return impl.PublishTaxiInserted(data)
}

func PublishTaxiUpdated(data schema.Taxi) error {
	return impl.PublishTaxiUpdated(data)
}

func PublishTaxiLocationUpdated(data schema.Taxi) error {
	return impl.PublishTaxiLocationUpdated(data)
}

func PublishServerCommand(data ServerCommand) error {
	return impl.PublishServerCommand(data)
}

func OnTaxiCancelProposal(f func(TaxiCancelProposalMessage)) error {
	return impl.OnTaxiCancelProposal(f)
}
func OnTaxiDropOff(f func(TaxiDropOffMessage)) error {
	return impl.OnTaxiDropOff(f)
}
func OnTaxiPickUp(f func(TaxiPickUpMessage)) error {
	return impl.OnTaxiPickUp(f)
}
func OnTaxiPropose(f func(TaxiProposeMessage)) error {
	return impl.OnTaxiPropose(f)
}
func OnUserAcceptProposal(f func(UserAcceptProposalMessage)) error {
	return impl.OnUserAcceptProposal(f)
}
func OnUserCancelBooking(f func(UserCancelBookingMessage)) error {
	return impl.OnUserCancelBooking(f)
}
func OnUserCancelRequestBooking(f func(UserCancelRequestBookingMessage)) error {
	return impl.OnUserCancelRequestBooking(f)
}
func OnUserRequestBooking(f func(UserRequestBookingMessage)) error {
	return impl.OnUserRequestBooking(f)
}
func OnTaxiInserted(f func(TaxiInsertMessage)) error {
	return impl.OnTaxiInserted(f)
}
func OnTaxiUpdated(f func(TaxiUpdateMessage)) error {
	return impl.OnTaxiUpdated(f)
}
func OnTaxiLocationUpdated(f func(TaxiLocationUpdateMessage)) error {
	return impl.OnTaxiLocationUpdated(f)
}
func OnServerCommand(f func(ServerCommandMessage)) error {
	return impl.OnServerCommand(f)
}
