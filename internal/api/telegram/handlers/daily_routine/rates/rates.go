package rates

import (
	"context"
	"time"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/api/telegram/handlers"
	"emwell/internal/core/diary"
	"emwell/internal/core/diary/entites"
)

type Handler struct {
	diaryService  *diary.Diary
	handlerType   string
	emotionalRate int32
	text          string
}

func NewDailyRoutineEmotionalHandler(
	handlerType string,
	emotionalRate int32,
	text string,
	diary *diary.Diary,
) *Handler {
	return &Handler{
		handlerType:   handlerType,
		diaryService:  diary,
		emotionalRate: emotionalRate,
		text:          text,
	}
}

func (h *Handler) CanHandle(upd entities.Update) bool {
	callback, ok := upd.Payload().(entities.Callback)
	if !ok {
		return false
	}

	if callback.Data == h.handlerType {
		return true
	}

	return false
}

func (h *Handler) Handle(ctx context.Context, upd entities.Update) ([]handlers.Response, error) {
	if !h.CanHandle(upd) {
		return nil, handlers.ErrCantHandle
	}

	user, ok := upd.User()
	if !ok {
		return nil, handlers.ErrUserNotSpecified
	}

	callback, _ := upd.Payload().(entities.Callback)

	if _, err := h.diaryService.SaveEmotionalInformation(
		ctx,
		entites.EmotionalInfo{
			UserID:        user.ID(),
			EmotionalRate: h.emotionalRate,
			ReferToDate:   time.Time{},
		},
	); err != nil {
		return nil, err
	}

	return []handlers.Response{
		handlers.NewMessage(
			handlers.MessagePayload{
				ToChatID: callback.Message.Chat.ID,
				Text:     h.text,
			},
		),
	}, nil
}
