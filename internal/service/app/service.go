package app

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
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

	tickers, err := s.getTickers()
	if err != nil {
		return fmt.Errorf("failed to get tickers: %w", err)
	}

	//s.logger.Info("Got tickers", "tickers", tickers)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Symbol"})

	for _, ticker := range tickers {
		row := []string{ticker.Symbol}
		table.Append(row)
	}

	table.Render()

	return nil
}

func (s *Service) getTickers() ([]bybit.GetTickersResultResultTicker, error) {
	res, err := s.bybitClient.GetTickers(&bybit.GetTickersProps{
		Category: "option",
		BaseCoin: "BTC",
	})
	if err != nil {
		return nil, err
	}

	from := time.Now()
	to := from.AddDate(0, 0, 30)

	filtered, err := filterByPeriod(res.Result.List, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to filter by period: %w", err)
	}

	return filtered, nil
}

func filterByPeriod(tickers []bybit.GetTickersResultResultTicker, from, to time.Time) ([]bybit.GetTickersResultResultTicker, error) {
	filtered := []bybit.GetTickersResultResultTicker{}
	for _, ticker := range tickers {
		symbolParts := strings.Split(ticker.Symbol, "-")
		if len(symbolParts) < 2 {
			return nil, errors.New("invalid symbol")
		}

		rawDate := symbolParts[1]

		date, err := time.Parse("2Jan06", rawDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse symbol date: %w", err)
		}

		if (date.Equal(from) || date.After(from)) && date.Before(to) {
			filtered = append(filtered, ticker)
		}
	}

	return filtered, nil
}
