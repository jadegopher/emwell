package statistic

import (
	"context"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/api/telegram/handlers"
)

type Handler struct{}

const (
	EmotionalStatisticsForMonth = "get_emotional_statistics.month"
	EmotionalStatisticsForWeek  = "get_emotional_statistics.week"
)

func NewStatisticHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CanHandle(upd entities.Update) bool {
	msg, ok := upd.Payload().(entities.Message)
	if !ok {
		return false
	}

	if msg.Text == "/get_emotional_statistics" {
		return true
	}

	return false
}

func (h *Handler) Handle(_ context.Context, upd entities.Update) ([]handlers.Response, error) {
	if !h.CanHandle(upd) {
		return nil, handlers.ErrCantHandle
	}

	msg, _ := upd.Payload().(entities.Message)

	return []handlers.Response{
		handlers.NewMessage(handlers.MessagePayload{
			ToChatID: msg.Chat.ID,
			Text:     "За какой период хочешь получить статистику?",
			InlineKeyboard: [][]handlers.Button{
				{
					{
						Text: "За месяц",
						Data: EmotionalStatisticsForMonth,
					},
					{
						Text: "За неделю",
						Data: EmotionalStatisticsForWeek,
					},
				},
			},
		}),
	}, nil
}
