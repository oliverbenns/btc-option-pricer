package app

import (
	"log/slog"
)

type Service struct {
	logger *slog.Logger
}

func NewService(logger *slog.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) Run() error {
	s.logger.Info("Service running")

	return nil
}
