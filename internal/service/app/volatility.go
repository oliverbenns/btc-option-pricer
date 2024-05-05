package app

import (
	"fmt"
	"strconv"

	"github.com/oliverbenns/btc-option-pricer/internal/bybit"
)

func (s *Service) getHistoricalVolatility(baseCoin string) (float64, error) {
	if volatility, ok := s.volatilityCache[baseCoin]; ok {
		return volatility, nil
	}

	res, err := s.bybitClient.GetHistoricalVolatility(&bybit.GetHistoricalVolatilityProps{
		Category: "option",
		BaseCoin: baseCoin,
		Period:   90,
	})
	if err != nil {
		return 0, err
	}

	if len(res.Result) == 0 {
		return 0, fmt.Errorf("no historical volatility data found")
	}

	volatility, err := strconv.ParseFloat(res.Result[0].Value, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse volatility: %w", err)
	}

	s.volatilityCache[baseCoin] = volatility

	return volatility, nil

}
