package data

import (
	"dompet-qu/features/transaction"
	ud "dompet-qu/features/user/data"
	"dompet-qu/utils/externalapi"

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

// // InsertTransaction implements transaction.TransactionDataInterface.
// func (repo *transactionQuery) TopUp(userIdLogin int, inputTrasaction transaction.Core) (*transaction.Core, error) {
// 	// var totalHargaKeseluruhan int
// 	// var packageGorm pd.Package
// 	// ts := repo.db.Where("tour_id = ? AND id = ?", inputTrasaction.TourID, inputBooking.PackageID).First(&packageGorm)
// 	// if ts.Error != nil {
// 	// 	return nil, ts.Error
// 	// }

// 	// if inputBooking.VoucherID != nil {
// 	// 	var voucherGorm vd.Voucher
// 	// 	ts := repo.db.Where("id = ?", inputBooking.VoucherID).First(&voucherGorm)
// 	// 	if ts.Error != nil {
// 	// 		return nil, ts.Error
// 	// 	}

// 	// 	var existingUseVoucher Booking
// 	// 	if err := repo.db.Where("user_id = ? AND voucher_id = ?", userIdLogin, inputBooking.VoucherID).First(&existingUseVoucher).Error; err == nil {
// 	// 		return nil, errors.New("user has already used this voucher")
// 	// 	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
// 	// 		return nil, err
// 	// 	}

// 	// 	totalHargaAwal := packageGorm.Price * inputBooking.Quantity
// 	// 	if totalHargaAwal < voucherGorm.DiscountValue {
// 	// 		return nil, errors.New("maaf, anda tidak bisa menggunakan voucher ini karena total pembayaran anda terlalu rendah")
// 	// 	} else {
// 	// 		// totalHargaKeseluruhan = ((packageGorm.JumlahTiket * packageGorm.Price) * inputBooking.Quantity) - voucherGorm.DiscountValue
// 	// 		totalHargaKeseluruhan = totalHargaAwal - voucherGorm.DiscountValue
// 	// 	}
// 	// } else {
// 	// 	// totalHargaKeseluruhan = (packageGorm.JumlahTiket * packageGorm.Price) * inputBooking.Quantity
// 	// 	totalHargaKeseluruhan = packageGorm.Price * inputBooking.Quantity
// 	// }

// 	// inputTrasaction.Amount = inputTrasaction.Amount

// 	payment, errPay := repo.paymentMidtrans.NewTransaction(inputTrasaction)
// 	if errPay != nil {
// 		return nil, errPay
// 	}

// 	transactionModel := CoreToModelTransction(inputTrasaction)
// 	transactionModel.PaymentType = payment.PaymentType
// 	transactionModel.Status = payment.Status
// 	transactionModel.VaNumber = payment.VaNumber
// 	transactionModel.ExpiredAt = payment.ExpiredAt

// 	tx := repo.db.Create(&transactionModel)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	transactionCore := ModelToCoreTransaction(transactionModel)

// 	return &transactionCore, nil
// }

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
