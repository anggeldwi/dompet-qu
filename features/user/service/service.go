package service

import (
	"dompet-qu/app/middlewares"
	"dompet-qu/features/user"
	"dompet-qu/utils/encrypts"
	"errors"

	"github.com/go-playground/validator"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// Insert implements user.UserServiceInterface.
func (service *userService) Insert(input user.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
	}

	if input.Role == "" {
		input.Role = "user"
	}

	err := service.userData.Insert(input)
	return err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(phoneNumber string, password string) (data *user.Core, token string, err error) {
	if phoneNumber == "" && password == "" {
		return nil, "", errors.New("no. hp dan password wajib diisi.")
	}
	if phoneNumber == "" {
		return nil, "", errors.New("no. hp wajib diisi.")
	}
	if password == "" {
		return nil, "", errors.New("password wajib diisi.")
	}

	data, err = service.userData.Login(phoneNumber, password)
	if err != nil {
		return nil, "", err
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai.")
	}

	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}

	return data, token, err
}

// SelectById implements user.UserServiceInterface.
func (service *userService) SelectById(userIdLogin int) (*user.Core, error) {
	result, err := service.userData.SelectById(userIdLogin)
	return result, err
}

// Update implements user.UserServiceInterface.
func (service *userService) Update(userIdLogin int, input user.Core) error {
	if userIdLogin <= 0 {
		return errors.New("invalid id.")
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
	}
	err := service.userData.Update(userIdLogin, input)
	return err
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(userIdLogin int) error {
	if userIdLogin <= 0 {
		return errors.New("invalid id")
	}
	err := service.userData.Delete(userIdLogin)
	return err
}
