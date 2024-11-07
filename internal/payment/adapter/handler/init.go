package handler

import (
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/core/ports"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/config"
	pbHandler "github.com/titikterang/hexagonal-arch-boilerplate/lib/protos/v1/payment"
)

type Handler struct {
	pbHandler.UnimplementedPaymentServiceServer
	config         *config.Config
	paymentService ports.PaymentServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.PaymentServiceAdapter) (*Handler, error) {
	return &Handler{
		config:         cfg,
		paymentService: adapter,
	}, nil
}
