package main

import (
	"log/slog"
	"os"

	"github.com/oliverbenns/btc-option-pricer/internal/bybit"
	"github.com/oliverbenns/btc-option-pricer/internal/service/app"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	bybitClient := bybit.NewClient("https://api.bybit.com")

	svc := app.NewService(logger, bybitClient)

	err := svc.Run()
	if err != nil {
		panic(err)
	}
}
