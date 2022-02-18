package daily_routine

import (
	"context"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/api/telegram/handlers"
)

const (
	DailyRoutineWorst   = "daily_routine.worst"
	DailyRoutineWorse   = "daily_routine.worse"
	DailyRoutineBad     = "daily_routine.bad"
	DailyRoutineNeutral = "daily_routine.neutral"
	DailyRoutineGood    = "daily_routine.good"
	DailyRoutineBetter  = "daily_routine.better"
	DailyRoutineBest    = "daily_routine.best"
)

type Handler struct{}

func NewDailyRoutineHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CanHandle(upd entities.Update) bool {
	msg, ok := upd.Payload().(entities.Message)
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

	msg, _ := upd.Payload().(entities.Message)

	return []handlers.Response{
		handlers.NewMessage(
			handlers.MessagePayload{
				ToChatID: msg.Chat.ID,
				Text:     "Скажи, как ты оцениваешь свой день?",
				InlineKeyboard: [][]handlers.Button{
					{
						{
							Text: "🤕",
							Data: DailyRoutineWorst,
						},
						{
							Text: "😪",
							Data: DailyRoutineWorse,
						},
						{
							Text: "😔",
							Data: DailyRoutineBad,
						},
						{
							Text: "😐",
							Data: DailyRoutineNeutral,
						},
						{
							Text: "😌",
							Data: DailyRoutineGood,
						},
						{
							Text: "☺️",
							Data: DailyRoutineBetter,
						},
						{
							Text: "😎",
							Data: DailyRoutineBest,
						},
					},
				},
			},
		),
	}, nil
}
