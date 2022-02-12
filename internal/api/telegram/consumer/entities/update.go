package entities

import (
	"fmt"

	"emwell/internal/core/user/entites"
)

type Update interface {
	ID() int64
	Type() UpdateType
	Sender() Sender
	User() (entites.User, bool)
	Payload() interface{}
}

var (
	ErrWrongPayload = fmt.Errorf("wrong payload type")
	ErrUnknownType  = fmt.Errorf("unknown type")
)

type UpdateEntity struct {
	id      int64
	typ     UpdateType
	sender  Sender
	user    *entites.User
	payload interface{}
}

func NewUpdate(
	id int64,
	typ UpdateType,
	sender Sender,
	user *entites.User,
	payload interface{},
) (*UpdateEntity, error) {
	switch typ {
	case UpdateTypeMessage:
		_, ok := payload.(Message)
		if !ok {
			return nil, ErrWrongPayload
		}
	case UpdateTypeCallback:
		_, ok := payload.(Callback)
		if !ok {
			return nil, ErrWrongPayload
		}
	default:
		return nil, ErrUnknownType
	}

	return &UpdateEntity{
		id:      id,
		typ:     typ,
		sender:  sender,
		user:    user,
		payload: payload,
	}, nil
}

func (m *UpdateEntity) ID() int64 {
	return m.id
}

func (m *UpdateEntity) Type() UpdateType {
	return m.typ
}

func (m *UpdateEntity) Sender() Sender {
	return m.sender
}

func (m *UpdateEntity) User() (entites.User, bool) {
	if m.user == nil {
		return entites.User{}, false
	}

	return *m.user, true
}

func (m *UpdateEntity) Payload() interface{} {
	return m.payload
}
