package services

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
)

// primary adapter
func (s *WalletService) GetUserBalance(ctx context.Context, userID string) (models.BalanceResponse, error) {
	return models.BalanceResponse{}, nil
}

// secondary adapter
func (s *WalletService) ReadBalanceInfoFromDatastore(ctx context.Context, userID string) (models.DatastoreBalanceResponse, error) {
	return models.DatastoreBalanceResponse{}, nil
}

func (s *WalletService) AppendBalanceInfoIntoDatastore(ctx context.Context, userID string, amount float64) error {
	return nil
}
