package handlers

import (
	"context"
	"errors"

	"emwell/internal/api/telegram/consumer/entities"
)

var (
	ErrCantHandle = errors.New("can't handle update")
)

type Handler interface {
	CanHandle(update entities.Update) bool
	Handle(ctx context.Context, update entities.Update) (resp []Response, err error)
}

type Response struct {
	Text    string
	Buttons [][]Button
}

type Button struct {
	Text string
	Data string
}
