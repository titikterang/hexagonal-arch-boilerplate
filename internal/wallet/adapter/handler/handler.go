package handler

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/wallet/core/models"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/protos/v1/old"
)

func (h *Handler) GetUserBalance(ctx context.Context, in *old.GetBalanceRequest) (*old.GetBalanceResponse, error) {
	resp, err := h.walletService.GetUserBalance(ctx, in.GetUserId())
	if err != nil {
		/// log err
		return nil, err
	}
	data := &old.GetBalanceResponse{
		UserId:  in.UserId,
		Balance: resp.AvailableBalance,
	}

	return data, nil
}

func (h *Handler) UpdateUserBalance(ctx context.Context, in *old.UpdateBalanceRequest) (*old.UpdateBalanceResponse, error) {
	var message = "success"
	amount, err := h.walletService.UpdateUserBalance(ctx, models.UpdateBalancePayload{
		UserID: in.GetUserId(),
		Amount: in.GetAmount(),
	})
	if err != nil {
		// log err
		message = err.Error()
	}

	return &old.UpdateBalanceResponse{
		Message:      message,
		Success:      err == nil,
		FinalBalance: amount,
	}, nil
}
