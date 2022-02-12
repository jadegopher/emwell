package user

import (
	"emwell/internal/core/user/repository"
	"emwell/internal/logger"
)

type Manager struct {
	logger         logger.ILogger
	userRepository repository.UserRepository
}

func NewManager(logger logger.ILogger, userRepository repository.UserRepository) *Manager {
	return &Manager{
		logger:         logger,
		userRepository: userRepository,
	}
}
