package repository

import (
	"context"
	"encoding/json"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/core/models"
	"net/http"
	"strings"
)

func (r *PaymentRepository) ReadBalanceInfoFromWallet(ctx context.Context, userID string) (float64, error) {
	// call ke localhost:3002/v1/wallet/balance/?user_id=10
	req, err := http.NewRequest("GET", "http://localhost:3002/v1/wallet/balance/?user_id="+userID, nil)
	if err != nil {
		// log error
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		// log error
	}
	defer resp.Body.Close()

	var data models.WalletBalanceResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		// log err
	}

	return float64(data.Balance), err
}

func (r *PaymentRepository) AppendBalanceInfoIntoWallet(ctx context.Context, userID string, amount float64) error {
	// call ke localhost:3002/v1/wallet/update
	url := "http://localhost:3002/v1/wallet/update"
	method := "POST"

	payload := models.UpdateBalancePayload{
		UserID: userID,
		Amount: amount,
	}
	payloadByte, err := json.Marshal(&payload)

	req, err := http.NewRequest(method, url, strings.NewReader(string(payloadByte)))
	if err != nil {
		// log error
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := r.Client.Do(req)
	if err != nil {
		// log error
	}
	defer res.Body.Close()

	var data models.WalletUpdateResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		// log err
	}

	return err
}
