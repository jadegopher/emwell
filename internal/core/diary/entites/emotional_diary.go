package entites

import (
	"errors"
	"time"
)

var (
	ErrInvalidUserID        = errors.New("invalid userID")
	ErrInvalidEmotionalRate = errors.New("invalid emotional rate")
)

const (
	MinEmotionalRate     = -1000
	WorstEmotionalRate   = -750
	WorseEmotionalRate   = -500
	BadEmotionalRate     = -250
	NeutralEmotionalRate = 0
	GoodEmotionalRate    = 250
	BetterEmotionalRate  = 500
	BestEmotionalRate    = 750
	MaxEmotionalRate     = 1000
)

type EmotionalDiary struct {
	id            int64
	UserID        int64
	EmotionalRate int32
	ReferToDate   time.Time
	createdAt     time.Time
}

func NewEmotionalDiaryEntity(id int64, createdAt time.Time, entity EmotionalDiary) EmotionalDiary {
	return EmotionalDiary{
		id:            id,
		UserID:        entity.UserID,
		EmotionalRate: entity.EmotionalRate,
		ReferToDate:   entity.ReferToDate,
		createdAt:     createdAt,
	}
}

func (e *EmotionalDiary) ID() int64 {
	return e.id
}

func (e *EmotionalDiary) CreatedAt() time.Time {
	return e.createdAt
}

func (e *EmotionalDiary) Validate() error {
	if e.UserID <= 0 {
		return ErrInvalidUserID
	}

	if e.EmotionalRate < MinEmotionalRate || e.EmotionalRate > MaxEmotionalRate {
		return ErrInvalidEmotionalRate
	}

	return nil
}
