package data

import (
	"dompet-qu/features/user"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	dataGorm := CoreToModel(input)

	tx := repo.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(phoneNumber string, password string) (data *user.Core, err error) {
	var userGorm User
	tx := repo.db.Where("phone_number = ?", phoneNumber).First(&userGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := userGorm.ModelToCore()
	return &result, nil
}

// SelectById implements user.UserDataInterface.
func (repo *userQuery) SelectById(userIdLogin int) (*user.Core, error) {
	var userDataGorm User
	tx := repo.db.First(&userDataGorm, userIdLogin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := userDataGorm.ModelToCore()
	return &result, nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(userIdLogin int, input user.Core) error {
	dataGorm := CoreToModel(input)
	tx := repo.db.Model(&User{}).Where("id = ?", userIdLogin).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(userIdLogin int) error {
	tx := repo.db.Delete(&User{}, userIdLogin)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}
