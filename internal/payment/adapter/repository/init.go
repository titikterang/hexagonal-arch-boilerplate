package repository

import (
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/core/ports"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/config"
	"net/http"
	"time"
)

type PaymentRepository struct {
	Client *http.Client
	cfg    *config.Config
}

func NewRepository(cfg *config.Config) ports.WalletRepositoryAdapter {
	return &PaymentRepository{
		cfg: cfg,
		Client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     10 * time.Second,
				MaxIdleConns:        100,
			},
		},
	}
}
