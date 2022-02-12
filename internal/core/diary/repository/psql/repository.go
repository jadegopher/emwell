package psql

import (
	"context"
	"database/sql"
	"errors"

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

func (r *Repository) Insert(ctx context.Context, diary entites.EmotionalDiary) (id int64, err error) {
	result, err := r.db.QueryContext(ctx, InsertQuery,
		diary.UserID,
		diary.EmotionalRate,
	)
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
