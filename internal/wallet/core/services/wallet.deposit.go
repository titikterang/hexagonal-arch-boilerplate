package services

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
)

// primary adapter
func (s *WalletService) UpdateUserBalance(ctx context.Context, payload models.UpdateBalancePayload) error {
	return nil
}
