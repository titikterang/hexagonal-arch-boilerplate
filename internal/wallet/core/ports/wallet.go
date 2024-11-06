package ports

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
)

// port primary
type WalletService interface {
	GetUserBalance(ctx context.Context, userID string) (models.BalanceResponse, error)
	UpdateUserBalance(ctx context.Context, payload models.UpdateBalancePayload) error
}

// port secondary
type WalletRepository interface {
	ReadBalanceInfoFromDatastore(ctx context.Context, userID string) (models.DatastoreBalanceResponse, error)
	AppendBalanceInfoIntoDatastore(ctx context.Context, userID string, amount float64) error
}
