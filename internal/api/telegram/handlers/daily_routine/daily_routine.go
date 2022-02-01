package daily_routine

import (
	"context"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/api/telegram/handlers"
)

type Handler struct{}

func NewDailyRoutineHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CanHandle(upd entities.Update) bool {
	msg, ok := upd.Message()
	if !ok {
		return false
	}

	if msg.Text == "/daily_routine" {
		return true
	}

	return false
}

func (h *Handler) Handle(_ context.Context, upd entities.Update) ([]handlers.Response, error) {
	if !h.CanHandle(upd) {
		return nil, handlers.ErrCantHandle
	}

	return []handlers.Response{
		{
			Text: "Скажи, как ты оцениваешь свой день?",
			Buttons: [][]handlers.Button{
				{
					{
						Text: "🤕",
						Data: "daily_routine.worst",
					},
					{
						Text: "😪",
						Data: "daily_routine.worse",
					},
					{
						Text: "😔",
						Data: "daily_routine.bad",
					},
					{
						Text: "😌",
						Data: "daily_routine.good",
					},
					{
						Text: "☺️",
						Data: "daily_routine.better",
					},
					{
						Text: "😎",
						Data: "daily_routine.best",
					},
				},
			},
		},
	}, nil
}
