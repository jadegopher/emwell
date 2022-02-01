package repository

import (
	"context"

	"emwell/internal/user/entites"
)

type UserRepository interface {
	Insert(ctx context.Context, user entites.User) (id int64, err error)
}
