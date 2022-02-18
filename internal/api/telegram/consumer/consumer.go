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
			case update, isOpened := <-updates:
				if !isOpened {
					return
				}
				e.logger.InfoKV(ctx, "Got new update", "update", update)
				upd, ok := e.convertUpdate(ctx, update)
				if ok {
					e.updateCh <- upd
				}
			case <-ctx.Done():
				e.logger.InfoKV(ctx, "Context Done signal received. Exiting telegram bot...")
				e.gracefulStop()
				return
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
	if update.Message != nil {
		return e.toMessage(ctx, update)
	}

	if update.CallbackQuery != nil {
		return e.toCallback(ctx, update)
	}

	return nil, false
}

func (e *EventConsumer) toMessage(ctx context.Context, update tgbotapi.Update) (entities.Update, bool) {
	if update.Message.From == nil && update.Message.Chat == nil {
		return nil, false
	}
	result, err := entities.NewUpdate(
		int64(update.UpdateID),
		entities.UpdateTypeMessage,
		fromToSender(update.Message.From),
		nil,
		convertMessage(update.Message),
	)
	if err != nil {
		e.logger.ErrorKV(ctx, "Error convert update", "err", err)
		return nil, false
	}

	return result, true
}

func (e *EventConsumer) toCallback(ctx context.Context, update tgbotapi.Update) (entities.Update, bool) {
	if update.CallbackQuery.From == nil {
		return nil, false
	}

	result, err := entities.NewUpdate(
		int64(update.UpdateID),
		entities.UpdateTypeCallback,
		fromToSender(update.CallbackQuery.From),
		nil,
		entities.Callback{
			ID:              update.CallbackQuery.ID,
			From:            fromToSender(update.CallbackQuery.From),
			Message:         convertMessage(update.CallbackQuery.Message),
			InlineMessageID: update.CallbackQuery.InlineMessageID,
			Data:            update.CallbackQuery.Data,
		},
	)
	if err != nil {
		e.logger.ErrorKV(ctx, "Error convert update", "err", err)
		return nil, false
	}

	return result, true
}

func fromToSender(from *tgbotapi.User) entities.Sender {
	if from == nil {
		return entities.Sender{}
	}

	return entities.Sender{
		ID:           from.ID,
		FirstName:    from.FirstName,
		LastName:     from.LastName,
		UserName:     from.UserName,
		LanguageCode: from.LanguageCode,
	}
}

func convertMessage(message *tgbotapi.Message) entities.Message {
	if message == nil {
		return entities.Message{}
	}

	return entities.Message{
		ID:   int64(message.MessageID),
		From: fromToSender(message.From),
		Chat: convertChat(message.Chat),
		Text: message.Text,
	}
}

func convertChat(chat *tgbotapi.Chat) entities.Chat {
	if chat == nil {
		return entities.Chat{}
	}

	return entities.Chat{
		ID: chat.ID,
	}
}
