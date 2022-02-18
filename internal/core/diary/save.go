package diary

import (
	"context"
	"time"

	"emwell/internal/core/diary/entites"
)

func (d *Diary) SaveEmotionalInformation(
	ctx context.Context,
	diary entites.EmotionalInfo,
) (entites.EmotionalInfo, error) {
	if err := diary.Validate(); err != nil {
		d.logger.ErrorKV(ctx, "Error validate diary", "err", err)
		return entites.EmotionalInfo{}, err
	}

	id, err := d.emotionalRepo.Insert(ctx, diary)
	if err != nil {
		d.logger.ErrorKV(ctx, "Error insert diary", "err", err)
		return entites.EmotionalInfo{}, err
	}

	return entites.NewEmotionalDiaryEntity(id, time.Time{}, diary), nil
}
