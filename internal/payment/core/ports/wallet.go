package ports

import (
	"context"
)

// WalletRepositoryAdapter - port secondary
type WalletRepositoryAdapter interface {
	ReadBalanceInfoFromWallet(ctx context.Context, userID string) (float64, error)
	AppendBalanceInfoIntoWallet(ctx context.Context, userID string, amount float64) error
}
