package link

import (
	"time"

	"emwell/internal/core/link/repository"
	"emwell/internal/logger"
)

type Service struct {
	logger     logger.ILogger
	timeGetter TimeGetter
	repo       repository.LinkRepository
	secret     string
}

type TimeGetter interface {
	Now() time.Time
}

func NewLinkService(logger logger.ILogger, secret string, getter TimeGetter, repo repository.LinkRepository) *Service {
	return &Service{
		logger:     logger,
		secret:     secret,
		timeGetter: getter,
		repo:       repo,
	}
}
