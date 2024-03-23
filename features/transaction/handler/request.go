package handler

import (
	"dompet-qu/features/transaction"

	"github.com/google/uuid"
)

type TopUpRequest struct {
	JenisTransaksi string `json:"jenis_transaksi" form:"jenis_transaksi"`
	PaymentType    string `json:"payment_type" form:"payment_type"`
	Amount         int    `json:"amount" form:"amount"`
	Bank           string `json:"bank" form:"bank"`
}

type TransferRequest struct {
	JenisTransaksi string `json:"jenis_transaksi" form:"jenis_transaksi"`
	PaymentType    string `json:"payment_type" form:"payment_type"`
	Amount         int    `json:"amount" form:"amount"`
	PhoneNumber    string `json:"phone_number" form:"phone_number"`
}

func RequestToCoreTopUp(input TopUpRequest, userIdLogin uint) transaction.Core {
	// Generate UUID
	uid := uuid.New()

	// Ambil 10 karakter pertama dari UUID sebagai ID
	id := uid.String()[:10]

	return transaction.Core{
		ID:             id,
		UserID:         userIdLogin,
		JenisTransaksi: input.JenisTransaksi,
		PaymentType:    input.PaymentType,
		Amount:         input.Amount,
		Bank:           input.Bank,
	}
}

func RequestToCoreTransfer(input TransferRequest, userIdLogin uint, phoneNumber string) transaction.Core {
	// Generate UUID
	uid := uuid.New()

	// Ambil 10 karakter pertama dari UUID sebagai ID
	id := uid.String()[:10]

	return transaction.Core{
		ID:             id,
		UserID:         userIdLogin,
		JenisTransaksi: input.JenisTransaksi,
		Amount:         input.Amount,
		PhoneNumber:    phoneNumber,
	}
}
