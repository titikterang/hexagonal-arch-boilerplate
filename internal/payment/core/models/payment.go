package models

type BalanceResponse struct {
	UserID           string
	AvailableBalance float64
}

type UpdateBalancePayload struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type TransferBalancePayload struct {
	SourceUserID string
	TargetUserID string
	Amount       float64
}

type WalletBalanceResp struct {
	UserId  string `json:"userId"`
	Balance int    `json:"balance"`
}

type WalletUpdateResponse struct {
	Message      string `json:"message"`
	Success      bool   `json:"success"`
	FinalBalance int    `json:"finalBalance"`
}
