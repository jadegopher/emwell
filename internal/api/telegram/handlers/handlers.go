package handlers

import (
	"context"
	"errors"

	"emwell/internal/api/telegram/consumer/entities"
)

var (
	ErrCantHandle       = errors.New("can't handle update")
	ErrUserNotSpecified = errors.New("user not specified")
)

type Handler interface {
	CanHandle(update entities.Update) bool
	Handle(ctx context.Context, update entities.Update) (resp []Response, err error)
}

type ResponseType int8

const (
	ResponseTypeUnknown = iota
	ResponseTypeMessage
	ResponseTypeCallback
	ResponseTypeKeyboard
)

type Response interface {
	Type() ResponseType
	Payload() interface{}
}
