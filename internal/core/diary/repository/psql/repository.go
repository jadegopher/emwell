package psql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"emwell/internal/core/diary/entites"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Insert(ctx context.Context, diary entites.EmotionalInfo) (id int64, err error) {
	result, err := r.db.QueryContext(ctx, InsertQuery, diary.UserID, diary.EmotionalRate)
	if err != nil {
		return 0, err
	}

	if !result.Next() {
		return 0, errors.New("no id in result set")
	}

	if err = result.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) SelectByUserID(
	ctx context.Context,
	userID int64,
	from, to time.Time,
) (infos []entites.EmotionalInfo, err error) {
	result, err := r.db.QueryContext(ctx, SelectByUserIDQuery, userID, from, to)
	if err != nil {
		return nil, err
	}

	infos = make([]entites.EmotionalInfo, 0)
	for result.Next() {
		id := int64(0)
		tmp := entites.EmotionalInfo{}
		createdAt := time.Time{}
		if err = result.Scan(
			&id,
			&tmp.UserID,
			&tmp.EmotionalRate,
			&tmp.ReferToDate,
			&createdAt,
		); err != nil {
			return nil, err
		}
		infos = append(infos, entites.NewEmotionalDiaryEntity(id, createdAt, tmp))
	}

	return infos, nil
}
