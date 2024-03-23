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

	payment, err := service.transactionData.TopUp(userIdLogin, inputTransaction)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// Transfer implements transaction.TransactionServiceInterface.
func (service *transactionService) Transfer(userIdLogin int, receiverPhoneNumber string, inputTransaction transaction.Core) (*transaction.Core, error) {
	if receiverPhoneNumber == "" {
		return nil, errors.New("receiver phone number is required")
	}

	transfer, err := service.transactionData.Transfer(userIdLogin, receiverPhoneNumber, inputTransaction)
	if err != nil {
		return nil, err
	}

	return transfer, nil
}
