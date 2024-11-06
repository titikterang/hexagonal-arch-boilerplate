package ports

import "context"

type PaymentService interface {
	Mutation(ctx context.Context) error
	Transfer(ctx context.Context) error
}
