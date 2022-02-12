package entites

import (
	"errors"
	"time"
)

var (
	ErrWrongTelegramID = errors.New("telegramID should be greater than 0")
)

type User struct {
	id         int64
	Name       string
	Language   LanguageType
	TelegramID int64
	createdAt  time.Time
}

type LanguageType = string

const (
	RU LanguageType = "ru"
)

func NewUserEntity(id int64, createdAt time.Time, entity User) User {
	return User{
		id:         id,
		Name:       entity.Name,
		Language:   entity.Language,
		TelegramID: entity.TelegramID,
		createdAt:  createdAt,
	}
}

func (e *User) ID() int64 {
	return e.id
}

func (e *User) CreatedAt() time.Time {
	return e.createdAt
}

func (e *User) Validate() error {
	if e.TelegramID <= 0 {
		return ErrWrongTelegramID
	}

	if e.Language == "" {
		e.Language = RU
	}

	if e.Name == "" {
		e.Name = "Robert Paulson"
	}

	return nil
}
