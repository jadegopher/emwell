package user

import (
	"emwell/internal/logger"
	"emwell/internal/user/repository"
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
