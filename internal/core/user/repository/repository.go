package repository

import (
	"context"

	"emwell/internal/core/user/entites"
)

type UserRepository interface {
	Insert(ctx context.Context, user entites.User) (id int64, err error)
}
