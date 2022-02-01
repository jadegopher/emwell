package middlewares

import (
	"context"

	"emwell/internal/api/telegram/consumer/entities"
)

type Middleware interface {
	Serve(ctx context.Context, upd entities.Update) (entities.Update, error)
}
