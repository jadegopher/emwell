package repository

import (
	"context"
	"time"

	"emwell/internal/core/diary/entites"
)

type EmotionalDiaryRepository interface {
	Insert(ctx context.Context, diary entites.EmotionalInfo) (id int64, err error)
	SelectByUserID(ctx context.Context, userID int64, from, to time.Time) (infos []entites.EmotionalInfo, err error)
}
