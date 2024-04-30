package app

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/oliverbenns/btc-option-pricer/internal/blackscholes"
	"github.com/oliverbenns/btc-option-pricer/internal/bybit"
	"github.com/oliverbenns/btc-option-pricer/internal/utils"
)

type Service struct {
	logger       *slog.Logger
	bybitClient  *bybit.Client
	riskFreeRate float64
}

func NewService(logger *slog.Logger, bybitClient *bybit.Client, riskFreeRate float64) *Service {
	return &Service{
		logger:       logger,
		bybitClient:  bybitClient,
		riskFreeRate: riskFreeRate,
	}
}

func (s *Service) Run() error {
	s.logger.Info("Service running")

	tickers, err := s.getTickers()
	if err != nil {
		return fmt.Errorf("failed to get tickers: %w", err)
	}

	s.logger.Info("Got tickers", "tickers", tickers)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Symbol", "Best Ask", "Best Bid", "Black Scholes"})

	for _, ticker := range tickers {

		now := time.Now()
		diff := ticker.ExpiryDate.Sub(now)

		value := blackscholes.Calculate(&blackscholes.CalculateProps{
			StrikePrice:     ticker.StrikePrice,
			UnderlyingPrice: ticker.UnderlyingPrice,
			TimeToExp:       utils.GetYearsFromDuration(diff),
			RiskFreeRate:    s.riskFreeRate,
			Volatility:      0, //todo
		})
		row := []string{
			ticker.Symbol,
			fmt.Sprintf("%.2f", ticker.BestAskPrice),
			fmt.Sprintf("%.2f", ticker.BestBidPrice),
			fmt.Sprintf("%.2f", value),
		}
		table.Append(row)
	}

	table.Render()

	return nil
}
