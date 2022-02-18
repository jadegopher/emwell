package repository

import "context"

type LinkRepository interface {
	Insert(ctx context.Context, key string, value []byte) (err error)
	GetByID(ctx context.Context, key string) (value []byte, err error)
}
