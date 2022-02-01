package unknown

import (
	"context"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/api/telegram/handlers"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CanHandle(_ entities.Update) bool {
	return true
}

func (h *Handler) Handle(_ context.Context, update entities.Update) (resp []handlers.Response, err error) {
	if !h.CanHandle(update) {
		return nil, handlers.ErrCantHandle
	}

	return []handlers.Response{
		{
			Text: "Извини, я не понимаю что ты говоришь. Попробуй воспользоваться списком доступных команд",
		},
	}, nil
}
