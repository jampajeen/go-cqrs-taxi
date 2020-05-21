package main

import (
	"net/http"

	"github.com/jampajeen/go-cqrs-taxi/db"
	"github.com/jampajeen/go-cqrs-taxi/event"
	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	"github.com/jampajeen/go-cqrs-taxi/schema"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func broadcastHandler(c echo.Context) error {

	type request struct {
		IDUser string `json:"id_user"`
		Body   string `json:"body" validate:"required"`
	}

	type response struct {
		ID string `json:"id"`
	}

	// ctx := c.Request().Context()

	payload := new(request)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid body",
		})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	cmd := event.ServerCommand{
		IDUser: payload.IDUser,
		Body:   payload.Body,
	}
	if err := event.PublishServerCommand(cmd); err != nil {
		Log.Error(err)
	}

	return c.JSON(http.StatusOK, response{ID: id.String()})
}

func updateTaxiLocationHandler(c echo.Context) error {

	type request struct {
		ID  string  `json:"id" validate:"required"`
		Lat float64 `json:"lat" validate:"required"`
		Lon float64 `json:"lon" validate:"required"`
	}

	type response struct {
		ID string `json:"id"`
	}

	ctx := c.Request().Context()

	payload := new(request)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid body",
		})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	found := schema.Taxi{}
	if err := db.Find(ctx, payload.ID, &found); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Failed to search taxi in DB",
		})
	}

	// Publish event
	if err := event.PublishTaxiLocationUpdated(found); err != nil {
		Log.Error(err)
	}

	return c.JSON(http.StatusOK, response{ID: id.String()})
}

func insertTaxiHandler(c echo.Context) error {

	type request struct {
		Taxi schema.Taxi `json:"taxi"`
	}

	type response struct {
		ID string `json:"id"`
	}

	ctx := c.Request().Context()

	payload := new(request)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid body",
		})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	taxi := payload.Taxi
	taxi.ID = id.String()

	if err := db.Insert(ctx, taxi); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create taxi in DB",
		})
	}

	// Publish event
	if err := event.PublishTaxiInserted(taxi); err != nil {
		Log.Error(err)
	}

	return c.JSON(http.StatusOK, response{ID: taxi.ID})
}

func updateTaxiHandler(c echo.Context) error {

	type request struct {
		Taxi schema.Taxi `json:"taxi"`
	}

	type response struct {
		ID   string      `json:"id"`
		Taxi schema.Taxi `json:"taxi"`
	}

	ctx := c.Request().Context()

	payload := new(request)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	taxi := payload.Taxi

	_, err := uuid.FromString(taxi.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No object id found",
		})
	}

	found := schema.Taxi{}
	if err := db.Find(ctx, taxi.ID, &found); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Failed to search taxi in DB",
		})
	}

	if err := db.Update(ctx, taxi); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update taxi in DB",
		})
	}

	// Publish event
	if err := event.PublishTaxiUpdated(taxi); err != nil {
		Log.Error(err)
	}

	return c.JSON(http.StatusOK, response{ID: taxi.ID})
}

func taxiCancelProposalHandler(c echo.Context) error {
	type response struct {
		ID                      string                         `json:"id"`
		TaxiCancelProposalEvent schema.TaxiCancelProposalEvent `json:"taxi_cancel_proposal_event"`
	}

	ctx := c.Request().Context()

	request := new(schema.TaxiCancelProposalEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	request.ID = id.String()

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create record in DB",
		})
	}

	// TODO: logic

	return c.JSON(http.StatusOK, &response{ID: id.String(), TaxiCancelProposalEvent: *request})

}
func taxiDropOffHandler(c echo.Context) error {
	type response struct {
		ID               string                  `json:"id"`
		TaxiDropOffEvent schema.TaxiDropOffEvent `json:"taxi_drop_off_event"`
	}

	ctx := c.Request().Context()

	request := new(schema.TaxiDropOffEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	request.ID = id.String()

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create record DB",
		})
	}

	// TODO: logic

	return c.JSON(http.StatusOK, &response{ID: id.String(), TaxiDropOffEvent: *request})
}

func taxiPickUpHandler(c echo.Context) error {
	type response struct {
		ID              string                 `json:"id"`
		TaxiPickUpEvent schema.TaxiPickUpEvent `json:"taxi_pick_up_event"`
	}

	ctx := c.Request().Context()

	request := new(schema.TaxiPickUpEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	request.ID = id.String()

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create record DB",
		})
	}

	// TODO: logic

	return c.JSON(http.StatusOK, &response{ID: id.String(), TaxiPickUpEvent: *request})
}

func taxiProposeHandler(c echo.Context) error {
	type response struct {
		ID               string                  `json:"id"`
		TaxiProposeEvent schema.TaxiProposeEvent `json:"taxi_propose_event"`
	}

	ctx := c.Request().Context()

	request := new(schema.TaxiProposeEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	request.ID = id.String()

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create record DB",
		})
	}

	// TODO: logic

	return c.JSON(http.StatusOK, &response{ID: id.String(), TaxiProposeEvent: *request})
}

func userAcceptProposalHandler(c echo.Context) error {
	type response struct {
		ID                      string                         `json:"id"`
		UserAcceptProposalEvent schema.UserAcceptProposalEvent `json:"user_accept_proposal_event"`
	}

	ctx := c.Request().Context()

	request := new(schema.UserAcceptProposalEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	request.ID = id.String()

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create record DB",
		})
	}

	// TODO: logic

	return c.JSON(http.StatusOK, &response{ID: id.String(), UserAcceptProposalEvent: *request})
}
func userCancelBookingHandler(c echo.Context) error {
	type response struct {
		ID                     string                        `json:"id"`
		UserCancelBookingEvent schema.UserCancelBookingEvent `json:"user_cancel_booking_event"`
	}

	ctx := c.Request().Context()

	request := new(schema.UserCancelBookingEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	request.ID = id.String()

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create record DB",
		})
	}

	// TODO: logic

	return c.JSON(http.StatusOK, &response{ID: id.String(), UserCancelBookingEvent: *request})
}
func userCancelRequestBookingHandler(c echo.Context) error {
	type response struct {
		ID                            string                               `json:"id"`
		UserCancelRequestBookingEvent schema.UserCancelRequestBookingEvent `json:"user_cancel_request_booking_event"`
	}

	ctx := c.Request().Context()

	request := new(schema.UserCancelRequestBookingEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	request.ID = id.String()

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create record DB",
		})
	}

	// TODO: logic

	return c.JSON(http.StatusOK, &response{ID: id.String(), UserCancelRequestBookingEvent: *request})
}

func userRequestBookingHandler(c echo.Context) error {
	type response struct {
		ID                      string                         `json:"id"`
		UserRequestBookingEvent schema.UserRequestBookingEvent `json:"user_request_booking_event"`
	}

	ctx := c.Request().Context()

	request := new(schema.UserRequestBookingEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid body",
		})
	}

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}

	request.ID = id.String()

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create record DB",
		})
	}

	// TODO: logic

	return c.JSON(http.StatusOK, &response{ID: id.String(), UserRequestBookingEvent: *request})
}
