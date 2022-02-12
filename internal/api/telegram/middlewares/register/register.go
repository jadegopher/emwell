package register

import (
	"context"

	"emwell/internal/api/telegram/consumer/entities"
	"emwell/internal/core/user"
	"emwell/internal/core/user/entites"
	"emwell/internal/logger"
)

type Middleware struct {
	logger      logger.ILogger
	userManager *user.Manager
}

func NewMiddleware(logger logger.ILogger, userManager *user.Manager) *Middleware {
	return &Middleware{
		logger:      logger,
		userManager: userManager,
	}
}

func (m *Middleware) Serve(ctx context.Context, upd entities.Update) (entities.Update, error) {
	senderInfo := upd.Sender()

	userInfo, err := m.userManager.CreateIfNotExists(ctx, entites.User{
		Name:       senderInfo.UserName,
		Language:   senderInfo.LanguageCode,
		TelegramID: senderInfo.ID,
	})
	if err != nil {
		return nil, err
	}

	newUpd, err := entities.NewUpdate(
		upd.ID(),
		upd.Type(),
		senderInfo,
		&userInfo,
		upd.Payload(),
	)
	if err != nil {
		m.logger.ErrorKV(ctx, "Error create new update", "err", err)
		return nil, err
	}

	return newUpd, nil
}
