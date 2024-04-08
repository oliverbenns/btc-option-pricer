package app

import (
	"fmt"
	"log/slog"

	"github.com/oliverbenns/btc-option-pricer/internal/bybit"
)

type Service struct {
	logger      *slog.Logger
	bybitClient *bybit.Client
}

func NewService(logger *slog.Logger, bybitClient *bybit.Client) *Service {
	return &Service{
		logger:      logger,
		bybitClient: bybitClient,
	}
}

func (s *Service) Run() error {
	s.logger.Info("Service running")

	res, err := s.bybitClient.GetTickers(&bybit.GetTickersProps{
		Category: "spot",
		Symbol:   "BTCUSDT",
	})
	if err != nil {
		return fmt.Errorf("failed to get tickers: %w", err)
	}

	s.logger.Info("Got tickers", "tickers", res)

	return nil
}
