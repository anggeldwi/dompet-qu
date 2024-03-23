package handler

import "dompet-qu/features/transaction"

type TopUpResponse struct {
	ID          string `json:"transaction_id"`
	UserID      uint   `json:"user_id"`
	PaymentType string `json:"payment_type"`
	Amount      int    `json:"amount"`
	Status      string `json:"status"`
	VaNumber    string `json:"va_number"`
	Bank        string `json:"bank"`
	ExpiredAt   string `json:"payment_expired"`
	CreatedAt   string `json:"created_at"`
}

func CoreToResponseTopUp(core *transaction.Core) TopUpResponse {
	return TopUpResponse{
		ID:          core.ID,
		UserID:      core.UserID,
		PaymentType: core.PaymentType,
		Amount:      core.Amount,
		Status:      core.Status,
		VaNumber:    core.VaNumber,
		Bank:        core.Bank,
		// PhoneNumber: core.PhoneNumber,
		ExpiredAt: core.ExpiredAt.Format("2006-01-02 15:04:05"),
		CreatedAt: core.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

type TransferResponse struct {
	ID          string `json:"transaction_id"`
	UserID      uint   `json:"user_id"`
	Amount      int    `json:"amount"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}

func CoreToResponseTransfer(core *transaction.Core) TransferResponse {
	return TransferResponse{
		ID:          core.ID,
		UserID:      core.UserID,
		Amount:      core.Amount,
		PhoneNumber: core.PhoneNumber,
		CreatedAt:   core.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
