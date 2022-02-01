package consumer

import (
	"context"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/logger"
)

type Consumer interface {
	StartConsuming(ctx context.Context) chan entities.Update
}

type EventConsumer struct {
	logger   logger.ILogger
	bot      *tgbotapi.BotAPI
	wg       *sync.WaitGroup
	termCh   chan struct{}
	updateCh chan entities.Update
}

func NewEventConsumer(logger logger.ILogger, token string, wg *sync.WaitGroup) (*EventConsumer, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &EventConsumer{
		logger:   logger,
		bot:      bot,
		termCh:   make(chan struct{}, 1),
		updateCh: make(chan entities.Update, 10),
		wg:       wg,
	}, nil
}

func (e *EventConsumer) StartConsuming(ctx context.Context) chan entities.Update {
	e.bot.Debug = true
	e.logger.InfoKV(ctx, "Consuming updates in progress", "account", e.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := e.bot.GetUpdatesChan(u)

	e.wg.Add(1)
	go func(ctx context.Context) {
		defer e.wg.Done()
		for {
			select {
			case update := <-updates:
				e.logger.InfoKV(ctx, "Got new update", "update", update)
				upd, ok := e.convertUpdate(ctx, update)
				if ok {
					e.updateCh <- upd
				}
			case <-ctx.Done():
				e.logger.InfoKV(ctx, "Context Done signal received. Exiting telegram bot...")
				e.gracefulStop()
				return
			case <-e.termCh:
				e.gracefulStop()
				return
			default:
				continue
			}
		}
	}(ctx)

	return e.updateCh
}

func (e *EventConsumer) gracefulStop() {
	// TODO save offset
	e.bot.StopReceivingUpdates()
	close(e.updateCh)
}

func (e *EventConsumer) convertUpdate(ctx context.Context, update tgbotapi.Update) (entities.Update, bool) {
	result := entities.Update(nil)
	err := error(nil)

	if update.Message != nil {
		if update.Message.From == nil && update.Message.Chat == nil {
			return nil, false
		}
		result, err = entities.NewUpdate(
			int64(update.UpdateID),
			entities.UpdateTypeMessage,
			entities.Sender{
				ID:           update.Message.From.ID,
				FirstName:    update.Message.From.FirstName,
				LastName:     update.Message.From.LastName,
				UserName:     update.Message.From.UserName,
				LanguageCode: update.Message.From.LanguageCode,
			},
			update.Message.Chat.ID,
			nil,
			entities.Message{
				ID: int64(update.Message.MessageID),
				From: entities.Sender{
					ID:           update.Message.From.ID,
					FirstName:    update.Message.From.FirstName,
					LastName:     update.Message.From.LastName,
					UserName:     update.Message.From.UserName,
					LanguageCode: update.Message.From.LanguageCode,
				},
				Chat: entities.Chat{
					ID: update.Message.Chat.ID,
				},
				Text: update.Message.Text,
			},
		)
		if err != nil {
			e.logger.ErrorKV(ctx, "Error convert update", "err", err)
			return nil, false
		}
	}

	if result == nil {
		return nil, false
	}

	return result, true
}
