package ports

import "context"

type PaymentService interface {
	Deposit(ctx context.Context) error
	Transfer(ctx context.Context) error
}

type PaymentRepository interface {
	DepositToWallet(ctx context.Context) error
	TransferToAccount(ctx context.Context) error
}
