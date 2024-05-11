package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/oliverbenns/btc-option-pricer/internal/bybit"
)

type Ticker struct {
	Symbol          string
	ExpiryDate      time.Time
	StrikePrice     float64
	UnderlyingPrice float64
	Kind            string // C(all) / P(ut)
	BestBidPrice    float64
	BestAskPrice    float64
	BaseCoin        string
}

func (s *Service) getTickers() ([]Ticker, error) {
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

	parsed := make([]Ticker, len(filtered))
	for i, ticker := range filtered {
		parsedTicker, err := parseTicker(ticker)
		if err != nil {
			return nil, fmt.Errorf("could not parse ticker: %w", err)
		}

		parsed[i] = *parsedTicker
	}

	return parsed, nil
}

func filterByPeriod(tickers []bybit.GetTickersResultResultTicker, from, to time.Time) ([]bybit.GetTickersResultResultTicker, error) {
	filtered := []bybit.GetTickersResultResultTicker{}
	for _, ticker := range tickers {
		details, err := parseSymbol(ticker.Symbol)
		if err != nil {
			return nil, fmt.Errorf("failed to parse symbol: %w", err)
		}

		isWithinRange := (details.expiryDate.Equal(from) || details.expiryDate.After(from)) && details.expiryDate.Before(to)
		if isWithinRange {
			filtered = append(filtered, ticker)
		}
	}

	return filtered, nil
}

func parseTicker(ticker bybit.GetTickersResultResultTicker) (*Ticker, error) {
	details, err := parseSymbol(ticker.Symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to parse symbol: %w", err)
	}

	underlyingPrice, err := strconv.ParseFloat(*ticker.UnderlyingPrice, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse underlying price: %w", err)
	}

	bestAsk, err := strconv.ParseFloat(ticker.Ask1Price, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ask price: %w", err)
	}

	bestBid, err := strconv.ParseFloat(ticker.Bid1Price, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse bid price: %w", err)
	}

	return &Ticker{
		Symbol:          ticker.Symbol,
		ExpiryDate:      details.expiryDate,
		StrikePrice:     details.strikePrice,
		UnderlyingPrice: underlyingPrice,
		Kind:            details.kind,
		BestAskPrice:    bestAsk,
		BestBidPrice:    bestBid,
		BaseCoin:        details.baseCoin,
	}, nil
}

type symbolDetails struct {
	baseCoin    string
	expiryDate  time.Time
	strikePrice float64
	kind        string // C(all) / P(ut)
}

// BTC-28APR24-67000-P
func parseSymbol(symbol string) (*symbolDetails, error) {
	symbolParts := strings.Split(symbol, "-")
	if len(symbolParts) < 4 {
		return nil, errors.New("invalid symbol")
	}

	rawDate := symbolParts[1]

	expiryDate, err := time.Parse("2Jan06", rawDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %w", err)
	}

	strikePrice, err := strconv.ParseFloat(symbolParts[2], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse strike price: %w", err)
	}

	return &symbolDetails{
		baseCoin:    symbolParts[0],
		expiryDate:  expiryDate,
		strikePrice: strikePrice,
		kind:        symbolParts[3],
	}, nil
}
