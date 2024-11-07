package services

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
	"strconv"
)

// GetUserBalance - primary adapter
func (s *WalletService) GetUserBalance(ctx context.Context, userID string) (models.BalanceResponse, error) {
	resp, err := s.repository.ReadBalanceInfoFromDatastore(ctx, userID)
	if err != nil {
		return models.BalanceResponse{}, err
	}

	var amountTotal float64
	for _, v := range resp.List {
		val, _ := strconv.ParseFloat(v, 64)
		amountTotal += val
	}

	return models.BalanceResponse{
		UserID:           userID,
		AvailableBalance: amountTotal,
	}, nil
}

// UpdateUserBalance - primary adapter
func (s *WalletService) UpdateUserBalance(ctx context.Context, payload models.UpdateBalancePayload) (float64, error) {
	// append balance
	err := s.repository.AppendBalanceInfoIntoDatastore(ctx, payload.UserID, payload.Amount)
	if err != nil {
		log.Error("failed to update balance")
	}
	// fetch latest balance
	resp, err := s.GetUserBalance(ctx, payload.UserID)
	if err != nil {
		log.Error("failed to get latest balance")
	}

	return resp.AvailableBalance, nil
}
