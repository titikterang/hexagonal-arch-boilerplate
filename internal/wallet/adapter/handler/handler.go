package handler

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/protos/v1/wallet"
)

func (h *Handler) GetUserBalance(ctx context.Context, in *wallet.GetBalanceRequest) (*wallet.GetBalanceResponse, error) {
	resp, err := h.walletService.GetUserBalance(ctx, in.GetUserId())
	if err != nil {
		/// log err
		return nil, err
	}
	data := &wallet.GetBalanceResponse{
		UserId:  in.UserId,
		Balance: resp.AvailableBalance,
	}

	return data, nil
}

func (h *Handler) UpdateUserBalance(ctx context.Context, in *wallet.UpdateBalanceRequest) (*wallet.UpdateBalanceResponse, error) {
	var message = "success"
	amount, err := h.walletService.UpdateUserBalance(ctx, models.UpdateBalancePayload{
		UserID: in.GetUserId(),
		Amount: in.GetAmount(),
	})
	if err != nil {
		// log err
		message = err.Error()
	}

	return &wallet.UpdateBalanceResponse{
		Message:      message,
		Success:      err == nil,
		FinalBalance: amount,
	}, nil
}
