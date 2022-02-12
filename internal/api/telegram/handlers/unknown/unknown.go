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

func (h *Handler) CanHandle(upd entities.Update) bool {
	_, ok := upd.Payload().(entities.Message)
	if !ok {
		return false
	}

	return true
}

func (h *Handler) Handle(_ context.Context, update entities.Update) (resp []handlers.Response, err error) {
	if !h.CanHandle(update) {
		return nil, handlers.ErrCantHandle
	}

	msg, _ := update.Payload().(entities.Message)

	return []handlers.Response{
		handlers.NewMessage(
			handlers.MessagePayload{
				ToChatID: msg.Chat.ID,
				Text:     "Извини, я не понимаю что ты говоришь. Попробуй воспользоваться списком доступных команд",
			},
		),
	}, nil
}
