package diary

import (
	"context"
	"errors"
	"time"

	"emwell/internal/core/diary/entites"
)

func validateInformation(userID int64, from, to time.Time) error {
	if userID <= 0 {
		return errors.New("userID incorrect")
	}

	if !to.After(from) {
		return errors.New("to should be greater than from")
	}

	return nil
}

func (d *Diary) GetStatistics(ctx context.Context, userID int64, from, to time.Time) (entites.EmotionalInfos, error) {
	if err := validateInformation(userID, from, to); err != nil {
		d.logger.ErrorKV(ctx, "error validate", "err", err)
		return nil, err
	}

	info, err := d.emotionalRepo.SelectByUserID(ctx, userID, from, to)
	if err != nil {
		d.logger.ErrorKV(ctx, "error select statistics by user id", "err", err)
		return nil, err
	}

	return d.converter.ConvertToPoints(info), nil
}
