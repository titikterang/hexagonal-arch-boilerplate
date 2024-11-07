package repository

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
)

func (s *WalletRepository) ReadBalanceInfoFromDatastore(ctx context.Context, userID string) (models.DatastoreBalanceResponse, error) {
	return models.DatastoreBalanceResponse{}, nil
}

func (s *WalletRepository) AppendBalanceInfoIntoDatastore(ctx context.Context, userID string, amount float64) error {
	return nil
}
