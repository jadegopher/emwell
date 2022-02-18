package sender

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"emwell/internal/api/telegram/handlers"
	"emwell/internal/logger"
)

type Sender interface {
	Send(ctx context.Context, response handlers.Response) error
}

var (
	ErrSendReply = fmt.Errorf("cannot send reply: ")
)

type ReplySender struct {
	logger logger.ILogger
	bot    *tgbotapi.BotAPI
}

func NewReplySender(logger logger.ILogger, token string) (*ReplySender, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &ReplySender{
		logger: logger,
		bot:    bot,
	}, nil
}

func (r *ReplySender) Send(ctx context.Context, response handlers.Response) error {
	chattableResponse, err := responseToChattable(response)
	if err != nil {
		return err
	}
	_, err = r.bot.Request(chattableResponse)
	if err != nil {
		r.logger.ErrorKV(ctx, "Send reply error", "err", err)
		return errors.Wrap(ErrSendReply, err.Error())
	}

	return nil
}

func responseToChattable(response handlers.Response) (tgbotapi.Chattable, error) {
	switch response.Type() {
	case handlers.ResponseTypeMessage:
		payload, ok := response.Payload().(handlers.MessagePayload)
		if !ok {
			return nil, errors.New("message payload incorrect type")
		}
		msg := tgbotapi.NewMessage(payload.ToChatID, payload.Text)

		if len(payload.InlineKeyboard) != 0 {
			msg.ReplyMarkup = convertInlineKeyboard(payload.InlineKeyboard)
		}

		return msg, nil
	case handlers.ResponseTypeCallback:
		payload, ok := response.Payload().(handlers.CallbackPayload)
		if !ok {
			return nil, errors.New("message payload incorrect type")
		}
		callback := tgbotapi.NewCallback(payload.CallbackID, payload.Text)

		return callback, nil
	default:
		return nil, errors.New("unknown type")
	}
}

func convertInlineKeyboard(buttons [][]handlers.Button) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, len(buttons))
	for i, row := range buttons {
		rows[i] = make([]tgbotapi.InlineKeyboardButton, len(row))
		for j, button := range row {
			rows[i][j] = tgbotapi.NewInlineKeyboardButtonData(button.Text, button.Data)
		}
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
