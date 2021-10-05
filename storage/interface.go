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
	URL string
)

type Storage interface {
	PutURL(ctx context.Context, url URL) (Key, error)
	GetURL(ctx context.Context, key Key) (URL, error)
}
