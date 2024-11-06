package main

import (
	"context"
	zerolog "github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	"github.com/go-kratos/kratos/v2/log"
	zlog "github.com/rs/zerolog"
	"os"
)

func main() {
	ctx := context.Background()

	zlogT := zlog.New(os.Stdout).With().CallerWithSkipFrameCount(4).Timestamp().Logger()
	logger := zerolog.NewLogger(&zlogT)
	log.SetLogger(logger)

	log.Info("Starting walletService")

	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("failed to parse config, %#v", err)
	}

	startLoanService(ctx, cfg)
}
