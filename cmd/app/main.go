package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/oliverbenns/btc-option-pricer/internal/bybit"
	"github.com/oliverbenns/btc-option-pricer/internal/service/app"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	bybitClient := bybit.NewClient("https://api.bybit.com")

	riskFreeRate := flag.Float64("riskFreeRate", 5.0, "risk free rate")
	flag.Parse()

	svc := app.NewService(logger, bybitClient, *riskFreeRate)

	err := svc.Run()
	if err != nil {
		panic(err)
	}
}
