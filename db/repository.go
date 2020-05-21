package db

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, item interface{}) error
	Update(ctx context.Context, item interface{}) error
	Delete(ctx context.Context, id string, item interface{}) error
	Find(ctx context.Context, id string, out interface{}) error
	FindAll(ctx context.Context, out interface{}) error
	Close()
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func Insert(ctx context.Context, item interface{}) error {
	return impl.Insert(ctx, item)
}
func Update(ctx context.Context, item interface{}) error {
	return impl.Update(ctx, item)
}
func Delete(ctx context.Context, id string, item interface{}) error {
	return impl.Delete(ctx, id, item)
}
func Find(ctx context.Context, id string, out interface{}) error {
	return impl.Find(ctx, id, out)
}
func FindAll(ctx context.Context, out interface{}) error {
	return impl.FindAll(ctx, out)
}
