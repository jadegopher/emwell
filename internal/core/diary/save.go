package diary

import (
	"context"
	"time"

	"emwell/internal/core/diary/entites"
)

func (d *Diary) SaveEmotionalInformation(
	ctx context.Context,
	diary entites.EmotionalDiary,
) (entites.EmotionalDiary, error) {
	if err := diary.Validate(); err != nil {
		d.logger.ErrorKV(ctx, "Error validate diary", "err", err)
		return entites.EmotionalDiary{}, err
	}

	id, err := d.emotionalRepo.Insert(ctx, diary)
	if err != nil {
		d.logger.ErrorKV(ctx, "Error insert diary", "err", err)
		return entites.EmotionalDiary{}, err
	}

	return entites.NewEmotionalDiaryEntity(id, time.Time{}, diary), nil
}
