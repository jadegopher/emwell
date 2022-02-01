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
	Send(ctx context.Context, toChatID int64, response handlers.Response) error
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

func (r *ReplySender) Send(ctx context.Context, toChatID int64, response handlers.Response) error {
	_, err := r.bot.Send(responseToChattable(toChatID, response))
	if err != nil {
		r.logger.ErrorKV(ctx, "Send reply error", "chat_id", toChatID, "err", err)
		return errors.Wrap(ErrSendReply, err.Error())
	}

	return nil
}

func responseToChattable(chatID int64, response handlers.Response) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, response.Text)

	if len(response.Buttons) != 0 {
		msg.ReplyMarkup = convertKeyBoard(response.Buttons)
	}

	return msg
}

func convertKeyBoard(buttons [][]handlers.Button) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, len(buttons))
	for i, row := range buttons {
		rows[i] = make([]tgbotapi.InlineKeyboardButton, len(row))
		for j, button := range row {
			rows[i][j] = tgbotapi.NewInlineKeyboardButtonData(button.Text, button.Data)
		}
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
