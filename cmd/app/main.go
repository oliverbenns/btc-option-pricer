package main

import (
	"log/slog"
	"os"

	"github.com/oliverbenns/btc-option-pricer/internal/service/app"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	svc := app.NewService(logger)

	err := svc.Run()
	if err != nil {
		panic(err)
	}
}
