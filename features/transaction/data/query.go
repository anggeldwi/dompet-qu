package data

import (
	"dompet-qu/features/transaction"
	ud "dompet-qu/features/user/data"
	"dompet-qu/utils/externalapi"
	"errors"
	"time"

	"gorm.io/gorm"
)

type transactionQuery struct {
	db              *gorm.DB
	paymentMidtrans externalapi.MidtransInterface
}

func New(db *gorm.DB, mi externalapi.MidtransInterface) transaction.TransactionDataInterface {
	return &transactionQuery{
		db:              db,
		paymentMidtrans: mi,
	}
}

func (repo *transactionQuery) TopUp(userIdLogin int, inputTransaction transaction.Core) (*transaction.Core, error) {
	// Membuat transaksi pembayaran
	payment, errPay := repo.paymentMidtrans.NewTransaction(inputTransaction)
	if errPay != nil {
		return nil, errPay
	}

	// Ambil data pengguna dari basis data
	user := &ud.User{}
	if err := repo.db.First(user, userIdLogin).Error; err != nil {
		return nil, err
	}

	// Menambahkan jumlah top up ke saldo pengguna
	user.Balance += inputTransaction.Amount

	// Menyimpan perubahan saldo kembali ke basis data
	if err := repo.db.Save(user).Error; err != nil {
		return nil, err
	}

	// Membuat entitas transaksi dan menyimpannya ke basis data
	transactionModel := CoreToModelTransction(inputTransaction)
	transactionModel.PaymentType = payment.PaymentType
	transactionModel.Status = payment.Status
	transactionModel.VaNumber = payment.VaNumber
	transactionModel.ExpiredAt = payment.ExpiredAt

	tx := repo.db.Create(&transactionModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Mengonversi model transaksi kembali ke entitas core
	transactionCore := ModelToCoreTransaction(transactionModel)

	return &transactionCore, nil
}

// func CoreToModelTransaction(inputTransaction transaction.Core) {
// 	panic("unimplemented")
// }

// Transfer implements transaction.TransactionDataInterface.
func (repo *transactionQuery) Transfer(userIdLogin int, receiverPhoneNumber string, inputTransaction transaction.Core) (*transaction.Core, error) {
	// Mengecek apakah pengguna pengirim memiliki saldo yang cukup
	sender := &ud.User{}
	if err := repo.db.First(sender, userIdLogin).Error; err != nil {
		return nil, err
	}
	if sender.Balance < inputTransaction.Amount {
		return nil, errors.New("saldo tidak mencukupi untuk melakukan transfer")
	}

	// Mengurangi saldo pengirim
	sender.Balance -= inputTransaction.Amount
	if err := repo.db.Save(sender).Error; err != nil {
		return nil, err
	}

	// Mendapatkan data penerima dari basis data berdasarkan nomor telepon
	receiver := &ud.User{}
	if err := repo.db.Where("phone_number = ?", receiverPhoneNumber).First(receiver).Error; err != nil {
		return nil, errors.New("pengguna dengan nomor telepon tersebut tidak ditemukan")
	}

	// Menambahkan saldo penerima
	receiver.Balance += inputTransaction.Amount
	if err := repo.db.Save(receiver).Error; err != nil {
		return nil, err
	}

	// Membuat entitas transaksi dan menyimpannya ke basis data
	transactionModel := CoreToModelTransction(inputTransaction)
	transactionModel.ExpiredAt = time.Now()

	tx := repo.db.Create(&transactionModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Mengonversi model transaksi kembali ke entitas core
	transactionCore := ModelToCoreTransaction(transactionModel)

	return &transactionCore, nil
}
