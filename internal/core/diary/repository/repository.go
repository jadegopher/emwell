package repository

import (
	"context"

	"emwell/internal/core/diary/entites"
)

type EmotionalDiaryRepository interface {
	Insert(ctx context.Context, diary entites.EmotionalDiary) (id int64, err error)
}
