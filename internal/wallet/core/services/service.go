package services

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
)

// GetUserBalance - primary adapter
func (s *WalletService) GetUserBalance(ctx context.Context, userID string) (models.BalanceResponse, error) {

	return models.BalanceResponse{}, nil
}

// UpdateUserBalance - primary adapter
func (s *WalletService) UpdateUserBalance(ctx context.Context, payload models.UpdateBalancePayload) error {

	return nil
}
