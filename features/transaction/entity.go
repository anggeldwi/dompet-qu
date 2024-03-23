package transaction

import (
	"dompet-qu/features/user"
	"time"
)

type Core struct {
	ID             string
	UserID         uint
	PaymentType    string
	JenisTransaksi string
	Amount         int
	Status         string
	VaNumber       string
	Bank           string
	PhoneNumber    string
	ExpiredAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           user.Core
}

type TransactionDataInterface interface {
	TopUp(userIdLogin int, inputTransaction Core) (*Core, error)
	Transfer(userIdLogin int, receiverPhoneNumber string, inputTransaction Core) (*Core, error)
}

type TransactionServiceInterface interface {
	TopUp(userIdLogin int, inputTransaction Core) (*Core, error)
	Transfer(userIdLogin int, receiverPhoneNumber string, inputTransaction Core) (*Core, error)
}
