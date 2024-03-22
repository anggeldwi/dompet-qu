package data

import (
	"dompet-qu/features/transaction"
	ud "dompet-qu/features/user/data"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID string `gorm:"type:varchar(36);primary_key" json:"id"`
	gorm.Model
	UserID         uint   `gorm:"not null"`
	JenisTransaksi string `gorm:"not null"`
	PaymentType    string
	Amount         int
	Status         string
	VaNumber       string
	Bank           string
	// BookingDate           string
	PhoneNumber string
	// Greeting              string
	// FullName              string
	// Email                 string
	// Quantity              int
	ExpiredAt time.Time
	User      ud.User
}

func CoreToModelTransction(input transaction.Core) Transaction {
	return Transaction{
		ID:             input.ID,
		UserID:         input.UserID,
		JenisTransaksi: input.JenisTransaksi,
		PaymentType:    input.PaymentType,
		Amount:         input.Amount,
		Status:         input.Status,
		VaNumber:       input.VaNumber,
		Bank:           input.Bank,
		PhoneNumber:    input.PhoneNumber,
		ExpiredAt:      input.ExpiredAt,
	}
}

func ModelToCoreTransaction(model Transaction) transaction.Core {
	return transaction.Core{
		ID:          model.ID,
		UserID:      model.UserID,
		Amount:      model.Amount,
		Status:      model.Status,
		VaNumber:    model.VaNumber,
		Bank:        model.Bank,
		PhoneNumber: model.PhoneNumber,
		ExpiredAt:   model.ExpiredAt,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
