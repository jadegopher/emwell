package user

import (
	"context"
	"time"

	"emwell/internal/user/entites"
)

func (m *Manager) CreateIfNotExists(ctx context.Context, user entites.User) (result entites.User, err error) {
	if err = user.Validate(); err != nil {
		m.logger.ErrorKV(ctx, "Error validate user", "err", err)
		return entites.User{}, err
	}

	id, err := m.userRepository.Insert(ctx, user)
	if err != nil {
		m.logger.ErrorKV(ctx, "Error insert user", "err", err)
		return entites.User{}, err
	}

	return entites.NewUserEntity(id, time.Time{}, user), nil
}
