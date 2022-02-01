package psql

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"emwell/internal/user/entites"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Insert(ctx context.Context, user entites.User) (id int64, err error) {
	result, err := r.db.QueryContext(ctx, InsertQuery,
		user.Name,
		user.Language,
		user.TelegramID,
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
