package main

import (
	"context"
	kratos "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/http"
	gorHdl "github.com/gorilla/handlers"
	"github.com/titikterang/amartha-loan-svc/internal/model"
	hdl "github.com/titikterang/amartha-loan-svc/internal/transport/handler"
	pbLoan "github.com/titikterang/amartha-loan-svc/protos/loan"
)

func startLoanService(ctx context.Context, cfg *model.Config) {
	uCase, err := InitUseCase(cfg)
	if err != nil {
		log.Fatalf("failed run InitUseCase: %v", err)
	}

	handler, err := hdl.NewHandler(cfg, uCase)
	if err != nil {
		log.Fatalf("failed initiate NewHandler: %v", err)
	}

	httpOpts := []http.ServerOption{
		http.Timeout(cfg.App.Timeout),
		http.Address(cfg.App.Address),
		http.Filter(
			gorHdl.CORS(
				gorHdl.AllowedOrigins([]string{"*"}),
				gorHdl.AllowedHeaders([]string{"Content-Type", "Authorization"}),
				gorHdl.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
			),
		),
		http.Middleware(
			metadata.Server(),
			logging.Server(log.GetLogger()),
		),
		http.Logger(log.GetLogger()),
	}

	httpServer := http.NewServer(
		httpOpts...,
	)

	pbLoan.RegisterLoanServiceHTTPServer(httpServer, handler)

	server := kratos.New(
		kratos.Name(cfg.App.Name),
		kratos.Server(
			httpServer,
		),
	)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
