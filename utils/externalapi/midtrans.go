package externalapi

import (
	"dompet-qu/app/configs"
	"dompet-qu/features/transaction"
	"errors"
	"time"

	mid "github.com/midtrans/midtrans-go"

	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransInterface interface {
	NewTransaction(transaction transaction.Core) (*transaction.Core, error)
	// CancelBookingPayment(bId string) error
}

type midtrans struct {
	client      coreapi.Client
	environment mid.EnvironmentType
}

func New() MidtransInterface {
	environment := mid.Sandbox
	var client coreapi.Client
	client.New(configs.MIDTRANS_SERVER_KEY, environment)

	return &midtrans{
		client: client,
	}
}

// NewBookingPayment implements Midtrans.
func (pay *midtrans) NewTransaction(transaction transaction.Core) (*transaction.Core, error) {
	req := new(coreapi.ChargeReq)
	req.TransactionDetails = mid.TransactionDetails{
		OrderID:  transaction.ID,
		GrossAmt: int64(transaction.Amount),
	}

	switch transaction.Bank {
	case "bca":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBca,
		}
	case "bni":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBni,
		}
	case "bri":
		req.PaymentType = coreapi.PaymentTypeBankTransfer
		req.BankTransfer = &coreapi.BankTransferDetails{
			Bank: mid.BankBri,
		}
	default:
		return nil, errors.New("sorry, payment not support")
	}

	res, err := pay.client.ChargeTransaction(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != "201" {
		return nil, errors.New(res.StatusMessage)
	}

	if len(res.VaNumbers) == 1 {
		transaction.VaNumber = res.VaNumbers[0].VANumber
	}

	if res.PaymentType != "" {
		transaction.PaymentType = res.PaymentType
	}

	if res.TransactionStatus != "" {
		transaction.Status = res.TransactionStatus
	}

	if expiredAt, err := time.Parse("2006-01-02 15:04:05", res.ExpiryTime); err != nil {
		return nil, err
	} else {
		transaction.ExpiredAt = expiredAt
	}

	return &transaction, nil
}
