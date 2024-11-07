package ports

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
)

// WalletServiceAdapter - port primary
type WalletServiceAdapter interface {
	GetUserBalance(ctx context.Context, userID string) (models.BalanceResponse, error)
	UpdateUserBalance(ctx context.Context, payload models.UpdateBalancePayload) error
}

// WalletRepositoryAdapter - port secondary
type WalletRepositoryAdapter interface {
	ReadBalanceInfoFromDatastore(ctx context.Context, userID string) (models.DatastoreBalanceResponse, error)
	AppendBalanceInfoIntoDatastore(ctx context.Context, userID string, amount float64) error
}
