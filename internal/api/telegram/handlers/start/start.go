package start

import (
	"context"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/api/telegram/handlers"
)

type Handler struct{}

func NewMenuHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CanHandle(upd entities.Update) bool {
	msg, ok := upd.Payload().(entities.Message)
	if !ok {
		return false
	}

	if msg.Text == "/start" {
		return true
	}

	return false
}

func (h *Handler) Handle(_ context.Context, upd entities.Update) ([]handlers.Response, error) {
	if !h.CanHandle(upd) {
		return nil, handlers.ErrCantHandle
	}

	return nil, nil
}
