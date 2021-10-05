package storage

import (
	"context"
	"errors"
	"fmt"
)

var (
	BaseError = errors.New("storage")
	CollisionError = fmt.Errorf("%w.collision", BaseError)
	NotFoundError = fmt.Errorf("%w.not_found", BaseError)
)

type (
	Key string
	Url string
)

type Storage interface {
	PutUrl(ctx context.Context, url Url) (Key, error)
	GetUrl(ctx context.Context, key Key) (Url, error)
}
