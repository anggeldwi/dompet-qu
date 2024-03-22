package service

import (
	"dompet-qu/features/transaction"
	"errors"
)

type transactionService struct {
	transactionData transaction.TransactionDataInterface
}

func New(repo transaction.TransactionDataInterface) transaction.TransactionServiceInterface {
	return &transactionService{
		transactionData: repo,
	}
}

// TopUp implements transaction.TransactionServiceInterface.
func (service *transactionService) TopUp(userIdLogin int, inputTransaction transaction.Core) (*transaction.Core, error) {
	if inputTransaction.Bank == "" {
		return nil, errors.New("bank is required")
	}
	// if inputTransaction.PhoneNumber == "" {
	// 	return nil, errors.New("phone number is required")
	// }
	// if inputBooking.Greeting == "" {
	// 	return nil, errors.New("greeting is required")
	// }
	// if inputBooking.FullName == "" {
	// 	return nil, errors.New("full name number is required")
	// }
	// if inputBooking.Email == "" {
	// 	return nil, errors.New("email is required")
	// }
	// if inputBooking.BookingDate == "" {
	// 	return nil, errors.New("booking date is required")
	// }

	payment, err := service.transactionData.TopUp(userIdLogin, inputTransaction)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
