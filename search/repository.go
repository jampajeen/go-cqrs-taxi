package search

import (
	"context"

	"github.com/jampajeen/go-cqrs-taxi/schema"
)

type Repository interface {
	Close()
	UpdateTaxiLocation(ctx context.Context, id string, lat float64, lon float64) error
	InsertTaxi(ctx context.Context, taxi schema.Taxi) error
	SearchTaxies(ctx context.Context, query string, start uint64, size uint64) ([]schema.Taxi, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func UpdateTaxiLocation(ctx context.Context, id string, lat float64, lon float64) error {
	return impl.UpdateTaxiLocation(ctx, id, lat, lon)
}

func InsertTaxi(ctx context.Context, taxi schema.Taxi) error {
	return impl.InsertTaxi(ctx, taxi)
}

func SearchTaxies(ctx context.Context, query string, start uint64, size uint64) ([]schema.Taxi, error) {
	return impl.SearchTaxies(ctx, query, start, size)
}
