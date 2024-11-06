package models

type BalanceResponse struct {
	UserID           string
	AvailableBalance float64
}

type UpdateBalancePayload struct {
	UserID string
	Amount float64
}

type DatastoreBalanceResponse struct {
	UserID string
	List   []string
}
