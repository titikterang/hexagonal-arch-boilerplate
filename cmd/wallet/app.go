package main

import (
	"github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	"github.com/go-kratos/kratos/v2/log"
	zlog "github.com/rs/zerolog"
	"os"
)

func main() {
	zlogT := zlog.New(os.Stdout).With().CallerWithSkipFrameCount(4).Timestamp().Logger()
	logger := zerolog.NewLogger(&zlogT)
	log.SetLogger(logger)

	log.Info("Starting backend service")

	cfg, err := initConfig("wallet.config.json")
	if err != nil {
		log.Fatalf("failed to parse config, %#v", err)
	}

	startService(cfg)
}
