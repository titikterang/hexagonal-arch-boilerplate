package services

import (
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/core/ports"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/config"
)

type PaymentService struct {
	config     *config.Config
	repository ports.WalletRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.WalletRepositoryAdapter) ports.PaymentServiceAdapter {
	return &PaymentService{
		config:     cfg,
		repository: repository,
	}
}
