package services

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/core/models"
)

func (s *PaymentService) TransferUserBalance(ctx context.Context, payload models.TransferBalancePayload) (float64, error) {
	// deduct source user balance
	s.repository.AppendBalanceInfoIntoWallet(ctx, payload.SourceUserID, payload.Amount*-1)
	// add destination user balance
	s.repository.AppendBalanceInfoIntoWallet(ctx, payload.TargetUserID, payload.Amount)
	// get final destination user balance
	resp, _ := s.repository.ReadBalanceInfoFromWallet(ctx, payload.TargetUserID)

	return resp, nil
}
