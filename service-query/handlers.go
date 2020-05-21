package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jampajeen/go-cqrs-taxi/db"
	"github.com/jampajeen/go-cqrs-taxi/event"
	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	"github.com/jampajeen/go-cqrs-taxi/schema"
	"github.com/jampajeen/go-cqrs-taxi/search"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func importTaxies() {
	Log.Info("Start importing taxies data")

	taxies := []schema.Taxi{}
	if err := db.FindAll(context.Background(), &taxies); err != nil {
		Log.Error(err)
		panic("Error on getting all taxies from DB")
	}

	for _, taxi := range taxies {
		if err := search.InsertTaxi(context.Background(), taxi); err != nil {
			Log.Error(err)
		}
	}
	Log.Info("Done importing taxies data")
}

func handleEvents(es *event.NatsEventBroker) {

	err := es.OnTaxiInserted(func(m event.TaxiInsertMessage) {
		data, _ := json.Marshal(m)
		Log.Info("Taxi Inserted: ", string(data))

		if err := search.InsertTaxi(context.Background(), m.Event); err != nil {
			Log.Error(err)
		}
	})
	if err != nil {
		Log.Error(err)
	}

	err = es.OnTaxiLocationUpdated(func(m event.TaxiLocationUpdateMessage) {
		data, _ := json.Marshal(m)
		Log.Info("Taxi Location Updated: ", string(data))

		if err := search.UpdateTaxiLocation(context.Background(), m.Event.ID, m.Event.Lat, m.Event.Lon); err != nil {
			Log.Error(err)
		}
	})
	if err != nil {
		Log.Error(err)
	}
}

func bookingsHandler(c echo.Context) error {

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}
	item := schema.Booking{
		ID: id.String(),
	}
	return c.JSON(http.StatusOK, item)
}

func paymentsHandler(c echo.Context) error {

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}
	item := schema.Payment{
		ID: id.String(),
	}
	return c.JSON(http.StatusOK, item)
}

func taxiesHandler(c echo.Context) error {

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}
	item := schema.Taxi{
		ID: id.String(),
	}
	return c.JSON(http.StatusOK, item)
}

func driversHandler(c echo.Context) error {

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}
	item := schema.Driver{
		ID: id.String(),
	}
	return c.JSON(http.StatusOK, item)
}

func usersHandler(c echo.Context) error {

	id, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate uuid",
		})
	}
	item := schema.User{
		ID: id.String(),
	}
	return c.JSON(http.StatusOK, item)
}

func testHandler(c echo.Context) error {

	type response struct {
		ID   string      `json:"id"`
		Out  interface{} `json:"out"`
		Out2 interface{} `json:"out2"`
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

	// request.ID = "1"

	if err := db.Insert(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create taxi in DB",
		})
	}

	request.Lat = 88
	if err := db.Update(ctx, request); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create taxi in DB",
		})
	}

	if err := db.Delete(ctx, "2", schema.TaxiCancelProposalEvent{}); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create taxi in DB",
		})
	}

	out := schema.TaxiCancelProposalEvent{}
	if err := db.Find(ctx, "1", &out); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create taxi in DB",
		})
	}

	out2 := []schema.TaxiCancelProposalEvent{}
	if err := db.FindAll(ctx, &out2); err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create taxi in DB",
		})
	}

	return c.JSON(http.StatusOK, &response{ID: id.String(), Out: out, Out2: out2})
}

func searchTaxiesHandler(c echo.Context) error {

	ctx := c.Request().Context()
	var err error

	query := c.FormValue("query")
	if len(query) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Missing query parameter",
		})
	}

	start := uint64(0)
	startStr := c.FormValue("start")
	size := uint64(100)
	sizeStr := c.FormValue("size")
	if len(startStr) != 0 {
		start, err = strconv.ParseUint(startStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid start parameter",
			})
		}
	}

	if len(sizeStr) != 0 {
		size, err = strconv.ParseUint(sizeStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid size parameter",
			})
		}
	}

	result, err := search.SearchTaxies(ctx, query, start, size)
	if err != nil {
		Log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
