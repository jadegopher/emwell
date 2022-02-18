package handlers

import (
	"errors"

	"emwell/internal/core/link"
	"emwell/internal/logger"
)

var (
	ErrNilRequest = errors.New("nil request")
)

type Handlers struct {
	logger      logger.ILogger
	linkService *link.Service
}

func NewHandlers(logger logger.ILogger, linkService *link.Service) *Handlers {
	return &Handlers{
		logger:      logger,
		linkService: linkService,
	}
}
