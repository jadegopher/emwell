package telegram

import (
	"context"

	"emwell/internal/api/telegram/consumer"
	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/api/telegram/handlers"
	"emwell/internal/api/telegram/middlewares"
	"emwell/internal/api/telegram/sender"
	"emwell/internal/logger"
)

type Telegram struct {
	logger      logger.ILogger
	consumer    consumer.Consumer
	middlewares []middlewares.Middleware
	handlers    []handlers.Handler
	sender      sender.Sender
}

func NewTelegramBotAPI(
	logger logger.ILogger,
	consumer consumer.Consumer,
	middlewares []middlewares.Middleware,
	handlers []handlers.Handler,
	sender sender.Sender,
) (*Telegram, error) {
	return &Telegram{
		logger:      logger,
		consumer:    consumer,
		middlewares: middlewares,
		handlers:    handlers,
		sender:      sender,
	}, nil
}

func (t *Telegram) HandleUpdates(ctx context.Context) (err error) {
	upds := t.consumer.StartConsuming(ctx)

	for upd := range upds {
		_ = t.process(ctx, upd)
	}

	return nil
}

func (t *Telegram) process(ctx context.Context, upd entities.Update) (err error) {
	for _, middleware := range t.middlewares {
		upd, err = middleware.Serve(ctx, upd)
		if err != nil {
			t.logger.ErrorKV(ctx, "middleware serve error", "err", err)
			return err
		}
	}

	for _, handler := range t.handlers {
		if !handler.CanHandle(upd) {
			continue
		}

		resp, err := handler.Handle(ctx, upd)
		if err != nil {
			t.logger.ErrorKV(ctx, "handle error", "err", err)
			return err
		}

		if resp == nil {
			continue
		}

		for _, r := range resp {
			if err = t.sender.Send(ctx, r); err != nil {
				t.logger.ErrorKV(ctx, "send error", "err", err)
				return err
			}
		}

		break
	}

	return nil
}
