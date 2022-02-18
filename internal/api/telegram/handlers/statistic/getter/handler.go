package getter

import (
	"context"
	"fmt"
	"time"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/api/telegram/handlers"
	"emwell/internal/core/diary"
	"emwell/internal/core/link"
)

type Handler struct {
	diaryService *diary.Diary
	linkService  *link.Service
	handlerType  string
	siteURL      string
	gapYear      int
	gapMonth     int
	gapDay       int
}

func NewStatisticGetterHandler(
	diaryService *diary.Diary,
	linkService *link.Service,
	handlerType string,
	siteURL string,
	gapYear int,
	gapMonth int,
	gapDay int,
) *Handler {
	return &Handler{
		diaryService: diaryService,
		linkService:  linkService,
		handlerType:  handlerType,
		siteURL:      siteURL,
		gapYear:      gapYear,
		gapMonth:     gapMonth,
		gapDay:       gapDay,
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

	to := time.Now().UTC()
	from := to.AddDate(h.gapYear, h.gapMonth, h.gapDay)
	stats, err := h.diaryService.GetStatistics(ctx, user.ID(), from, to)
	if err != nil {
		return nil, err
	}

	chart, err := stats.Visualize()
	if err != nil {
		return nil, err
	}

	pwd, err := h.linkService.SaveContent(ctx, user.ID(), chart)
	if err != nil {
		return nil, err
	}

	return []handlers.Response{
		handlers.NewMessage(
			handlers.MessagePayload{
				ToChatID: callback.Message.Chat.ID,
				Text: fmt.Sprintf(
					"График досутпен по ссылке: %s/emotional/statistics?password=%s",
					h.siteURL, pwd,
				),
			},
		),
	}, nil
}
