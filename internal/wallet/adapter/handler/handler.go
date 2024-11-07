package handler

import (
	"context"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/protos/v1/wallet"
)

func (h *Handler) GetUserBalance(ctx context.Context, in *wallet.GetBalanceRequest) (*wallet.GetBalanceResponse, error) {
	return &wallet.GetBalanceResponse{
		UserId:  "1245",
		Balance: 40000,
	}, nil
}
